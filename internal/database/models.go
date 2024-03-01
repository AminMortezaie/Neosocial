package db

type User struct {
	Age      int    `json:"age"`
	Name     string `json:"name"`
	Password string `json:"-"`
}