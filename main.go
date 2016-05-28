package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
)

var (
	str      = `R0lGODlhAQABAIAAAP///wAAACwAAAAAAQABAAACAkQBADs=`
	logger   = log.New(os.Stdout, "", 0)
	hostPort = ":9090"
)

func main() {
	http.HandleFunc("/h1", printWithFmt)
	http.HandleFunc("/h2", printWithLog)
	err := http.ListenAndServe(hostPort, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func printWithFmt(w http.ResponseWriter, r *http.Request) {
	go func() {
		fmt.Fprintln(os.Stdout, str)
	}()
	fmt.Fprintln(w, "h1")
}

func printWithLog(w http.ResponseWriter, r *http.Request) {
	go func() {
		logger.Println(str)
	}()
	fmt.Fprintln(w, "h2")
}
