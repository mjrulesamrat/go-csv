package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"fmt"
	"log"
	"encoding/json"
)

type Product struct{
	Name string `json:"name"`
	Price float64 `json:"price"`
}


func main() {
	// loading a csv file
	file := "sample.csv"
	log.Println("Opening file...")

	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	csvr := csv.NewReader(f)
	csvr.Comma = ','
	csvr.Comment = '#'

	log.Println("Will Read file...")
	counter := 1
	products := []Product{}
	for {
		row, err := csvr.Read()
		
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		
		price, _ := strconv.ParseFloat(row[1], 64)
		p := Product{string(row[0]), price}
		products = append(products, p)
		counter++
	}

	// encoding products slice to json object so later we can store that data to mongodb
	b, err := json.Marshal(products)
	fmt.Println(string(b))
}