package gmail

import (
	"fmt"
	"net/smtp"
)

type IGmail interface {
	SendCode(toGmail, code string) error
}
type Gmail struct {
	GmailLogin string
	GmailPass  string
	GmailHost  string
	GmailPort  string
}

func NewGmail(gmailLogin, gmailPass, gmailHost, gmailPort string) *Gmail {
	return &Gmail{
		GmailLogin: gmailLogin,
		GmailPass:  gmailPass,
		GmailHost:  gmailHost,
		GmailPort:  gmailPort,
	}
}

func (g *Gmail) SendCode(toGmail, code string) error {
	message := []byte(fmt.Sprintf("Subject: Your Code\r\n\r\nHi! Your code to register with Gmail: %s", code))

	auth := smtp.PlainAuth("", g.GmailLogin, g.GmailPass, g.GmailHost)

	err := smtp.SendMail(g.GmailHost+":"+g.GmailPort, auth, g.GmailLogin, []string{toGmail}, message)
	if err != nil {
		return err
	}

	return nil
}
