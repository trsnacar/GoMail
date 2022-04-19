package main

import (
	"fmt"

	gomail "gopkg.in/mail.v2"
)

func main() {
	abc := gomail.NewMessage()

	abc.SetHeader("From", "trsn.acr@gmail.com")
	abc.SetHeader("To", "abc@gmail.com")
	abc.SetHeader("Subject", "Test Email Subject abc")
	abc.SetHeader("text/plain", "This is the Test Body")

	a := gomail.NewDialer("smtp.gmail.com", 587, "trsn.acr@gmail.com", "password")

	if err := a.DialAndSend(abc); err != nil {
		fmt.Println(err)
	}

}
