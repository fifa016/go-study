package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	name string
	sex  string
	age  int
}

func (u User) setName(name string) User {
	u.name = name
	return u
}

func (u *User) setNameP(name string) {
	u.name = name
}

type Admin struct {
	user      User
	privilege int
}

type Admin2 struct {
	User
	privl int
}

func main() {
	fmt.Println("=========start run=========")
	//testAnonymousStruct()
	ul := UserList{
		User{name: "u1"},
		User{name: "u2"},
	}
	fmt.Println("before swap u name : ", ul[0].name)
	ul.listSwap(0,1)
	fmt.Println("after swap u name : ", ul[0].name)
}

type UserList []User

func (u UserList) listSwap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
func testSwapStruct(a User, b User) {
	a, b = b, a
}
func testSetName() {
	u := User{
		name: "name1",
	}
	//u2 := &User{
	//	name:"name2",
	//}

	fmt.Printf("user name before modify: %s\n", u.name)
	u = u.setName("new name")
	fmt.Printf("user name: %s\n", u.name)
	u.setNameP("new namep")
	fmt.Printf("user namep: %s\n", u.name)
}

func testAnonymousStruct() {
	admin := Admin2{
		User: User{
			name: "admin2",
			sex:  "sex",
			age:  1,
		},
		privl: 2,
	}
	admin.setNameP("asdf")
	admin.setName("asdf")
	fmt.Printf("admin: %+v\n", admin)

}

func testStructInStruct() {
	admin := Admin{
		user: User{
			name: "Joey",
			sex:  "female",
			age:  18,
		},
		privilege: 1,
	}
	fmt.Printf("admin: %+v\n", admin)
}

func testUserStruct() {
	user := User{
		age: 1,
	}

	user1 := User{
		name: "Joey",
		age:  19,
	}
	user2 := User{
		"Joey", "male", 10,
	}
	fmt.Printf("user:%+v\n", user)
	fmt.Printf("user1:%+v\n", user1)
	fmt.Printf("user1 name:%s\n", user1.name)
	fmt.Printf("user1 age:%d\n", user1.age)
	fmt.Printf("user2:%+v\n", user2)

	var user3 User

	fmt.Printf("user3 memsize: %d\n", unsafe.Sizeof(user3))
	fmt.Printf("user1 memsize: %d\n", unsafe.Sizeof(user1))
	fmt.Printf("user3 : %v\n", user3)
}
