package main

import (
	"fmt"
	"encoding/json"
)

type Contact struct {
	Email string
	Phone int64
}

func main () {
	data := `{
		"email": "anish.budha@gmail.com",
		"phone": 4372587777	
	}`

	var contact Contact
	err := json.Unmarshal([]byte(data), &contact)

	fmt.Println(contact, err)
}
