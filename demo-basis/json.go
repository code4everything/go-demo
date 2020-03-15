package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	pa := &Address{"chengdu", "china"}
	wa := &Address{City: "shanghai", Country: "china"}
	u := &User{"ease", []*Address{pa, wa}, "none"}
	// JSON编码
	j, _ := json.Marshal(u)
	fmt.Println(string(j))
	// JSON解码
	var uc User
	if err := json.Unmarshal(j, &uc); err == nil {
		fmt.Println(uc)
	} else {
		fmt.Println(err.Error())
	}
}

type Address struct {
	City    string
	Country string
}

type User struct {
	Name      string
	Addresses []*Address
	Remark    string
}

func (u User) String() string {
	us, err := json.MarshalIndent(u, "", "\t")
	if err == nil {
		return string(us)
	}
	return "to json string error: " + err.Error()
}
