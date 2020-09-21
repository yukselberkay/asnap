package main

// url : https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-ASN-CSV&license_key=<key>&suffix=zip

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func download(fname string) {

	if _, err := os.Stat(fname); !os.IsNotExist(err) {
		fmt.Println("[-] Database already downloaded.")
		fmt.Println("Geolite databases updates once a week. To update your current database use -update argument.")
		os.Exit(1)
	}

	out, err := os.Create(string(fname))
	checkErr(err)
	defer out.Close()

	url := fmt.Sprintf("https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-ASN-CSV&license_key=%s&suffix=zip", conf())
	fmt.Println("GET -> " + url)
	resp, err := http.Get(url)
	checkErr(err)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("[-] Bad status : %v\n", resp.Status)
		os.Exit(1)
	} else if resp.StatusCode == http.StatusOK {
		fmt.Printf(`[+]File "%s" downloaded successfully.`, fname)
		fmt.Println()
	}

	_, err = io.Copy(out, resp.Body)
	checkErr(err)

}

func update(fname string) {
	cmd := exec.Command("rm", fname)
	err := cmd.Run()
	checkErr(err)

	download(fname)
}
