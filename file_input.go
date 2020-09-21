package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func fileInput(inFile string, mode string, outFile string) {

	f1, err := os.Create(outFile + ".txt")
	checkErr(err)

	f2, err := os.Open(inFile)
	checkErr(err)

	scanner := bufio.NewScanner(f2)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	for _, eachLn := range text {
		if mode == "ipv4" {
			fp, err := os.Open("GeoLite2-ASN-Blocks-IPv4.csv")
			checkErr(err)
			db := csv.NewReader(fp)

			for {
				db, err := db.Read()

				if err == io.EOF {
					break
				}

				matchedCn, err := regexp.MatchString(strings.ToLower(eachLn),
					strings.ToLower(db[2]))
				checkErr(err)

				if matchedCn == true {
					fmt.Println(db[0], "AS"+db[1], db[2])
					_, err := f1.WriteString(db[0] + "\n")
					checkErr(err)
				}
			}
		} else if mode == "ipv6" {
			fp, err := os.Open("GeoLite2-ASN-Blocks-IPv6.csv")
			checkErr(err)
			db := csv.NewReader(fp)

			for {
				db, err := db.Read()

				if err == io.EOF {
					break
				}

				matchedCn, err := regexp.MatchString(strings.ToLower(eachLn),
					strings.ToLower(db[2]))
				checkErr(err)

				if matchedCn == true {
					fmt.Println(db[0], "AS"+db[1], db[2])
					_, err := f1.WriteString(db[0] + "\n")
					checkErr(err)
				}
			}
		}

	}

}
