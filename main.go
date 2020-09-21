package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	// default name for out file.
	currentTime := time.Now()
	now := currentTime.Format("01-02-2006")
	out := string(now) + "_out"

	// custom help message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Println("-download		Download database for the first usage.")
		fmt.Println("-update			Update downloaded database. (Geolite databases updates once a week.).")
		fmt.Println("-search			Specify search.")
		fmt.Println("-ipv4			Specify ipv4 database to search.")
		fmt.Println("-ipv6 			Specify ipv6 database to search.")
		fmt.Println("-company 		Search by company name.")
		fmt.Println("-asn			Search by as number.")
		fmt.Println("-outfile 		Specifies a name for the output text. By default, output file is named: MM-DD-YYYY_out.txt")
		fmt.Println("-infile			Use specified .txt file as input. Asnap will iterate every line, and treats them as company names and searches specified database with given inputs.")
		fmt.Println("-nmap			Passes found ip addresses to nmap.")
		fmt.Println()

		fmt.Println("Examples:")
		fmt.Println(`"$asnap -download" -> Downloads database with given key, for the first time.`)
		fmt.Println(`"$asnap -update" -> Updates database.`)
		fmt.Println(`"$asnap -search -ipv4 -company="example" " -> Search ipv4 database by company name "example"`)
		fmt.Println(`"$asnap -search -ipv6 -asn 13337" -> Search ipv6 database by as number "13337"`)
		fmt.Println(`"$asnap -search -ipv4 -company="github" -outfile /path/to/output/file" -> Search ipv4 database by company name "test" and save output to specified path.`)
		fmt.Println(`"$asnap -search -ipv4 -infile /path/to/input/file.txt -nmap" -> Give a list of company names as input, search it inside ipv4 database and pass found ip addresses to nmap for port scanning. `)
		fmt.Println()
	}

	fmt.Println(`
	
	█████╗ ███████╗███╗   ██╗ █████╗ ██████╗ 
	██╔══██╗██╔════╝████╗  ██║██╔══██╗██╔══██╗
	███████║███████╗██╔██╗ ██║███████║██████╔╝
	██╔══██║╚════██║██║╚██╗██║██╔══██║██╔═══╝ 
	██║  ██║███████║██║ ╚████║██║  ██║██║     
	╚═╝  ╚═╝╚══════╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     									
	Author : Mehmet Berkay Yuksel | Blog -> https://yukselberkay.me
												
	 `)
	dwPtr := flag.Bool("download", false, "Downloads Database.")
	upPtr := flag.Bool("update", false, "Updates Database.")
	srcPtr := flag.Bool("search", false, "Specifies search.")
	ipv4Ptr := flag.Bool("ipv4", false, "Search for ipv4.")
	ipv6Ptr := flag.Bool("ipv6", false, "Search for ipv6.")
	asnPtr := flag.String("asn", "", "Search database ASN.")
	cnPtr := flag.String("company", "", "Search database by company name")
	inPtr := flag.String("infile", "", "text file as a input.")
	nmPtr := flag.Bool("nmap", false, "pass the results to nmap")

	outPtr := flag.String("outfile", out, "outfile")

	flag.Parse()

	fname := "db.zip"

	//arguments and function calls
	if *dwPtr == true {
		download(fname)

		files, err := unzip(fname, "extracted_files")
		checkErr(err)
		fmt.Println("[+] Unzipped following files : \n" + strings.Join(files, "\n"))

		move()
	} else if *upPtr == true {
		update(fname)

		files, err := unzip(fname, "extracted_files")
		checkErr(err)
		fmt.Println("[+] Unzipped following files : \n" + strings.Join(files, "\n"))

		move()

	} else if *srcPtr == true && *ipv4Ptr == true {
		if *cnPtr != "" {
			searchCn(*cnPtr, "ipv4", *outPtr)
		} else if *asnPtr != "" {
			searchAsn(*asnPtr, "ipv4", *outPtr)
		}
	} else if *srcPtr == true && *ipv6Ptr == true {
		if *cnPtr != "" {
			searchCn(*cnPtr, "ipv6", *outPtr)
		} else if *asnPtr != "" {
			searchAsn(*asnPtr, "ipv6", *outPtr)
		}
	} else if *inPtr != "" && *ipv4Ptr == true {
		fileInput(*inPtr, "ipv4", *outPtr)

	} else if *inPtr != "" && *ipv6Ptr == true {
		fileInput(*inPtr, "ipv6", *outPtr)
	} else if len(flag.Args()) == 0 {
		flag.Usage()
	}
	if *nmPtr == true {
		nmap(out + ".txt")
	}
}
