package handlers

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
)

type Hello struct{
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello{
	return &Hello{l}
}

	// Description: open page /nguyen, shall return Hello world in log.
	// and if the server receive an invalid http request, shall return data bad request.

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request){
	h.l.Println("Hello World")
	
	d, err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw,"Hello %s \n", d)
}