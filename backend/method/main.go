package main

import "fmt"

func main() {
	fmt.Println("main")

	akshat := User{"Akshat", "ahaja@gmail.com", true, 12}
	fmt.Println(akshat)
	fmt.Printf("Name is %v and email is %v\n", akshat.Name, akshat.Email)

	akshat.GetStatus() //calling method
	akshat.NewEmail()
	fmt.Println(akshat)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// method
func (u User) GetStatus() {
	fmt.Println("Is user is active\n", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@gmal.com"
	fmt.Println("Email of the user is:", u.Email) // it is doesnot change the actual point,
}
