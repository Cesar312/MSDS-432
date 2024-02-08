package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kelvins/geocoder"
)

type TaxiTrip struct {
	TripID     string `json:"trip_id"`
	TaxiID     string `json:"taxi_id"`
	PickupLat  string `json:"pickup_centroid_latitude"`
	PickupLong string `json:"pickup_centroid_longitude"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		panic("API key not set in .env file")
	}

	url := "https://data.cityofchicago.org/resource/wrvz-psew.json?$limit=2"
	getResponse, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer getResponse.Body.Close()

	var TaxiTrips []TaxiTrip
	json.NewDecoder(getResponse.Body).Decode(&TaxiTrips)

	geocoder.ApiKey = apiKey

	// Print the addresses
	for _, trip := range TaxiTrips {
		latitude_float, _ := strconv.ParseFloat(trip.PickupLat, 64)
		longitude_float, _ := strconv.ParseFloat(trip.PickupLong, 64)

		location := geocoder.Location{
			Latitude:  latitude_float,
			Longitude: longitude_float,
		}

		address_list, _ := geocoder.GeocodingReverse(location)
		if err != nil {
			fmt.Println("Error fetching address: ", err)
			continue
		}

		if len(address_list) > 0 {
			address := address_list[0].FormattedAddress
			fmt.Printf("Trip ID: %s\nLatitude: %f\tLongitude: %f\tAddress: %s\n", trip.TripID, latitude_float, longitude_float, address)
		} else {
			fmt.Println("No address found for the given coordinates")
		}
	}
}
