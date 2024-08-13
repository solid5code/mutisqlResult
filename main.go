package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type (
	User struct {
		ID         int    `gorm:"column:Id"`
		Name       string `gorm:"column:Name"`
		Account    string `gorm:"column:Account"`
		Password   string `gorm:"column:Password"`
		CreateTime string `gorm:"column:CreateTime"`
	}

	Member struct {
		ID   int    `gorm:"column:Id"`
		Name string `gorm:"column:Name"`
	}
)

func main() {
	db, err := gorm.Open(mysql.Open("root:1234@tcp(localhost:3306)/datacenter?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	rows, err := db.Raw("CALL datacenter.GetUserAndMember").Rows()
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		if err := db.ScanRows(rows, &members); err != nil {
			panic(err)
		}
	}

	var users []User
	if rows.NextResultSet() {
		for rows.Next() {
			if err := db.ScanRows(rows, &users); err != nil {
				panic(err)
			}
		}
	}
	userb, _ := json.MarshalIndent(users, "", "\t")
	memberb, _ := json.MarshalIndent(members, "", "\t")

	fmt.Printf("users: %+v\n", string(userb))
	fmt.Printf("members: %+v\n", string(memberb))
}
