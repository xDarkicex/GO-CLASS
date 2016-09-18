package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var goHTML *template.Template

func init() {
	goHTML = template.Must(template.ParseFiles("index.gohtml"))
}

type lineItem struct {
	Date     time.Time
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   int64
	Adjclose float64
}

func handler(res http.ResponseWriter, req *http.Request) {
	fmt.Println(req.RequestURI)
	if len(req.RequestURI) == 1 {
		root(res)
	} else {
		//
	}
}
func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
func root(res http.ResponseWriter) {

	file, err := os.Open("main.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rdr := csv.NewReader(file)
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}
	lineItems := make([]lineItem, 0, len(rows))
	for i, row := range rows {
		if i == 0 {
			continue
		}

		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)
		high, _ := strconv.ParseFloat(row[2], 64)
		low, _ := strconv.ParseFloat(row[3], 64)
		close, _ := strconv.ParseFloat(row[4], 64)
		volume, _ := strconv.ParseInt(row[5], 10, 64)
		adjclose, _ := strconv.ParseFloat(row[6], 64)

		lineItems = append(lineItems, lineItem{
			Date:     date,
			Open:     open,
			High:     high,
			Low:      low,
			Close:    close,
			Volume:   volume,
			Adjclose: adjclose,
		})
	}
	err2 := goHTML.Execute(res, lineItems)
	if err2 != nil {
		log.Fatalln(err)
	}
}
