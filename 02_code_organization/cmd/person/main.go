package main

import (
	"fmt"

	architecture "github.com/KenShiiii/Golang-architecture/02_code_organization"
	"github.com/KenShiiii/Golang-architecture/02_code_organization/database/harddrive"
	"github.com/KenShiiii/Golang-architecture/02_code_organization/database/mongo"
)

func main() {
	dbm := mongo.Db{}
	dbh := harddrive.Db{}

	u1 := architecture.User{
		First: "James",
	}

	u2 := architecture.User{
		First: "Jenny",
	}

	architecture.Put(dbm, 1, u1)
	architecture.Put(dbh, 2, u2)

	fmt.Println(architecture.Get(dbm, 1))
	fmt.Println(architecture.Get(dbh, 2))
	fmt.Println(architecture.Get(dbh, 3))

}
