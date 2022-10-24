package main

import (
	"fmt"
)

type user struct {
	name      string
	age       int
	cellphone string
}

func (u *user) hello() {
	fmt.Printf("hello:%s\n", u.name)
}

func picUser(u *user) string {
	fmt.Printf("a user:", u)

	u.age = u.age + 1

	return "ok"
}

func main() {
	u := user{
		name:      "jamie",
		age:       32,
		cellphone: "13042321103",
	}

	u.hello()

	_ = picUser(&u)
	_ = picUser(&u)
}
