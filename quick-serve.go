package main

import (
"net/http"
"fmt"
)

func main() {
	fileHandler := http.FileServer(http.Dir("./"))
	http.Handle("/", fileHandler)
	
	fmt.Println("Serving on *:80")
	http.ListenAndServe(":80", nil)
}