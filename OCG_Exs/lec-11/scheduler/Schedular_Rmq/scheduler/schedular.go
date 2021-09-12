package scheduler

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
	"worker/mail"
	"worker/rmq"

	"github.com/robfig/cron/v3"
)

/*
Goal:
 - is to scan order table for getting new order to send thank you email
 - run at every minute
 - after send the email -> set thankyou_email_sent to true

Definition:
 - if created_at < now - 1 minutes && thankyou_email_sent == false -> schedule to send email
*/

const (
	DefaultThankyouSubject   = "Thank you for purchasing from mystore.com"
	DefaultThankyouBodyPlain = "Thank you for purchasing from our store. Here's your order details:"
	DefaultThankyouBodyHtml  = "<strong>Thank you for purchasing from our store. Here's your order details:</strong>"
	DefaultFromName          = "My Store Owner"
	DefaultFromEmail         = "support@mystore.com"
)

type Scheduler struct {
	db *sql.DB
	c  *cron.Cron
	// outChan chan<- *mail.EmailContent
	rmqChan *rmq.Rabbit
	ctx     context.Context
}

func NewScheduler(ctx context.Context, db *sql.DB, rmqChan *rmq.Rabbit) *Scheduler {
	return &Scheduler{
		ctx:     ctx,
		db:      db,
		c:       cron.New(cron.WithSeconds()), //New returns a new Cron job runner, in the Local time zone.
		rmqChan: rmqChan,
		// outChan: ch,
	}
}

func (sched *Scheduler) Start() {
	// runs this function every minute
	sched.c.AddFunc("0 * * * * *", sched.scheduleJob)
	sched.c.Start()
}

func (sched *Scheduler) Stop() {
	fmt.Println("Stopping scheduler")
	sched.c.Stop()
}

func (sched *Scheduler) scheduleJob() {
	fmt.Printf("Scanning for new order(s) at %v\n", time.Now().Format("2006-Jan-02 15:04:05"))
	resp, err := sched.getEmailForSending()
	if err != nil {
		return
	}
	fmt.Printf("Scheduling %v email(s) at %v\n", len(resp), time.Now().Format("2006-Jan-02 15:04:05"))
	for _, em := range resp {
		body := em.String()
		err = sched.rmqChan.Publish(body)
		if err != nil {
			fmt.Println("Can't channel mes")
		}
	}
}

// getEmailForSending get email and fill up enough information ready for sending
func (sched *Scheduler) getEmailForSending() ([]*mail.EmailContent, error) {
	resp, err := sched.scanFromDB()
	if err != nil {
		return resp, err
	}
	// fill FromUser
	// why we can set FromUser here? aws:
	// I think Sender usualy only  uses a fixed mail or if we want divide the work for other admin mail
	for _, emailContent := range resp {
		emailContent.FromUser = &mail.EmailUser{
			Name:  DefaultFromName,
			Email: DefaultFromEmail,
		}
	}

	return resp, err
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// scanFromDB get all orders that match the predefined condition (created_at < now - 1 min && thankyou_email_sent == falses)
func (sched *Scheduler) scanFromDB() ([]*mail.EmailContent, error) {
	var resp []*mail.EmailContent
	fromTime := time.Now().Add(-time.Minute * 2)
	// Stmt is a struct which  prepared statement                                                                                   // subtract by 2 minutes - why not one?  because if query take more than one minute,overflow time will not be queried in next  (eg: time name 7:30, query from time 7:29, but query take 70 seconds, so the next timenow is 7:31:10 and start time at 7:30:10,  so 10 seconds between them can not query  )
	stmt, err := sched.db.Prepare("SELECT id, customer_name, email FROM `order` WHERE created_at >= ? AND thankyou_email_sent = ?;") // Prepare query to use latter? but why? When you expect to execute the sameSQL
	if err != nil {                                                                                                                  // repeatedly, you can use an sql.Stmt to prepare the SQL statement in advance, then execute it as needed.
		fmt.Println("Cannot prepare statement, ", err)
		return nil, err
	}
	rows, err := stmt.Query(fromTime, false)
	if err != nil || rows == nil {
		fmt.Printf("Cannot query from db due to error: %v, %v\n", err, rows == nil)
		return nil, err
	}
	// MUST to call this function at the end to free connection to mysql
	defer rows.Close()

	var id int64
	var email, name string
	for rows.Next() {
		err = rows.Scan(&id, &name, &email)
		if err != nil {
			fmt.Println("Cannot scan row due to error: ", err)
			continue
		}
		resp = append(resp, &mail.EmailContent{
			ID:               id,
			Subject:          DefaultThankyouSubject,
			PlainTextContent: DefaultThankyouBodyPlain,
			HtmlContent:      DefaultThankyouBodyHtml,
			ToUser: &mail.EmailUser{
				Name:  name,
				Email: email,
			},
		})
	}
	return resp, nil
}
