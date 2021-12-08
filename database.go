package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)


func init(){
	var URI string = "postgres://foudystf:0ZNypQVJNOhsxEq_yPyd_kW2Fz9Y3jKR@castor.db.elephantsql.com/foudystf"
	db, err := sql.Open("postgres", URI)
	if err != nil{
		log.Fatalf("Error while connecting to database: %s",err)
	}
	err = db.Ping()
	if err != nil{
		log.Fatalf("Error while connecting to database: %s",err)
	}
	fmt.Println("Successfully created connection to database")
}