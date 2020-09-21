package main

import "os/exec"

func move() {

	cmd := exec.Command("./move.sh")
	err := cmd.Run()
	checkErr(err)
}
