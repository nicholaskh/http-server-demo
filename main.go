package main

import (
    "fmt"
    "net/http"
    //"runtime"
)

func init() {
    //runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
    http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request){
        w.Write([]byte("hello world"))
    });
    port := ":8877"
    fmt.Println("Http server start on " + port)
    http.ListenAndServe(port, nil);
}
