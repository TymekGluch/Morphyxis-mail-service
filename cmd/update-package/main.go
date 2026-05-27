package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	log.Print("Starting to Upload new version of Morphyxis Mail Service ...")

	packageToken := os.Getenv("MORPHYXIS_MAIL_SERVICE_PACKAGE_TOKEN")
	if packageToken == "" {
		log.Print("Environment variable MORPHYXIS_MAIL_SERVICE_PACKAGE_TOKEN not set")
		os.Exit(1)
	}

	gitTag, err := exec.Command("git", "describe", "--tags", "--abbrev=0").Output()
	if err != nil || gitTag == nil {
		log.Print("Failed to get Git tag")
		os.Exit(1)
	}

	isBetaVersion := os.Getenv("IS_BETA_VERSION")
	if isBetaVersion == "" {
		log.Print("Environment variable IS_BETA_VERSION not set")
		os.Exit(1)
	}

	log.Printf("Run ci with variables: ")
	log.Printf("MORPHYXIS_MAIL_SERVICE_PACKAGE_TOKEN: %s", "[REDACTED]")
	log.Printf("GIT_TAG: %s", gitTag)
	log.Printf("IS_BETA_VERSION: %s", isBetaVersion)

	log.Print("Finished uploading new version of Morphyxis Mail Service")
}
