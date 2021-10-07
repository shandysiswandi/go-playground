package main

import (
	"fmt"

	gomail "gopkg.in/mail.v2"
)

var (
	smtpHost = "smtp.mailtrap.io"
	smtpPort = 2525 // 25(error localhost) or 465(error localhost) or 587 (tls) or 2525 (tls)
	smtpUser = "42239792fa9ea4"
	smtpPass = "486010d4347713"
)

func main() {
	m := gomail.NewMessage()

	// m.SetHeader("From", "from@gmail.com") // alternative
	m.SetAddressHeader("From", "from@gmail.com", "Golang Indonesia")

	// m.SetAddressHeader("To", "to1@gmail.com", "Anggota #1 Golang")
	// m.SetAddressHeader("To", "to2@gmail.com", "Anggota #2 Golang")
	// m.SetAddressHeader("To", "to3@gmail.com", "Anggota #3 Golang")
	m.SetHeader("To", "to1@example.com", "to2@example.com", "to3@example.com") // alternative

	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	// m.SetHeader("Cc", "cc@gmail.com")  // alternative

	//
	// absPath, _ := filepath.Abs("mailer/image.png")
	// m.Attach(absPath)

	// //
	// m.Embed(absPath)

	m.SetHeader("Subject", "Gomail test subject")
	m.SetBody("text/html", `<h1>This is Gomail test body</h1> <img src="cid:image.png" alt="My image" />`)

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: false} // avoid this on production level

	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("success")
	}
}
