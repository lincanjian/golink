package email

import "net/smtp"

type golink_email struct {
	emailClient *smtp.Client
	smtpServer  string
	smtpPort    string
	from        string
	licenseCode string
}
