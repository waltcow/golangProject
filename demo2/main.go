package main

import (
	"github.com/waltcow/golangProject/demo2/storage"
	"github.com/waltcow/golangProject/demo2/handles"
	"log"
	"net/http"
)

func main() {
	db := storage.NewMemHashTable()
	mux := http.NewServeMux()
	mux.Handle("/get", handles.GetKey(db))
	mux.Handle("/set", handles.SetKey(db))
	log.Printf("serving on port 8080")

	// http.ListenAndServe takes in an http.Handler as its second parameter.
	// since ServeMux implements a ServeHTTP function, it is also an http.Handler,
	// so we can pass it here.
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}