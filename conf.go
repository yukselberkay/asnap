package main

import (
	"io/ioutil"
	"strings"
)

func conf() string {
	key, err := ioutil.ReadFile("asnap_conf.txt")

	keyStr := string(key)

	keyStr = strings.TrimSuffix(keyStr, "\n")

	checkErr(err)

	return keyStr

}
