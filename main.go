package main

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
)

func main() {
	fmt.Println("Starting email sender")

	// Connect to the remote SMTP server.
	// Note would need to lookup gmail smtp server
	c, err := smtp.Dial("smtp.gmail.com:25")
	if err != nil {
		log.Fatal(err)
	}

	// Set the sender and recipient first
//	if err := c.Mail("security.operations@target.com"); err != nil {

	if err := c.Mail("vermeerp@gmail.com"); err != nil {
		log.Fatal(err)
	}
	if err := c.Rcpt("vermeerp@gmail.com"); err != nil {
		log.Fatal(err)
	}

	// Send the email body.
	wc, err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	buf := bytes.NewBufferString("Subject: Digital Security Alert\r\n\r\nThis is the email body.")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}

	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Send the QUIT command and close the connection.
	err = c.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
