// Web Server

package main

import (
   "fmt"
   "log"
   "net/http"
   )

   func main() {
       http.HandleFunc("/", handler) //cada petición manda a llamar un handler
       log.Fatal(http.ListenAndServe("localhost:8000", nil))
   }

   //handler echoes the Path component of the request URL r.
   func handler(w http.ResponseWriter, r *http.Request) {
    resp, err := http.Get("http://google.com/")
    if err != nil {
      panic(err)
    }
   fmt.Fprintf(w, "URL.Path = %q\n", resp)
   }