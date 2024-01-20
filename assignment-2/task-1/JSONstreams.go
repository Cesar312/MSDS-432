package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

type Data struct {
	Key string `json:"key"`
	Val int    `json:"value"`
}

var serializedData []byte

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

func serialize(buf *bytes.Buffer, DataRecords []Data, elapsedSerialize *time.Duration) {
	// Start CPU profiling for serialization
	fSer, err := os.Create("cpu_serialize.prof")
	if err != nil {
		log.Fatal("Could not create CPU profile for serialization: ", err)
	}
	pprof.StartCPUProfile(fSer)

	startSerialize := time.Now() // Start time serialization

	encoder := json.NewEncoder(buf)
	err = Serialize(encoder, DataRecords)
	if err != nil {
		log.Fatal("Serialization error: ", err)
		return
	}

	*elapsedSerialize = time.Since(startSerialize) // End timing serialization
	serializedData = buf.Bytes()
	fmt.Print("After Serialize:", string(serializedData))

	pprof.StopCPUProfile() // Stop CPU profiling
	fSer.Close()
}

func deserialize(buf *bytes.Buffer, elapsedDeserialize *time.Duration) {

	// Start CPU profiling for deserialization
	fDes, err := os.Create("cpu_deserialize.prof")
	if err != nil {
		log.Fatal("Could not create CPU profile for deserialization: ", err)
	}
	pprof.StartCPUProfile(fDes)

	startDeserialize := time.Now() // Start time deserialization

	decoder := json.NewDecoder(buf)
	var temp []Data
	err = DeSerialize(decoder, &temp)
	if err != nil {
		log.Fatal("Deserialization error: ", err)
		return
	}

	*elapsedDeserialize = time.Since(startDeserialize) // End timing deserialization

	fmt.Println("After DeSerialize:")
	for index, value := range temp {
		fmt.Println(index, value)
	}

	pprof.StopCPUProfile() // Stop CPU profiling
	fDes.Close()
}

func main() {

	var DataRecords []Data

	// Create sample data
	var i int
	var t Data
	for i = 0; i < 100000; i++ {
		t = Data{
			Key: getString(5),
			Val: random(1, 100),
		}
		DataRecords = append(DataRecords, t)
	}

	var elapsedSerialize, elapsedDeserialize time.Duration

	// bytes.Buffer is both an io.Reader and io.Writer
	buf := new(bytes.Buffer)

	// Serialization with Memory Profiling
	serialize(buf, DataRecords, &elapsedSerialize)
	writeMemProfile("memory_profile_serialize.prof")

	// Reset buffer for Deserialization
	buf.Reset()
	buf.Write(serializedData)

	// Deserialization with Memory Profiling
	deserialize(buf, &elapsedDeserialize)
	writeMemProfile("memory_profile_deserialize.prof")

	fmt.Printf("Serialization took %s \n", elapsedSerialize)
	fmt.Printf("Deserialization took %s \n", elapsedDeserialize)
}

// Write memory profile
func writeMemProfile(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal("Could not create memory profile: ", err)
	}
	defer f.Close()

	runtime.GC()
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal("Could not write memory profile: ", err)
	}
}
