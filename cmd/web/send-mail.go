package main

import (
	"log"
	"os"
	"time"

	"github.com/bishal7679/SpiceEx/internal/models"
	"github.com/joho/godotenv"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
			sendMsg(msg)
		}
	}()
}

func init() {
	 if err := godotenv.Load(); err != nil {
        log.Print("No .env file found")
    }
}
func sendMsg(m models.MailData) {
	// its weird that its actually creating a SMTP server not client
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
	server.Username = m.From
	server.Password = os.Getenv("GMAIL_PASS")
	server.Encryption = mail.EncryptionSTARTTLS

	// creating a client to server
	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextHTML, m.Content)
	// email.Attach(&mail.File{FilePath: "/static/images/logo.png", Name:"Gopher.png", Inline: true})

	err = email.Send(client)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Email sent!")
	}
}
