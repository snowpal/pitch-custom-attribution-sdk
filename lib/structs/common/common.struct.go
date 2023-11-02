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

type SelectOption struct {
	ID    int    `json:"ID"`
	Value string `json:"value"`
}

type NSSChild struct {
	Name      string       `json:"name"`
	FieldType string       `json:"fieldType"`
	Options   *[]NSSOption `json:"options,omitempty"`
}

type NSSOption struct {
	ID    int       `json:"ID"`
	Value string    `json:"value"`
	Child *NSSChild `json:"child,omitempty"`
}

type BagAttribute struct {
	ID       string `json:"ID"`
	Required bool   `json:"required"`
	List     bool   `json:"list"`
	Metadata bool   `json:"metadata"`
}

type BagPrimitiveAttr struct {
	Value  *string   `json:"value"`
	Values *[]string `json:"values"`
}

type SelectValue struct {
	SelectedID  *string   `json:"selectedID"`
	SelectedIDs *[]string `json:"selectedIDs"`
}

type BagSelectAttr struct {
	Value  *SelectValue   `json:"value"`
	Values *[]SelectValue `json:"values"`
}

type BagSelectOption struct {
	*SelectValue
	*NSSValue
}

type NSSValueChild struct {
	Text   *BagPrimitiveAttr `json:"text"`
	Number *BagPrimitiveAttr `json:"number"`
	Date   *BagPrimitiveAttr `json:"date"`
	File   *BagPrimitiveAttr `json:"file"`
	Option *BagSelectOption  `json:"option"`
}

type NSSValue struct {
	SelectedID string         `json:"selectedID"`
	Child      *NSSValueChild `json:"child"`
}

type BagNSSAttr struct {
	Value  *NSSValue   `json:"value"`
	Values *[]NSSValue `json:"values"`
}
