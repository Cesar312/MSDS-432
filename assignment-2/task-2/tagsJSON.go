package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type MSDSCourse struct {
	CID     string `json:"course_id"`
	CNAME   string `json:"course_name"`
	CPREREQ string `json:"prerequisite,omitempty"`
}

func main() {

	// Array of courses
	var coursesArray [5]MSDSCourse
	coursesArray[0] = MSDSCourse{CID: "MSDS-400", CNAME: "Math for Modelers", CPREREQ: ""}
	coursesArray[1] = MSDSCourse{CID: "MSDS-422", CNAME: "Practical Machine Learning", CPREREQ: "MSDS-430"}
	coursesArray[2] = MSDSCourse{CID: "MSDS-432", CNAME: "Foundations of Data Engineering", CPREREQ: "MSDS-420"}
	coursesArray[3] = MSDSCourse{CID: "MSDS-460", CNAME: "Decision Analytics", CPREREQ: "MSDS-400"}
	coursesArray[4] = MSDSCourse{CID: "MSDS-498", CNAME: "AI Capstone Project", CPREREQ: "MSDS-458"}

	// Slice of courses
	coursesSlice := []MSDSCourse{
		{CID: "MSDS-400", CNAME: "Math for Modelers", CPREREQ: ""},
		{CID: "MSDS-422", CNAME: "Practical Machine Learning", CPREREQ: "MSDS-430"},
		{CID: "MSDS-432", CNAME: "Foundations of Data Engineering", CPREREQ: "MSDS-420"},
		{CID: "MSDS-460", CNAME: "Decision Analytics", CPREREQ: "MSDS-400"},
		{CID: "MSDS-498", CNAME: "AI Capstone Project", CPREREQ: "MSDS-458"},
	}

	// Map of courses
	coursesMap := map[string]MSDSCourse{
		"MSDS-400": {CID: "MSDS-400", CNAME: "Math for Modelers", CPREREQ: ""},
		"MSDS-422": {CID: "MSDS-422", CNAME: "Practical Machine Learning", CPREREQ: "MSDS-430"},
		"MSDS-432": {CID: "MSDS-432", CNAME: "Foundations of Data Engineering", CPREREQ: "MSDS-420"},
		"MSDS-460": {CID: "MSDS-460", CNAME: "Decision Analytics", CPREREQ: "MSDS-400"},
		"MSDS-498": {CID: "MSDS-498", CNAME: "AI Capstone Project", CPREREQ: "MSDS-458"},
	}

	// Marshal and print the Array to JSON
	printJSON("Array", coursesArray[:])

	// Marshal and print the Slice to JSON
	printJSON("Slice", coursesSlice)

	// Marshal and print the Map to JSON
	printJSON("Map", coursesMap)

}

func printJSON(dataType string, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling %s: %v", dataType, err)
	}
	fmt.Printf("%s: %s\n\n", dataType, jsonData)
}
