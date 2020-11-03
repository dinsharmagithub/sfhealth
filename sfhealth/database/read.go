package database

import (
	"context"
	"fmt"
	"log"

	"github.com/dinsharmagithub/sfhealth/proto"
)

//Read Reads all records from the application table in the database
func Read(ctx context.Context, req *proto.ReadRequest) ([]*proto.Record, error) {
	//TODO:
	// 1. implement paging
	// 2. query based read
	// 3. check and implement input primary key ID if provided by the client.
	log.Printf("Inside Read()\n")

	var toReturn []*proto.Record

	fmt.Printf("Querying\n")
	rows, err := GetDbConn().Query("SELECT * FROM restaurant_scores")
	// rows, err := GetDbConn().Query("SELECT business_id, business_postal_code, inspection_score  FROM restaurant_scores")

	if err != nil {
		fmt.Printf("Error in executing query statement %v", err)
		return toReturn, err
	}
	for rows.Next() {

		record := proto.Record{}

		if err := rows.Scan(&record.Id, &record.BusinessId, &record.BusinessName, &record.BusinessAddress, &record.BusinessCity, &record.BusinessState, &record.BusinessPostalCode, &record.BusinessLatitude, &record.BusinessLongitude, &record.BusinessLocation, &record.BusinessPhoneNumber, &record.InspectionId, &record.InspectionDate, &record.InspectionScore, &record.InspectionType, &record.ViolationId, &record.ViolationDescription, &record.RiskCategory, &record.NeighborhoodsOld, &record.PoliceDistricts, &record.SupervisorDistricts, &record.FirePreventionDistricts, &record.ZipCodes, &record.AnalysisNeighborhoods); err != nil {
			fmt.Printf("Error in executing query statement %v", err)
			return toReturn, err
		}
		toReturn = append(toReturn, &record)
	}
	log.Printf("Returning (%d) records from Read()\n", len(toReturn))
	return toReturn, nil
}
