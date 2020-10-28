package main

import "fmt"

type User struct {
	id   int
	name string
}

type Manager struct {
	User
}

func (user *User) ToString() string { // receiver = &(Manager.User)
	return fmt.Sprintf("User: %p, %v", user, user)
}

func main() {
	m := Manager{User: User{1, "Tom"}}
	fmt.Printf("Manager: %p\n", &m)
	fmt.Println(m.ToString())
}
