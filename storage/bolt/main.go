package main

import (
	"encoding/json"
	"fmt"
	"github.com/bukowa/goutils/storage/bolt/pkg"
)

type User struct {
	Name string `json:"Name"`
}

func (u *User) Key() []byte {
	return []byte(u.Name)
}

func main() {
	db := &pkg.BoltDatabase{}
	if err := db.Init(nil, "database.test", &User{}); err != nil {
		panic(err)
	}
	cntrl := &pkg.Controller{BoltDatabase: db}
	cntrl.Create(&User{Name: "asxx"})
	b, _ := cntrl.Get(&User{Name: "new"})
	if b == nil {
		panic(b)
	}
	var user User
	if err := json.Unmarshal(b, &user); err != nil {
		panic(err)
	}
	fmt.Println(user)
}
