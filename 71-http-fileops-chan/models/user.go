package models

type User struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	Status       string `json:"status"`
	LastModified int64  `json:"last_modified"`
}

func NewUser(name, email, mobile, status string) *User {
	return &User{Name: name, Email: email, Mobile: mobile, Status: status}
}

// There is a concept of serialization and deserialization

// `json:"status"` --> called as tags. These work based on reflection, it is very easy for serialize and deserialize
// you can json, yaml, gorm etc.. many formats..
// some project they use some third party packages even to do validations by using tags
