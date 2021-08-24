package apis

import (
	"fmt"
	"net/http"
)

func FindAll(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "This is find all ")
}

func Search(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "This is search one ")
}
