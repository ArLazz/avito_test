package main

import (
	"flag"
	"log"
	"net/http"
	"slugservice/service"
	"slugservice/pkg"
	"slugservice/endpoint"
)


func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
	)

	flag.Parse()

	srv := service.NewService()

	endpoints := endpoint.Endpoints{
		CreateSegmentEndpoint:   endpoint.MakeCreateSegmentEndpoint(srv),
		DeleteSegmentEndpoint: endpoint.MakeDeleteSegmentEndpoint(srv),
		AddUserEndpoint: endpoint.MakeAdduserEndpoint(srv),
		UserSegmentsEndpoint: endpoint.MakeUserSegmentsEndpoint(srv),
		UsersHistoryEndpoint: endpoint.MakeUsersHistoryEndpoint(srv),
	}

	handler := pkg.NewHTTPServer(endpoints)

	log.Printf("slugservice is running on %s\n", *httpAddr)

	if err := http.ListenAndServe(*httpAddr, handler); err != nil {
		log.Println(err)
	}
}
