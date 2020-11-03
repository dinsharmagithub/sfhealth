package database

import (
	"context"
	"fmt"
	"log"

	"github.com/dinsharmagithub/sfhealth/proto"
)

//Update Updates a record in the application table in the DB
func Update(ctx context.Context, req *proto.UpdateRequest) error {
	log.Printf("Inside Update()\n")

	txn, err := GetDbConn().Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer txn.Rollback()
	log.Printf("Transaction started \n")

	fmt.Printf("Creating prepared statement\n")
	stmtStr, err := txn.Prepare("UPDATE restaurant_scores SET business_id=$1, business_name=$2, business_address=$3, business_city=$4, business_state=$5, business_postal_code=$6, business_latitude=$7, business_longitude=$8, business_location=$9, business_phone_number=$10, inspection_id=$11, inspection_date=$12, inspection_score=$13, inspection_type=$14, violation_id=$15, violation_description=$16, risk_category=$17, neighborhoods_old=$18, police_districts=$19, supervisor_districts=$20, fire_prevention_districts=$21, zip_codes=$22, analysis_neighborhoods=$23 WHERE id=$24")
	//TODO clean
	// stmtStr, err := txn.Prepare("INSERT INTO restaurant_scores(business_id, business_name, business_address, business_city, business_state, business_postal_code, business_latitude, business_longitude, business_location, business_phone_number, inspection_id, inspection_date, inspection_score, inspection_type, violation_id, violation_description, risk_category, neighborhoods_old, police_districts, supervisor_districts, fire_prevention_districts, zip_codes, analysis_neighborhoods) VALUES( 1, 2, 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test', 'test' )")
	if err != nil {
		fmt.Printf("Error in creating statement %v", err)
		return err
	}
	defer stmtStr.Close()
	log.Printf("Statement created \n")

	log.Printf("Executing the statement for business_id %v \n", req.GetRecord().GetBusinessId())
	//Keeping long statement as punch cards time has gone
	res, err := stmtStr.Exec(req.GetRecord().GetBusinessId(), req.GetRecord().GetBusinessName(), req.GetRecord().GetBusinessAddress(), req.GetRecord().GetBusinessCity(), req.GetRecord().GetBusinessState(), req.GetRecord().GetBusinessPostalCode(), req.GetRecord().GetBusinessLatitude(), req.GetRecord().GetBusinessLongitude(), req.GetRecord().GetBusinessLocation(), req.GetRecord().GetBusinessPhoneNumber(), req.GetRecord().GetInspectionId(), req.GetRecord().GetInspectionDate(), req.GetRecord().GetInspectionScore(), req.GetRecord().GetInspectionType(), req.GetRecord().GetViolationId(), req.GetRecord().GetViolationDescription(), req.GetRecord().GetRiskCategory(), req.GetRecord().GetNeighborhoodsOld(), req.GetRecord().GetPoliceDistricts(), req.GetRecord().GetSupervisorDistricts(), req.GetRecord().GetFirePreventionDistricts(), req.GetRecord().GetZipCodes(), req.GetRecord().GetAnalysisNeighborhoods(), req.GetId())
	if err != nil {
		log.Printf("Error while inserting rows %v", err)
	}
	log.Printf("UPDATE done with Result = %v\n doing commit now \n", res)

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Exiting Update()\n")
	return nil
}
