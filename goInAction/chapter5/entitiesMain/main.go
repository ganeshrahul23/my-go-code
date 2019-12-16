package main

import (
	"fmt"
	"github.com/ganeshrahul23/goInAction/chapter5/entities"
)

func main() {
	user1 := entities.User{Name: "Ganesh"}
	fmt.Println(user1)
	user1.SetEmail("ganesheiepec@gmail.com")
	fmt.Println(&user1)
	fmt.Println(user1.String())
	fmt.Println("========================================")

	admin1 := entities.Admin{Rights: 10}
	admin1.Name = "Rahul"
	//admin1.user.Name = "Rahul" will not compile
	//because user is a unexported type
	admin1.Email = "avenger.ganesh46@gmail.com"

	fmt.Println(admin1)
}
