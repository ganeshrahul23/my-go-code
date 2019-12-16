package entities

import "fmt"

type User struct {
	Name  string
	email string
}

func (u *User) String() string {
	return fmt.Sprintf("Name : %v\tMail : %v", u.Name, u.email)
}

func (u *User) Email() string {
	return u.email
}

func (u *User) SetEmail(email string) {
	u.email = email
}

type user struct {
	Name  string
	Email string
}

type Admin struct {
	user
	Rights int
}

func (a Admin) String() string {
	return fmt.Sprintf("Admin Name : %v\tAdmin Mail : %v\tAdmin Rights : %v",
		a.user.Name, a.Email, a.Rights)
}
