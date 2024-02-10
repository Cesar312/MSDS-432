package main

import (
	"fmt"

	"github.com/Cesar312/MSDS-432/assignment-5/catalog/post05-main"
)

func main() {
	post05.Hostname = "localhost"
	post05.Port = 5433
	post05.Username = "postgres"
	post05.Password = "root"
	post05.Database = "msds"

	// Define and add MSDS courses
	courses := []post05.MSDSCourse{
		{CID: "MSDS401", CNAME: "Applied Statistics with R", CPREREQ: "MSDS400"},
		{CID: "MSDS420", CNAME: "Database Systems", CPREREQ: "MSDS402"},
		{CID: "MSDS432", CNAME: "Foundations of Data Engineering", CPREREQ: "MSDS420"},
		{CID: "MSDS453", CNAME: "Natural Language Processing", CPREREQ: "MSDS422"},
		{CID: "MSDS498", CNAME: "Capstone Project", CPREREQ: "MSDS453"},
	}

	for _, course := range courses {
		err := post05.AddCourse(course)
		if err != nil {
			fmt.Printf("Error adding course %s: %v\n", course.CID, err)
		} else {
			fmt.Printf("Successfully added course %s\n", course.CID)
		}
	}

}
