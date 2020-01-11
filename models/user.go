package models

import (
	"time"
)

// User struct represents an example user found in the database.
type User struct {
	ID							int			`json:"id,omitempty"`
	Email						string		`json:"email,omitempty"`
	Username					string		`json:"username,omitempty"`
	Password					string		`json:"password,omitempty"`
	IsActivated					bool		`json:"is_activated,omitempty"`
	ActivationTimestamp			time.Time	`json:"activation_timestamp,omitempty"`
	MailSentTimestamp			time.Time	`json:"mail_sent_timestamp,omitempty"`
	ActivationToken				string		`json:"activation_token,omitempty"`
	ActivationTokenExpiryDate	int			`json:"activation_token_expiry_date,omitempty"`
}