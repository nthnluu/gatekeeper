package ability

import (
	"testing"
)

type User struct {
	IsAdmin bool
	ID int
}

func defineAbility(user User) (a *Ability) {
	a = NewAbility()

	if user.IsAdmin {
		a.Can("create", "User")
		a.Can("update", "User")
		a.Can("delete", "User")
	}

	a.Can("read", "User")

	return
}

func TestAbs(t *testing.T) {
	newAbility := NewAbility()
	newAbility.Can("read", "User")

	if res := newAbility.Check("read", "User"); res != true {
		t.Errorf("Check(read, User) returned %v when it was expected to return true", res)
	}

	if res := newAbility.Check("update", "User"); res != false {
		t.Errorf("Check(update, User) returned %v when it was expected to return false", res)
	}

	adminUser := User{
		IsAdmin: true,
		ID:      0,
	}

	regularUser := User{
		IsAdmin: false,
		ID:      1,
	}

	adminAbility := defineAbility(adminUser)
	regularAbility := defineAbility(regularUser)

	if res := adminAbility.Check("read", "User"); res != true {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

	if res := regularAbility.Check("read", "User"); res != true {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

	if res := adminAbility.Check("update", "User"); res != true {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

	if res := adminAbility.Check("create", "User"); res != true {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

	if res := adminAbility.Check("delete", "User"); res != true {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

	// ------

	if res := regularAbility.Check("update", "User"); res != false {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

	if res := regularAbility.Check("create", "User"); res != false {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

	if res := regularAbility.Check("delete", "User"); res != false {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

}