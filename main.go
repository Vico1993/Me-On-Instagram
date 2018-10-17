package main

import (
	"fmt"

	goinsta "github.com/ahmdrz/goinsta"
)

const (
	username = "****"
	password = "****"
)

func main() {
	insta := goinsta.New(username, password)

	if err := insta.Login(); err != nil {
		fmt.Println("Can't login :" + err.Error())
		return
	}

	user, err := insta.Profiles.ByName("vicop21")
	if err != nil {
		fmt.Println("Can't get access to progiles by name :" + err.Error())
		return
	}

	users := user.Followers()
	if err != nil {
		fmt.Println(err)
		return
	}

	i := 1
	for users.Next() {
		fmt.Println("Next:", users.NextID)
		for _, user := range users.Users {
			i++
			fmt.Printf("  - %s\n", user.Username)
		}
	}
	fmt.Println("Followers:", i)
}
