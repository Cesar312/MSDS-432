package main

////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////
// The following program will read data for
// Traffic Crashes from the City of Chicago data portal
// we are reading the data from CSV file and storing it in a map data structure.
////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

// The following is a sample record from the Traffic Crashes dataset retrieved from the City of Chicago Data Portal

////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

//Crash_record_id    2210669e39d023a283a253efbf227d83d5aa6a3e4af83277cfa86fbe27a972e07e0f2ca474bcc1043fdde3cf040281292f077a5a03bbde580b90aef73cc7c747
//Crash_date_est_i
//Crash_date    10/22/2021 20:57
//Posted_speed_limit    30
//Traffic_control_device    TRAFFIC SIGNAL
//Device_condition    FUNCTIONING PROPERLY
//Weather_condition    CLEAR
//Lighting_condition    DARKNESS, LIGHTED ROAD
//First_crash_type    ANGLE
//Trafficway_type    FOUR WAY
//Lane_cnt
//Alignment    STRAIGHT AND LEVEL
//Roadway_surface_cond    DRY
//Road_defect    NO DEFECTS
//Report_type    ON SCENE
//Crash_type    INJURY AND / OR TOW DUE TO CRASH
//Intersection_related_i
//Not_right_of_way_i
//Hit_and_run_i    Y
//Damage    OVER $1,500
//Date_police_notified    10/22/2021 20:57
//Prim_contributory_cause    DISREGARDING TRAFFIC SIGNALS
//Sec_contributory_cause    DISREGARDING TRAFFIC SIGNALS
//Street_no    2
//Street_direction    E
//Street_name    MARQUETTE RD
//Beat_of_occurrence    322
//Photos_taken_i
//Statements_taken_i    Y
//Dooring_i
//Work_zone_i
//Work_zone_type
//Workers_present_i
//Num_units    3
//Most_severe_injury    REPORTED, NOT EVIDENT
//Injuries_total    2
//Injuries_fatal    0
//Injuries_incapacitating    0
//Injuries_non_incapacitating    0
//Injuries_reported_not_evident    2
//Injuries_no_indication    3
//Injuries_unknown    0
//Crash_hour    20
//Crash_day_of_week    6
//Crash_month    10
//Latitude    41.77283498
//Longitude    -87.62507028
//Location    POINT (-87.625070277068 41.772834981744)
////////////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////////////

import (
	"encoding/csv"
	"encoding/json"
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
	Day_of_the_week               string
}

func main() {

	// import API_KEY
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file could not be loaded")
	}

	apiKey := os.Getenv("API_KEY")
	// fmt.Println(apiKey)

	// // open  csv dataset
	// dataset, err := os.Open("Traffic_Crashes_20240113.csv")
	dataset, err := os.Open("Traffic_Crashes_Subset.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer dataset.Close()

	// read crash data using csv.Reader
	csvReader := csv.NewReader(dataset)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// converts data to map with crashID as key and CrashData struct as values
	createCrashMap(data, apiKey)

	//uncomment below lines to print the entire map of data
	// crashMap := createCrashMap(data, apiKey)
	//fmt.Printf("\n\n\n")

	// prints the first entry in Hash Map of Crashes
	// fmt.Printf("%+v\n", crashMap[data[1][0]])

}

func createCrashMap(data [][]string, apiKey string) map[string]CrashData {

	crashMap := make(map[string]CrashData)
	crash_on_day_mp := make(map[string]int)
	crash_in_zip_code_mp := make(map[string]int)
	zip_code_with_Hit_and_run_2020_mp := make(map[string]int) // Specifically for 2020 Hit & Runs

	fmt.Println("CreateCrashMap: Creating Crash Map from Data")

	// Set the API key for the geocoder
	geocoder.ApiKey = apiKey

	// uncomment below line to process the entire data set
	// for i := 1; i < len(data); i++ {
	// uncomment below line to process 500 records only
	for i := 1; i < 5; i++ {

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

		// Processing the dataset to print Crashes on Each day of the week in 2021
		parsed_date, err := time.Parse("01/02/2006 03:04:05 PM", crashRecord.Crash_date)
		if err != nil {
			panic(err)
		}

		if parsed_date.Year() == 2021 {
			day_of_week := parsed_date.Weekday().String()
			crashRecord.Day_of_the_week = day_of_week
			crash_on_day_mp[day_of_week]++
		}

		// Processing the dataset to print Crashes at each Zip code
		if crashRecord.Zipcode != "" && crash_in_zip_code_mp[crashRecord.Zipcode] != 0 {
			crash_in_zip_code_mp[crashRecord.Zipcode] = crash_in_zip_code_mp[crashRecord.Zipcode] + 1
		} else if crashRecord.Zipcode != "" {
			crash_in_zip_code_mp[crashRecord.Zipcode] = 1
		}

		// Processing the dataset to print Hit and Runs at Each Zip code in 2020
		if crashRecord.Hit_and_run_i == "Y" && parsed_date.Year() == 2020 {
			if zip_code := crashRecord.Zipcode; zip_code != "" {
				zip_code_with_Hit_and_run_2020_mp[zip_code]++
			}
		}

		crashMap[data[i][0]] = crashRecord

	}

	fmt.Println("---------------------Crashes on Each day of the week in Year 2021-----------------")
	crash_on_day, err := json.MarshalIndent(crash_on_day_mp, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(crash_on_day))
	fmt.Printf("\n\n")

	fmt.Println("---------------------Crashes at Each Zip code-----------------")
	crash_in_zip_code, err := json.MarshalIndent(crash_in_zip_code_mp, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(crash_in_zip_code))
	fmt.Printf("\n\n")

	fmt.Println("---------------------Hit and Runs at Each Zip code in Year 2020-----------------")
	zip_code_with_Hit_and_run, err := json.MarshalIndent(zip_code_with_Hit_and_run_2020_mp, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(zip_code_with_Hit_and_run))
	fmt.Printf("\n\n")

	return crashMap
}
