package model

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`

	// holds ip of instance that served this user
	ServedBy string `json:"servedBy"`

	// just to play with invoking other services
	Quote Quote `json:"quote"`
}

type Quote struct {
	Text     string `json:"quote"`
	ServedBy string `json:"ipAddress"`
	Language string `json:"language"`
}
