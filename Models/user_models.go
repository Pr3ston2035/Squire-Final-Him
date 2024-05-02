package Models

type Profile struct {
	ID       string `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

type Credentials struct {
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	ID       string `json:"id"`
}
