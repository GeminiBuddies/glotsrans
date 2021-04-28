package main

import (
	"bufio"
	"glotsrans/router"
	"net"
	"net/http"
	"os"
)

func main() {
	l, _ := net.Listen("tcp", ":22345")

	s := &http.Server{
		Addr:           l.Addr().String(),
		Handler:        router.Create(),
		MaxHeaderBytes: 1 << 20,
	}
	s.SetKeepAlivesEnabled(true)

	go s.Serve(l)

	bufio.NewReader(os.Stdin).ReadLine()

	s.Close()

	return
}
