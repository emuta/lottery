package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"

	"lottery/ssc/cqssc/web"
)

var (
	port string
	addr string
)

func init() {
	flag.StringVar(&port, "port", "80", "The port of service listening")
	flag.StringVar(&addr, "addr", "localhost:3721", "The address of grpc server address listening")

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetLevel(log.DebugLevel)
}

func newRouter() *mux.Router {
	handler := web.NewHandler(addr)
	routes := handler.GetRoutes()
	router := mux.NewRouter()
	for _, route := range routes {
		router.Path(route.Path).Handler(route.Handler).Methods(route.Method)
	}

	return router
}

func main() {
	flag.Parse()

	l := fmt.Sprintf(":%s", port)
	log.Infof("Server starting with %s", l)
	if err := http.ListenAndServe(l, newRouter()); err != nil {
		log.Fatal(err)
	}

}
