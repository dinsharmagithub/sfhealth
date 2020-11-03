package database

import (
	"context"
	"fmt"
	"log"

	"github.com/dinsharmagithub/sfhealth/proto"
)

//Delete Removes a record from the application table in the DB
func Delete(ctx context.Context, req *proto.DeleteRequest) error {
	log.Printf("Inside Delete()\n")

	txn, err := GetDbConn().Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer txn.Rollback()
	log.Printf("Transaction started \n")

	fmt.Printf("Creating prepared statement\n")
	stmtStr, err := txn.Prepare("DELETE FROM restaurant_scores WHERE ID = $1")
	if err != nil {
		fmt.Printf("Error in creating statement %v", err)
		return err
	}
	defer stmtStr.Close()
	log.Printf("Statement created \n")

	log.Printf("Executing the statement for ID %v \n", req.GetId())
	//Keeping long statement as punch cards time has gone
	res, err := stmtStr.Exec(req.GetId())
	if err != nil {
		log.Printf("Error while deleting: %v", err)
	}
	log.Printf("Delete done with Result = %v\n doing commit now \n", res)

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Exiting Delete()\n")
	return nil
}
