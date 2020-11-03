package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/dinsharmagithub/sfhealth/config"
	"github.com/dinsharmagithub/sfhealth/database"
	"github.com/dinsharmagithub/sfhealth/proto"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var logFilePath *string

func init() {
	// TODO removing hardcoding, check the value with Lookup
	logFilePath = flag.String("logfile", "./log/sfhealthserver.log", "Log file path")
}

type sfHealthServer struct {
}

func (s *sfHealthServer) Create(ctx context.Context, request *proto.CreateRequest) (*proto.CreateResponse, error) {
	log.Printf("Create() handler called\n")

	err := database.Insert(ctx, request)
	if err != nil {
		log.Fatal(err)
		//TODO change hardcoding below
		return &proto.CreateResponse{Result: 2}, err
	}

	log.Printf("Exiting Create() handler\n")
	//TODO change hardcoding below
	return &proto.CreateResponse{Result: 0}, nil
}
func (s *sfHealthServer) Delete(ctx context.Context, request *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	log.Printf("Delete() handler called\n")

	err := database.Delete(ctx, request)
	if err != nil {
		log.Fatal(err)
		//TODO change hardcoding below
		return &proto.DeleteResponse{Result: 2}, err
	}

	log.Printf("Exiting Delete() handler\n")
	//TODO change hardcoding below
	return &proto.DeleteResponse{Result: 0}, nil
}
func (s *sfHealthServer) Update(ctx context.Context, request *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	log.Printf("Create() handler called\n")

	err := database.Update(ctx, request)
	if err != nil {
		log.Fatal(err)
		//TODO change hardcoding below
		return &proto.UpdateResponse{Result: 2}, err
	}

	log.Printf("Exiting Update() handler\n")
	//TODO change hardcoding below
	return &proto.UpdateResponse{Result: 0}, nil
}
func (s *sfHealthServer) Read(ctx context.Context, request *proto.ReadRequest) (*proto.ReadResponse, error) {
	log.Printf("Read() handler called\n")

	toReturn, err := database.Read(ctx, request)
	if err != nil {
		log.Fatal(err)
		return &proto.ReadResponse{Result: toReturn}, err
	}

	log.Printf("Exiting Read() handler\n")
	return &proto.ReadResponse{Result: toReturn}, nil
}

//Initialize Initializes the service
func Initialize(logPath string) context.Context {
	_ = logPath
	ctx := context.Background()

	if err := config.Initialize(ctx); err != nil {
		log.Fatal(err)
	}

	if err := database.Initialize(ctx, *config.Get()); err != nil {
		log.Fatal(err)
	}
	return ctx
}
func main() {
	fmt.Printf("Starting SF Health Server\n")

	flag.Parse()
	if *logFilePath == "" {
		fmt.Println("logfile option not specified")
		os.Exit(1)
	}
	//TODO if itme permits,
	//   use the provide file for logging. Its unused as of now
	//   Make the below code better

	ctx := Initialize(*logFilePath)
	defer database.CloseDbConn(ctx)

	//TODO use config file for below hardcode info
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterCrudServiceServer(srv, &sfHealthServer{})
	reflection.Register(srv)
	if e := srv.Serve(listener); e != nil {
		panic(err)
	}
}
