package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/fcgi"
	"runtime"
)

var runserver = flag.Bool("s", false, "Flag for running server")

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	Mux := SetupMux()
	flag.Parse()
	var err error
	if *runserver {
		Log("STARTING SERVER")
		fmt.Println("STARTING SERVER")
		err = http.ListenAndServe(":8080", Mux)
	} else {
		err = fcgi.Serve(nil, Mux)
	}
	if err != nil {
		ErrLog(err)
	}
}
