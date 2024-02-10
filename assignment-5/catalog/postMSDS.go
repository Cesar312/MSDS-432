package main

import (
	"fmt"
	// "math/rand"
	//"time"

	"github.com/mactsouk/post05"
)

func main() {
	post05.Hostname = "localhost"
	post05.Port = 5433
	post05.Username = "postgres"
	post05.Password = "root"
	post05.Database = "msds"

	// data, err := post05.ListCourses()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// for _, c := range data {
	// 	fmt.Println(c)
	// }

	// // Seed for random generation
	// SEED := time.Now().Unix()
	// rand.Seed(SEED)

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
