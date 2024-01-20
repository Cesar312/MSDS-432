package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"
)

type Data struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

var DataRecords []Data

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

var MIN = 0
var MAX = 26

func getString(l int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1
	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == l {
			break
		}
		i++
	}
	return temp
}

// DeSerialize decodes a serialized slice with JSON records
func DeSerialize(e *json.Decoder, slice interface{}) error {
	return e.Decode(slice)
}

// Serialize serializes a slice with JSON records
func Serialize(e *json.Encoder, slice interface{}) error {
	return e.Encode(slice)
}

func main() {

	// Create sample data
	var i int
	var t Data
	for i = 0; i < 1000; i++ {
		t = Data{
			Key: getString(5),
			Val: random(1, 100),
		}
		DataRecords = append(DataRecords, t)
	}

	// bytes.Buffer is both an io.Reader and io.Writer
	buf := new(bytes.Buffer)

	// Serialization
	// Start CPU profiling for serialization
	fSer, err := os.Create("cpu_serialize.prof")
	if err != nil {
		log.Fatal("Could not create CPU profile for serialization: ", err)
	}
	pprof.StartCPUProfile(fSer)

	// Start time serialization
	startSerialize := time.Now()

	encoder := json.NewEncoder(buf)
	err = Serialize(encoder, DataRecords)
	if err != nil {
		fmt.Println("Serialization error: ", err)
		return
	}

	// End timing serialization
	elapsedSerialize := time.Since(startSerialize)
	fmt.Print("After Serialize:", buf)

	// Stop CPU profiling for serialization
	pprof.StopCPUProfile()
	fSer.Close()

	// Deserialization
	// Start CPU profiling for deserialization
	fDes, err := os.Create("cpu_deserialize.prof")
	if err != nil {
		log.Fatal("Could not create CPU profile for deserialization: ", err)
	}
	pprof.StartCPUProfile(fDes)

	// Start time deserialization
	startDeserialize := time.Now()

	decoder := json.NewDecoder(buf)
	var temp []Data
	err = DeSerialize(decoder, &temp)
	if err != nil {
		fmt.Println("Deserialization error: ", err)
		return
	}

	// End timing deserialization
	elapsedDeserialize := time.Since(startDeserialize)

	fmt.Println("After DeSerialize:")

	// Stop CPU profiling for  deserialization
	pprof.StopCPUProfile()
	fDes.Close()

	for index, value := range temp {
		fmt.Println(index, value)
	}

	fmt.Print("\n")
	fmt.Printf("Serialization took %s \n", elapsedSerialize)
	fmt.Printf("Deserialization took %s \n", elapsedDeserialize)
}
