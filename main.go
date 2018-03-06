package main

import (
	"fmt"
	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
	"os"
)

const version = "1.1.1"

func doSelfUpdate() {

	latest, found, err := selfupdate.DetectLatest("maki5/go_test")
	if err != nil {
		fmt.Println("Error occurred while detecting version:", err)
		return
	}

	v := semver.MustParse(version)
	if !found || latest.Version.Equals(v) {
		fmt.Println("Current version is the latest")
		return
	}

	if err := selfupdate.UpdateTo(latest.AssetURL, os.Args[0]); err != nil {
		fmt.Println("Error occurred while updating binary:", err)
		return
	}
	fmt.Println("Successfully updated to version", latest.Version)

}

func main(){
	fmt.Println("Hell World")
	fmt.Println("current version: ", version)

	doSelfUpdate()
}
