package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelvins/geocoder"
)

type CrashData struct {
	Crash_record_id               string
	Crash_date_est_i              string
	Crash_date                    string
	Posted_speed_limit            string
	Traffic_control_device        string
	Device_condition              string
	Weather_condition             string
	Lighting_condition            string
	First_crash_type              string
	Trafficway_type               string
	Lane_cnt                      string
	Alignment                     string
	Roadway_surface_cond          string
	Road_defect                   string
	Report_type                   string
	Crash_type                    string
	Intersection_related_i        string
	Not_right_of_way_i            string
	Hit_and_run_i                 string
	Damage                        string
	Date_police_notified          string
	Prim_contributory_cause       string
	Sec_contributory_cause        string
	Street_no                     string
	Street_direction              string
	Street_name                   string
	Beat_of_occurrence            string
	Photos_taken_i                string
	Statements_taken_i            string
	Dooring_i                     string
	Work_zone_i                   string
	Work_zone_type                string
	Workers_present_i             string
	Num_units                     string
	Most_severe_injury            string
	Injuries_total                string
	Injuries_fatal                string
	Injuries_incapacitating       string
	Injuries_non_incapacitating   string
	Injuries_reported_not_evident string
	Injuries_no_indication        string
	Injuries_unknown              string
	Crash_hour                    string
	Crash_day_of_week             string
	Crash_month                   string
	Latitude                      string
	Longitude                     string
	Location                      string
	Zipcode                       string
	Address                       geocoder.Address
}

const workerCount = 10 // Number of concurrent workers

func main() {

	// Import API_KEY
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file could not be loaded")
	}

	fmt.Println("Successfully loaded API key")

	apiKey := os.Getenv("API_KEY")

	// Open CSV dataset
	dataset, err := os.Open("Traffic_Crashes_Subset.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer dataset.Close()

	csvReader := csv.NewReader(dataset)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Start time
	startTime := time.Now()

	crashMap := processCrashData(data, apiKey)

	// Elapsed time
	elapsedTime := time.Since(startTime)

	fmt.Println("Number of crash records processed: ", len(crashMap))
	fmt.Println("Elapsed time: ", elapsedTime)
}

func processCrashData(data [][]string, apiKey string) map[string]CrashData {
	crashMap := make(map[string]CrashData)

	// Channels for passing data between stages
	records := make(chan []string)
	processed := make(chan CrashData)

	// Start goroutines
	var wg sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for record := range records {
				crashData, err := processRecord(record, apiKey)
				if err != nil {
					log.Println("Error processing record: ", err)
					continue
				}
				processed <- crashData
			}
		}()
	}

	// Goroutine to close the processed channel when workers are done
	go func() {
		wg.Wait()
		close(processed)
	}()

	// Send records to workers
	go func() {
		for _, record := range data {
			// fmt.Println("Sending record to worker")
			records <- record
		}
		close(records)
	}()

	// Collect processed records
	for crashData := range processed {
		recordID := crashData.Crash_record_id
		zipCode := crashData.Zipcode
		fmt.Printf("Record ID: %s, Zip Code: %s\n", recordID, zipCode)
		crashMap[recordID] = crashData
	}

	return crashMap
}

func processRecord(record []string, apiKey string) (CrashData, error) {

	// Set the API key for the geocoder
	geocoder.ApiKey = apiKey

	var crashData CrashData

	crashData.Crash_record_id = record[0]
	crashData.Crash_date = record[1]
	crashData.Crash_date_est_i = record[2]
	crashData.Posted_speed_limit = record[3]
	crashData.Traffic_control_device = record[4]
	crashData.Device_condition = record[5]
	crashData.Weather_condition = record[6]
	crashData.Lighting_condition = record[7]
	crashData.First_crash_type = record[8]
	crashData.Trafficway_type = record[9]
	crashData.Lane_cnt = record[10]
	crashData.Alignment = record[11]
	crashData.Roadway_surface_cond = record[12]
	crashData.Road_defect = record[13]
	crashData.Report_type = record[14]
	crashData.Crash_type = record[15]
	crashData.Intersection_related_i = record[16]
	crashData.Not_right_of_way_i = record[17]
	crashData.Hit_and_run_i = record[18]
	crashData.Damage = record[19]
	crashData.Date_police_notified = record[20]
	crashData.Prim_contributory_cause = record[21]
	crashData.Sec_contributory_cause = record[22]
	crashData.Street_no = record[23]
	crashData.Street_direction = record[24]
	crashData.Street_name = record[25]
	crashData.Beat_of_occurrence = record[26]
	crashData.Photos_taken_i = record[27]
	crashData.Statements_taken_i = record[28]
	crashData.Dooring_i = record[29]
	crashData.Work_zone_i = record[30]
	crashData.Work_zone_type = record[31]
	crashData.Workers_present_i = record[32]
	crashData.Num_units = record[33]
	crashData.Most_severe_injury = record[34]
	crashData.Injuries_total = record[35]
	crashData.Injuries_fatal = record[36]
	crashData.Injuries_incapacitating = record[37]
	crashData.Injuries_non_incapacitating = record[38]
	crashData.Injuries_reported_not_evident = record[39]
	crashData.Injuries_no_indication = record[40]
	crashData.Injuries_unknown = record[41]
	crashData.Crash_hour = record[42]
	crashData.Crash_day_of_week = record[43]
	crashData.Crash_month = record[44]
	crashData.Latitude = record[45]
	crashData.Longitude = record[46]
	crashData.Location = record[47]

	// Using latitude and longitude in geocoder.GeocodingReverse
	// we could find the crash Zip-Code and Address
	latitude_float, err := strconv.ParseFloat(record[45], 64)
	if err != nil {
		return CrashData{}, fmt.Errorf("invalid latitude: %s", record[45])
	}

	longitude_float, err := strconv.ParseFloat(record[46], 64)
	if err != nil {
		return CrashData{}, fmt.Errorf("invalid longitude: %s", record[46])
	}

	// Gecoding using latitude and longitude
	location := geocoder.Location{
		Latitude:  latitude_float,
		Longitude: longitude_float,
	}

	address_list, err := geocoder.GeocodingReverse(location)
	if err != nil {
		return CrashData{}, err
	}

	// Ignoring the entry if location of the data is not available
	if len(address_list) > 0 {
		crashData.Zipcode = address_list[0].PostalCode
		crashData.Address = address_list[0]
	}

	return crashData, nil

}
