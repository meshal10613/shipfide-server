package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"path/filepath"
	"server/internal/config"
)

type Sender interface {
	SendEmail(to string, subject string, templateFile string, data any) error
	SendVerificationEmail(to, name, otp string) error
	SendPasswordResetEmail(to, name, otp string) error
	SendAdminWelcomeEmail(to, name, tempPassword, loginLink string) error
}

type emailSender struct {
	config *config.Config
}

func NewEmailSender(cfg *config.Config) Sender {
	return &emailSender{config: cfg}
}

func (s *emailSender) SendEmail(to string, subject string, templateFile string, data any) error {
	// 1. Resolve template path and parse it with <% and %> delimiters
	tmplPath := filepath.Join("pkg", "templates", templateFile)
	tmpl, err := template.New(templateFile).Delims("<%", "%>").ParseFiles(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	// 2. Build RFC 822 email message
	from := s.config.SmtpFromEmail
	if from == "" {
		from = s.config.SmtpUsername
	}
	if from == "" {
		from = "no-reply@shipfide.com"
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subjectHeader := fmt.Sprintf("Subject: %s\n", subject)
	fromHeader := fmt.Sprintf("From: %s\n", from)
	toHeader := fmt.Sprintf("To: %s\n", to)

	msg := []byte(fromHeader + toHeader + subjectHeader + mime + body.String())

	// 3. Send email using SMTP or log to console in dev mode
	if s.config.SmtpHost == "" {
		log.Printf("\n==================================================\n"+
			"[DEV EMAIL SENDER (SMTP Not Configured)]\n"+
			"To: %s\n"+
			"Subject: %s\n"+
			"Content:\n%s\n"+
			"==================================================", to, subject, body.String())
		return nil
	}

	addr := fmt.Sprintf("%s:%s", s.config.SmtpHost, s.config.SmtpPort)
	auth := smtp.PlainAuth("", s.config.SmtpUsername, s.config.SmtpPassword, s.config.SmtpHost)

	err = smtp.SendMail(addr, auth, from, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send SMTP email: %w", err)
	}

	return nil
}

func (s *emailSender) SendVerificationEmail(to, name, otp string) error {
	data := map[string]any{
		"Name": name,
		"OTP":  otp,
	}
	return s.SendEmail(to, "Verify Your Email Address", "email_verification.ejs", data)
}

func (s *emailSender) SendPasswordResetEmail(to, name, otp string) error {
	data := map[string]any{
		"Name": name,
		"OTP":  otp,
	}
	return s.SendEmail(to, "Reset Your Password", "reset_password.ejs", data)
}

func (s *emailSender) SendAdminWelcomeEmail(to, name, tempPassword, loginLink string) error {
	data := map[string]any{
		"Name":      name,
		"Email":     to,
		"Password":  tempPassword,
		"LoginLink": loginLink,
	}
	return s.SendEmail(to, "Welcome to the Shipfide Team - Admin Portal Credentials", "admin_welcome.ejs", data)
}
