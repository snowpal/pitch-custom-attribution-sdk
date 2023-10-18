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

type LastModified struct {
	By string `json:"by"`
	On string `json:"on"`
}

type UserMetadata struct {
	CreatedBy    string       `json:"createdBy"`
	UpdateBy     []string     `json:"updatedBy"`
	LastModified LastModified `json:"lastModified"`
}
