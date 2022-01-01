package policy

import (
	"testing"
)

type User struct {
	IsAdmin bool
	ID      int
}

func definePolicy(user User) (a *Policy) {
	a = NewPolicy()

	if user.IsAdmin {
		a.Can("create", "User")
		a.Can("update", "User")
		a.Can("delete", "User")
	}

	a.Can("read", "User")

	return
}

func TestAbs(t *testing.T) {
	newPolicy := NewPolicy()
	newPolicy.Can("read", "User")

	if res := newPolicy.Check("read", "User"); res != true {
		t.Errorf("Check(read, User) returned %v when it was expected to return true", res)
	}

	if res := newPolicy.Check("update", "User"); res != false {
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

	adminPolicy := definePolicy(adminUser)
	regularPolicy := definePolicy(regularUser)

	if res := adminPolicy.Check("read", "User"); res != true {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

	if res := regularPolicy.Check("read", "User"); res != true {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

	if res := adminPolicy.Check("update", "User"); res != true {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

	if res := adminPolicy.Check("create", "User"); res != true {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

	if res := adminPolicy.Check("delete", "User"); res != true {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

	// ------

	if res := regularPolicy.Check("update", "User"); res != false {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

	if res := regularPolicy.Check("create", "User"); res != false {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

	if res := regularPolicy.Check("delete", "User"); res != false {
		t.Errorf("Check(update, User) returned %v when it was expected to return true", res)
	}

}
