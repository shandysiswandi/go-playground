package main

import (
	"fmt"
	"net/smtp"
	"time"
)

var (
	smtpHost = "smtp.mailtrap.io"
	smtpPort = 2525
	smtpAddr = fmt.Sprintf("%s:%d", smtpHost, smtpPort)
	smtpUser = "42239792fa9ea4"
	smtpPass = "486010d4347713"
)

func main() {
	//
	from := "shandy@jojonomic.com"
	to := []string{"shandy@jojonomic.com"}

	// standart header
	msg := "Mime-Version: 1.0\r\n"
	msg += "Date: " + time.Now().Format(time.RFC1123Z) + "\r\n"
	msg += "From: \"AWS HRD\" <hrd@aws.com> \r\n"
	msg += "To: \"Shandy\" <shnady@gmail.com> \r\n"
	msg += "Cc: \"Alex\" <alex@aws.com>, \"John\" <john@aws.com> \r\n"
	msg += "Subject: Internship Software Engineer \r\n"
	msg += "Content-Type: text/html; charset=UTF-8\r\n"
	msg += "Content-Transfer-Encoding: quoted-printable\r\n"

	// anti-spam header
	msg += "Message-Id: 950124.162336@aws.com\r\n"

	// content
	msg += "\r\n"
	msg += `
	<html>
		<head>
			<title> Mail From AWS </title>
		</head>

		<body>
			<p>Here’s the space for our great sales pitch</p>
		</body>
	</html>`

	//
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost) // plain
	// auth := smtp.CRAMMD5Auth(smtpUser, smtpPass) // md5

	//
	err := smtp.SendMail(smtpAddr, auth, from, to, []byte(msg))
	if err != nil {
		fmt.Println("err", err)
	} else {
		fmt.Println("Email Sent Successfully!")
	}
}
