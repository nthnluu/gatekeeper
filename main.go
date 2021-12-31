package gatekeeper

import (
	"fmt"
	"gatekeeper/ability"
)

type User struct {
	IsAdmin bool
	ID int
}

func defineAbility(user User) (a *ability.Ability) {
	a = ability.NewAbility()

	if user.IsAdmin {
		a.Can("create", "User")
		a.Can("update", "User")
		a.Can("delete", "User")
	}

	a.Can("read", "User")

	return
}

func main() {
	user1 := User{
		IsAdmin: false,
		ID:      0,
	}

	user2 := User{
		IsAdmin: true,
		ID:      1,
	}

	newAbility1 := defineAbility(user1)
	newAbility2 := defineAbility(user2)

	fmt.Println(newAbility1.Check("update", "User"))
	fmt.Println(newAbility2.Check("update", "User"))
}
