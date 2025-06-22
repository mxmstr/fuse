package main

import (
	"fuse/sessionmanager"
	"os"
)

func ttt() {
	ss, _ := os.ReadFile("/tmp/soldiers.base64")
	sessionmanager.ReadSoldierData(string(ss))
}

func main() {
	ttt()
}
