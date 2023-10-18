package common

type ProfileName struct {
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	LastName   string `json:"LastName"`
}

type PhoneNumber struct {
	Primary     bool   `json:"primary"`
	PhoneNumber string `json:"phoneNumber"`
}

type EmailAddress struct {
	Primary      bool   `json:"primary"`
	EmailAddress string `json:"emailAddress"`
}
