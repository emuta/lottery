package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"github.com/emuta/go-rabbitmq/pubsub/publisher"
	
	pb "lottery/ssc/cqssc/proto"
	"lottery/ssc/cqssc/server/repository"
	"lottery/ssc/cqssc/server"
	
)

const ServiceName = "lottery.ssc.cqssc"

var (
	port     string
	exchange string
)

func init() {
	flag.StringVar(&port, "port", "3721", "The port of service listen")
	flag.StringVar(&exchange, "exchange", "pubsub", "The exchange name of RabbitMQ used")
	flag.Parse()

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func NewPostgres(url string, logMode bool) *gorm.DB {
	if url == "" {
		log.Fatal("Empty URL of postgresql for gorm connection")
		return nil
	}
	conn, err := gorm.Open("postgres", url)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect PostgreSQL server")
		return nil
	}
	log.Info("Connected to PostgreSQL server")
	conn.LogMode(logMode)
	return conn
}

func main() {
	db := NewPostgres(os.Getenv("PG_URL"), true)
	puber := publisher.NewAppPublisher(os.Getenv("RABBITMQ_URL"), exchange, ServiceName)
	repo := repository.NewRepository(db)
	srv := server.NewCqsscServiceServer(repo, puber)

	s := grpc.NewServer()
	pb.RegisterCqsscServiceServer(s, srv)

	addr := fmt.Sprintf("0.0.0.0:%s", port)
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.WithError(err).Fatal("Failed to listen %s", addr)
	}

	log.Info("Initialized all components")
	log.Infof("Server starting with addr -> %s", addr)
	log.Fatal(s.Serve(l))
}
