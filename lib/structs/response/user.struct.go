package response

type Users struct {
	Users []User `json:"users"`
}

type UserRegistration struct {
	User User `json:"user"`
}

type User struct {
	ID          string `json:"id"`
	Uuid        string `json:"uuid"`
	Email       string `json:"email"`
	Inactive    bool   `json:"inactive"`
	Deactivated bool   `json:"userDeactivated"`
	Dormant     bool   `json:"dormant"`
	AvatarUrl   string `json:"avatarUrl"`
	JwtToken    string `json:"jwtToken"`
}
