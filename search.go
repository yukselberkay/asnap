package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func readCsvCn(srcStr string, fp *os.File, outfile string) {
	f, err := os.Create(outfile + ".txt")
	checkErr(err)

	db := csv.NewReader(fp)
	for {
		db, err := db.Read()

		if err == io.EOF {
			break
		}

		matchedCn, err := regexp.MatchString(strings.ToLower(srcStr),
			strings.ToLower(db[2]))
		checkErr(err)

		if matchedCn == true {
			fmt.Println(string(ColorGreen), db[0], "AS"+db[1], db[2])
			_, err := f.WriteString(db[0] + "\n")
			checkErr(err)

		}
	}

	defer fp.Close()
}

func readCsvAsn(srcStr string, fp *os.File, outfile string) {
	f, err := os.Create(outfile + ".txt")
	checkErr(err)

	db := csv.NewReader(fp)
	for {
		db, err := db.Read()

		if err == io.EOF {
			break
		}

		matchedCn, err := regexp.MatchString(strings.ToLower(srcStr),
			strings.ToLower(db[1]))
		checkErr(err)

		if matchedCn == true {
			fmt.Println(string(ColorGreen), db[0], "AS"+db[1], db[2])
			_, err := f.WriteString(db[0] + "\n")
			checkErr(err)
		}
	}

	defer fp.Close()
}
func searchCn(srcStr string, mode string, outfile string) {
	//colorGreen := "\033[32m"

	if mode == "ipv4" {
		file, err := os.Open("GeoLite2-ASN-Blocks-IPv4.csv")
		checkErr(err)
		readCsvCn(srcStr, file, outfile)

	} else if mode == "ipv6" {
		file, err := os.Open("GeoLite2-ASN-Blocks-IPv6.csv")
		checkErr(err)
		readCsvCn(srcStr, file, outfile)

	}
}

func searchAsn(srcStr string, mode string, outfile string) {

	if mode == "ipv4" {
		file, err := os.Open("GeoLite2-ASN-Blocks-IPv4.csv")
		checkErr(err)
		readCsvAsn(srcStr, file, outfile)

	} else if mode == "ipv6" {
		file, err := os.Open("GeoLite2-ASN-Blocks-IPv6.csv")
		checkErr(err)
		readCsvAsn(srcStr, file, outfile)

	}
}
