package main

import (
	"fmt"
	"testing"

	"github.com/dinsharmagithub/sfhealth/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

//TODO: make better if time permits. Query exact value and then delete
var idForUpdate int64 = 2
var genRec = proto.Record{
	Id:                      &idForUpdate,
	BusinessId:              4794,
	BusinessName:            "VICTOR'S",
	BusinessAddress:         "210 TOWNSEND St",
	BusinessCity:            "San Francisco",
	BusinessState:           "CA",
	BusinessPostalCode:      94107,
	BusinessLatitude:        37.778634,
	BusinessLongitude:       -122.393089,
	BusinessLocation:        "(37.778634, -122.393089)",
	BusinessPhoneNumber:     "14155607018",
	InspectionId:            "4794_20181030",
	InspectionDate:          "10/30/2018 12:00:00 AM",
	InspectionScore:         71,
	InspectionType:          "Routine - Unscheduled",
	ViolationId:             "4794_20181030_103138",
	ViolationDescription:    "Improper storage use or identification of toxic substances",
	RiskCategory:            "Low Risk",
	NeighborhoodsOld:        "34",
	PoliceDistricts:         "2",
	SupervisorDistricts:     "9",
	FirePreventionDistricts: "6",
	ZipCodes:                "28856",
	AnalysisNeighborhoods:   "34",
}

func TestIntegration(t *testing.T) {
	//TODO: If time permits, separate out tests for each operation
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		t.Errorf("Failed to Dial: %v", err)
	}

	ctx := context.Background()
	client := proto.NewCrudServiceClient(conn)

	req := &proto.CreateRequest{Record: &genRec}

	if response, err := client.Create(ctx, req); err == nil {
		fmt.Printf("Returned value = %v", response.Result)
	} else {
		t.Errorf("Returned error = %v", err)
	}
	id := int64(2)
	readReq := &proto.ReadRequest{Id: &id}

	if response, err := client.Read(context.Background(), readReq); err == nil {
		fmt.Printf("Returned value = %v\n", response.Result)
	} else {
		t.Errorf("Returned error = %v\n", err)
	}

	updateReq := &proto.UpdateRequest{Record: &genRec}

	//TODO make it better. Just update a record with hardcoded primary key value
	*updateReq.Record.Id = 2
	if response, err := client.Update(context.Background(), updateReq); err == nil {
		fmt.Printf("Returned value = %v\n", response.Result)
	} else {
		t.Errorf("Returned error = %v\n", err)
	}

	//TODO make it better. Just deleting the one created with ID 2
	delReq := &proto.DeleteRequest{Id: 2}

	if response, err := client.Delete(context.Background(), delReq); err == nil {
		fmt.Printf("Returned value = %v", response.Result)
	} else {
		t.Errorf("Returned error = %v", err)
	}
}
