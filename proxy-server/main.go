package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	"net/http"
	"os"
	"http-gateway/proxy-server/proxy"
	"http-gateway/proxy-server/proxy/handler"
)

var opts struct {
	Port int  `short:"p" long:"port" description:"the server port to listen" default:"9999"`
	Help bool `short:"h" long:"help" descrition:"the help message"`
}

func main() {
	server := proxy.ReverseProxy{}

	server.RequestHandler = handler.RequestHandler

	server.ResponseHandler = handler.ResponseHandler

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server.ServeHTTP(w, r)
	})

	parser := flags.NewParser(&opts, flags.None)
	_, err := parser.ParseArgs(os.Args)

	if err != nil {
		fmt.Println("parse args error")
		printUsage(parser)
		return
	}

	if opts.Help {
		printUsage(parser)
		return
	}

	log.Println("listening on port:", opts.Port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", opts.Port), nil)

	if err != nil {
		log.Println(err)
	}
}

func printUsage(parser *flags.Parser) {
	parser.WriteHelp(os.Stdout)
}
