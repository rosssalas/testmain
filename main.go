package main

import (
	"github.com/rosssalas/mydb"
	"github.com/rosssalas/myhttp"
	"net/http"
	"log"
)

func main() {
	mydb.Init()
	http.HandleFunc("/",myhttp.MyHandler)
    log.Fatal(http.ListenAndServe(":1234", nil))
}
