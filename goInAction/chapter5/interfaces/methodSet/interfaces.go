package main

import "fmt"

type notifier interface {
	notify()
}
type user struct {
	name, email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

func main() {
	u := user{"Ganesh", "ganesheiepec@gmail.com"}
	sendNotification(&u)
}

func sendNotification(n notifier) {
	n.notify()
}
