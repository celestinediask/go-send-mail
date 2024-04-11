package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func createConfig() {

	// Create .env file

	filename := ".env"

	// Check if file already exists
	if _, err := os.Stat(filename); err == nil {
		fmt.Println("File " + filename + " found")
		return
	} else if !os.IsNotExist(err) {
		fmt.Println("Error file "+filename+" checking file status:", err)
		return
	}

	// Create or open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating "+filename+"file:", err)
		return
	}
	defer file.Close()

	// Write lines to the file
	lines := []string{
		"EMAIL_FROM=",
		"EMAIL_PASSWORD=",
		"EMAIL_TO=",
		"EMAIL_SUB=",
		"EMAIL_MSG=",
	}
	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("File " + filename + " created successfully. Set appropriate values in this file and run again")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func sendMailSimple(email_to, sub, msg string) {
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")

	addr := "smtp.gmail.com:587"
	to := []string{email_to}

	smsg := "Subject: " + sub + "\n" + msg

	auth := smtp.PlainAuth(
		"", // Identity
		from,
		password,         // password
		"smtp.gmail.com", // host
	)

	err := smtp.SendMail(addr, auth, from, to, []byte(smsg))

	if err != nil {
		fmt.Println("Error: sending email", err)
	} else {
		fmt.Println("Successfully send email!")
	}
}

func main() {

	createConfig()

	//load .env file
	loadEnv()

	// Email data
	to := os.Getenv("EMAIL_TO")
	sub := os.Getenv("EMAIL_SUB")
	msg := os.Getenv("EMAIL_MSG")

	// Send email
	sendMailSimple(to, sub, msg)

}
