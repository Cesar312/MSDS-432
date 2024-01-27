package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelvins/geocoder"
)

type CrashData struct {
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

func main() {

	// import API_KEY
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file could not be loaded")
	}

	apiKey := os.Getenv("API_KEY")

	// Open  CSV dataset
	dataset, err := os.Open("Traffic_Crashes_Subset.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer dataset.Close()

	// Read crash data using csv.Reader
	csvReader := csv.NewReader(dataset)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	createCrashMap := func(data [][]string, apiKey string) map[string]CrashData {

		crashMap := make(map[string]CrashData)

		fmt.Println("CreateCrashMap: Creating Crash Map from Data")

		// Set the API key for the geocoder
		geocoder.ApiKey = apiKey

		// uncomment below line to process the entire data set
		// for i := 0; i < len(data); i++ {
		// uncomment below line to process 10 records only
		for i := 1; i < 11; i++ {

			// Initilze the list
			var crashRecord CrashData

			crashRecord.Crash_date_est_i = data[i][1]
			crashRecord.Crash_date = data[i][2]
			crashRecord.Posted_speed_limit = data[i][3]
			crashRecord.Traffic_control_device = data[i][4]
			crashRecord.Device_condition = data[i][5]
			crashRecord.Weather_condition = data[i][6]
			crashRecord.Lighting_condition = data[i][7]
			crashRecord.First_crash_type = data[i][8]
			crashRecord.Trafficway_type = data[i][9]
			crashRecord.Lane_cnt = data[i][10]
			crashRecord.Alignment = data[i][11]
			crashRecord.Roadway_surface_cond = data[i][12]
			crashRecord.Road_defect = data[i][13]
			crashRecord.Report_type = data[i][14]
			crashRecord.Crash_type = data[i][15]
			crashRecord.Intersection_related_i = data[i][16]
			crashRecord.Not_right_of_way_i = data[i][17]
			crashRecord.Hit_and_run_i = data[i][18]
			crashRecord.Damage = data[i][19]
			crashRecord.Date_police_notified = data[i][20]
			crashRecord.Prim_contributory_cause = data[i][21]
			crashRecord.Sec_contributory_cause = data[i][22]
			crashRecord.Street_no = data[i][23]
			crashRecord.Street_direction = data[i][24]
			crashRecord.Street_name = data[i][25]
			crashRecord.Beat_of_occurrence = data[i][26]
			crashRecord.Photos_taken_i = data[i][27]
			crashRecord.Statements_taken_i = data[i][28]
			crashRecord.Dooring_i = data[i][29]
			crashRecord.Work_zone_i = data[i][30]
			crashRecord.Work_zone_type = data[i][31]
			crashRecord.Workers_present_i = data[i][32]
			crashRecord.Num_units = data[i][33]
			crashRecord.Most_severe_injury = data[i][34]
			crashRecord.Injuries_total = data[i][35]
			crashRecord.Injuries_fatal = data[i][36]
			crashRecord.Injuries_incapacitating = data[i][37]
			crashRecord.Injuries_non_incapacitating = data[i][38]
			crashRecord.Injuries_reported_not_evident = data[i][39]
			crashRecord.Injuries_no_indication = data[i][40]
			crashRecord.Injuries_unknown = data[i][41]
			crashRecord.Crash_hour = data[i][42]
			crashRecord.Crash_day_of_week = data[i][43]
			crashRecord.Crash_month = data[i][44]

			if data[i][45] == "" || data[i][46] == "" {
				continue
			} else {
				crashRecord.Latitude = data[i][45]
				crashRecord.Longitude = data[i][46]
				crashRecord.Location = data[i][47]
			}

			// Using latitude and longitude in geocoder.GeocodingReverse
			// we could find the crash Zip-Code and Address

			latitude_float, _ := strconv.ParseFloat(crashRecord.Latitude, 64)
			longitude_float, _ := strconv.ParseFloat(crashRecord.Longitude, 64)

			location := geocoder.Location{
				Latitude:  latitude_float,
				Longitude: longitude_float,
			}

			address_list, _ := geocoder.GeocodingReverse(location)

			// Ignoring the entry if location of the data is not available
			if len(address_list) == 0 {
				fmt.Printf("No results found for crash at latitude : %f and Longitude : %f \n", latitude_float, longitude_float)
				continue
			}

			address := address_list[0]
			zip_code := address.PostalCode

			crashRecord.Zipcode = zip_code
			crashRecord.Address = address

			fmt.Printf("Record ID: %s, Zip Code: %s\n", data[i][0], crashRecord.Zipcode)

			crashMap[data[i][0]] = crashRecord

		}

		return crashMap
	}

	// Start time
	startTime := time.Now()

	crashMap := createCrashMap(data, apiKey)

	// Elapsed time
	elapsedTime := time.Since(startTime)

	fmt.Println("Number of records processed: ", len(crashMap))
	fmt.Println("Elapsed time: ", elapsedTime)

}
