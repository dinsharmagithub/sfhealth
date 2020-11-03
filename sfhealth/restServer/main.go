package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/dinsharmagithub/sfhealth/proto"
	"google.golang.org/grpc"
)

var ctx context.Context
var client proto.CrudServiceClient

//TODO use config packge to pick from file
var restPort = ":8080"

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Inside rootHandler()\n")
	json.NewEncoder(w).Encode(struct{}{})
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	var record proto.Record
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	req := &proto.CreateRequest{Record: &record}

	toReturn := struct{ Response string }{Response: "Operation successful"}
	if response, err := client.Create(ctx, req); err == nil {
		log.Printf("Returned value = %v", response.Result)
	} else {
		log.Printf("Returned error = %v", err)
		toReturn = struct{ Response string }{Response: err.Error()}
	}
	json.NewEncoder(w).Encode(toReturn)
}
func readHandler(w http.ResponseWriter, r *http.Request) {
	id := int64(2)
	readReq := &proto.ReadRequest{Id: &id}

	response, err := client.Read(context.Background(), readReq)
	if err == nil {
		log.Printf("Returned %v\n", response.Result)
	} else {
		fmt.Printf("Returned error = %v\n", err)
		return
	}
	json.NewEncoder(w).Encode(response.Result)
}
func updateHandler(w http.ResponseWriter, r *http.Request) {
	var record proto.Record
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		log.Printf("Error encountered: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//TODO more data validation
	updateReq := &proto.UpdateRequest{Record: &record}
	response, err := client.Update(context.Background(), updateReq)
	if err == nil {
		log.Printf("Returned value = %v\n", response.Result)
		json.NewEncoder(w).Encode(response.Result)
	} else {
		log.Printf("Returned error = %v\n", err)
		json.NewEncoder(w).Encode(err)
	}
}
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	//TODO : to be implemented
	// delReq := &proto.DeleteRequest{Id: 2}

	// if response, err := client.Delete(context.Background(), delReq); err == nil {
	// 	fmt.Printf("Returned value = %v", response.Result)
	// } else {
	// 	fmt.Printf("Returned error = %v", err)
	// }
}

func initialize() {
	//TODO use config package for getting following hardcoded values from file
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	ctx = context.Background()
	client = proto.NewCrudServiceClient(conn)

	// Start REST server
	log.Println("Starting REST Server at http://localhost%s\n", restPort)
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/list", readHandler)
	http.HandleFunc("/insert", createHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)
	//TODO Monitor, handle errors and bring down gracefully in case of panic
	http.ListenAndServe(restPort, nil)
}
func main() {
	//TODO reorganize and make code robust
	initialize()
}
