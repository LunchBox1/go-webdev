package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

/*
-- encoding/csv
 - package "encoding" defines interfaces shared by other packages that conver data to and from byte-level and textual representations.
 - package "csv" reads and writes comma-separated values (CSV) files.
 (encoding methods are not concurrency-safe)
 (an interface is a custom type that is used to specify a set of one or more method signatures)
 - package "time" provides functionality for measuring and displaying time.
 - package "os" provides a platform-independant interface to operating system functionality.
*/

type Record struct {
	Date time.Time //type Time from package time. Programs using time should typically store and pass them as values, not pointers.
	//That is, time variables and struct fields should be of type time.Time, not *time.Time.
	//Date and Open are the two fields that we are pulling out of the CSV file and placing into a data structure. So to specify which fields we want to pull out, we can create a struct that will hold that data.
	Open float64
}

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}
func foo(res http.ResponseWriter, req *http.Request) {
	//parse csv
	records := prs("table.csv")

	//parse template
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	//execute template
	err = tpl.Execute(res, records)
	if err != nil {
		log.Fatalln(err)
	}
}
func prs(filePath string) []Record { //the csv file will be parsed here. Two fields from the csv will be placed into a data structure. Function prs takes in a string and returns a []Record.
	src, err := os.Open(filePath) //func Open, from package os, opens the named file for reading. It takes in a string and returns a *File (File is a type from os, I believe a location for the file is created in
	if err != nil {               //in memmory), and an error. We assign the returned values to identifiers src and err.
		log.Fatalln(err)
	}
	defer src.Close() //src.Close() closed the open file. "defer" will stop the closing until all of the code has finished running and close the file at the end of excecution.

	rdr := csv.NewReader(src) //func NewReader returns a new Reader that reads from r (r is of type io.Reader and is taken in by the NewReader func). A Reader reads records from a CSV-encoded file.
	// func NewReader(r io.Reader) *Reader; rdr is a value of type *Reader.
	rows, err := rdr.ReadAll() //ReadAll is a method, from package csv, that is attached to any value of type *Reader. func (r *Reader) ReadAll() (records [][]string, err error)
	if err != nil {            //ReadAll reads all the remaining records from r. Each record is a slice of fields(cells). I belive that ReadAll reads a row at a time from the CSV file.
		log.Fatalln(err)
	}

	records := make([]Record, 0, len(rows)) //making an empty []Record who's length is set to the amount of rows in the CSV file.

	for i, row := range rows {
		if i == 0 { //this skips the first row in the CSV file. The first row is the names of the columns, so essencially, we are throwing away that row.
			continue //i is the index position and it is evaluated at index postion 0. Since the evaluation is true (i == 0), "continue" returns to the beggining of the for loop, this time at index position 1
		}

		date, _ := time.Parse("2006-01-02", row[0]) //func Parse(layout, value string) (Time, error); Parse parses a formatted string and returns the time value it represents. The second argument (value string),
		//must be parseable using the format sting (layout) provided as the first argument. date is of type Time that is returned by the time.Parse function. This func
		//takes the value that is in the index positon 0 of row (row[0]), and formats it to what is specified in the first argument. "row[0]" is a good representation
		//of a column.
		open, _ := strconv.ParseFloat(row[1], 64) //package strconv implements conversions to and from string representations of basic data types. ParseFloat converts the string s to a floating-point numbers
		//with the precision specified by bitSize: 32 for float32, or 64 for float64. func ParseFloat(s string, bitSize int) (float54, error). This converts the sring in row[1] to a float64.

		records = append(records, Record{
			Date: date,
			Open: open,
		})
	}
	return records
}
