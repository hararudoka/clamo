package object

// User will be used to store user information in DB and in JSON responses
type User struct {
	ID int `json:"id"` // unique id for user

	Username string `json:"username"` // unique name that user will be able to change

	Password string `json:"password"` // just password for authentication
}

type Message struct {
	ID int `json:"id"` // unique id for message

	FromID int `json:"from_id"` // id of user who sent the message
	ToID   int `json:"to_id"`   // id of user who will receive the message

	Text string `json:"text"` // content of the message
}
