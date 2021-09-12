package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
	"worker/mail"
	"worker/rmq"
)

// Worker defines a worker
type Worker struct {
	wg      *sync.WaitGroup
	mailer  mail.Mailer
	rmqChan *rmq.Rabbit
	ctx     context.Context
	db      *sql.DB
}

// NewWorker creates new worker and gets incoming parammeters
func NewWorker(ctx context.Context, wg *sync.WaitGroup, db *sql.DB, mailer mail.Mailer, rmqChan *rmq.Rabbit) *Worker {
	return &Worker{
		ctx:     ctx,
		wg:      wg,
		mailer:  mailer,
		rmqChan: rmqChan,
		db:      db,
	}
}

// Start starts worker to process message
//  Processing logic:
//    1. Wait for message
//    2. Send email with mailer (Sendgrid client)
//    3. Update database (thankyou_email_sent) to prevent duplicated emails
func (w *Worker) Start() {
	if w.mailer == nil || w.db == nil {
		fmt.Println("cannot start worker since mailer is nil")
		return
	}
	sttm, err := w.db.Prepare("UPDATE `order` SET thankyou_email_sent = ? WHERE id =?;")
	if err != nil {
		fmt.Println("Can't prepare statement to update thankyou_email_sent")
	}
	msgs, err := w.rmqChan.Consume()
	rmq.FailOnError(err, "Failed to register a consumer")

	for {
		select {
		case msg := <-msgs:
			var em mail.EmailContent
			//convert byte to struct
			err := json.Unmarshal(msg.Body, &em)
			if err != nil {
				fmt.Println("Can't unMarshal body msgs")
				continue
			}
			err = w.mailer.Send(&em)
			if err != nil {
				fmt.Println("Can't send msgs in sendgrid")
				continue
			}
			_, err = sttm.Query(true, em.ID)
			if err != nil {
				fmt.Println("Can't query to update thankyou_email_sent to true")
			}
		case <-w.ctx.Done():
			fmt.Println("Exiting worker")
			w.wg.Done()
			return
		}
	}
}

// for {
// 	select {
// 	case em := <-w.inChan: // call interface Send from maller.pkg and send after that update rows
// 		err := w.mailer.Send(em)
// 		if err != nil {
// 			fmt.Println("Cannot send email due to error: ", err)
// 			continue
// 		}
// 		// update sql data
// 		_, err = w.db.Exec("UPDATE `order` SET thankyou_email_sent = ? WHERE id = ?", true, em.ID)
// 		if err != nil {
// 			fmt.Println("Cannot update thankyou_email_sent to true")
// 		}
// 	case <-w.ctx.Done():
// 		fmt.Println("Exiting worker")
// 		w.wg.Done()
// 		return
// 	}
// }
