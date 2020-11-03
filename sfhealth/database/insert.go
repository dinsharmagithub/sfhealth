package database

import (
	"context"
	"fmt"
	"log"

	"github.com/dinsharmagithub/sfhealth/proto"
)

//Insert Inserts one record in application specific table
//TODO it has to named better or under different package like sfhealth/crud
func Insert(ctx context.Context, req *proto.CreateRequest) error {
	log.Printf("Inside Insert()\n")

	txn, err := GetDbConn().Begin()

	if err != nil {
		log.Fatal(err)
	}
	defer txn.Rollback()
	log.Printf("Transaction started \n")

	fmt.Printf("Creating prepared statement\n")
	stmtStr, err := txn.Prepare("INSERT INTO restaurant_scores(business_id, business_name, business_address, business_city, business_state, business_postal_code, business_latitude, business_longitude, business_location, business_phone_number, inspection_id, inspection_date, inspection_score, inspection_type, violation_id, violation_description, risk_category, neighborhoods_old, police_districts, supervisor_districts, fire_prevention_districts, zip_codes, analysis_neighborhoods) VALUES( $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23 )")
	//TODO clean
	// stmtStr, err := txn.Prepare("INSERT INTO restaurant_scores(business_id, business_name, business_address, business_city, business_state, business_postal_code, business_latitude, business_longitude, business_location, business_phone_number, inspection_id, inspection_date, inspection_score, inspection_type, violation_id, violation_description, risk_category, neighborhoods_old, police_districts, supervisor_districts, fire_prevention_districts, zip_codes, analysis_neighborhoods) VALUES( 1, 2, 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test' )")
	if err != nil {
		fmt.Printf("Error in creating statement %v", err)
		return err
	}
	defer stmtStr.Close()
	log.Printf("Statement Insertd \n")

	log.Printf("Executing the statement for business_id %v \n", req.GetRecord().GetBusinessId())
	//Keeping long statement as punch cards time has gone
	res, err := stmtStr.Exec(req.GetRecord().GetBusinessId(), req.GetRecord().GetBusinessName(), req.GetRecord().GetBusinessAddress(), req.GetRecord().GetBusinessCity(), req.GetRecord().GetBusinessState(), req.GetRecord().GetBusinessPostalCode(), req.GetRecord().GetBusinessLatitude(), req.GetRecord().GetBusinessLongitude(), req.GetRecord().GetBusinessLocation(), req.GetRecord().GetBusinessPhoneNumber(), req.GetRecord().GetInspectionId(), req.GetRecord().GetInspectionDate(), req.GetRecord().GetInspectionScore(), req.GetRecord().GetInspectionType(), req.GetRecord().GetViolationId(), req.GetRecord().GetViolationDescription(), req.GetRecord().GetRiskCategory(), req.GetRecord().GetNeighborhoodsOld(), req.GetRecord().GetPoliceDistricts(), req.GetRecord().GetSupervisorDistricts(), req.GetRecord().GetFirePreventionDistricts(), req.GetRecord().GetZipCodes(), req.GetRecord().GetAnalysisNeighborhoods())
	if err != nil {
		log.Printf("Error while inserting rows %v", err)
	}
	log.Printf("INSERT done with Result = %v\n doing commit now \n", res)

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Exiting Insert()\n")
	return nil
}
