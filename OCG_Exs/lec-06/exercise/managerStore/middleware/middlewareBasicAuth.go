package middleware

import (
	"fmt"
	"lec-06/exercise/managerStore/model"
	"net/http"
)

func BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		usename, password, ok := r.BasicAuth()
		fmt.Println("ur", usename)
		fmt.Println("pw", password)
		if !ok || !checkUser(usename, password) {
			rw.Header().Set("WWW-Authenticate", `Basic realm="Account Invalid"`)
			rw.WriteHeader(401)
			rw.Write([]byte("Unauthorised\n"))
			fmt.Println(accounts)
			return
		}
		handler(rw, r)
	}
}

//dummy data
var (
	accounts = model.User{UserName: "Duy", Password: "123"}
)

func checkUser(username, password string) bool {
	return username == accounts.UserName && password == accounts.Password
}
