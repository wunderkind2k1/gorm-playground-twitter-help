package main

import (
	"fmt"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

// TestTwitter1722317893096874229 explains https://gorm.io/docs/query.html#Selecting-Specific-Fields
func TestTwitter1722317893096874229(t *testing.T) {
	var users []User
	user := User{Name: "jinzhu", Age: 99, Active: true}
	DB.Create(&user)
	defer DB.Delete(&user)

	result := DB.Select([]string{"Name", "Active"}).Find(&users)

	if result.Error != nil {
		t.Fatalf("Epected user not found. Got error: %v", result.Error)
	}

	fmt.Printf("Users found: %v", users)
}

// TestTwitter1722317893096874229ResponseStruct explains https://gorm.io/docs/advanced_query.html
func TestTwitter1722317893096874229ResponseStruct(t *testing.T) {

	type ResponseUser struct {
		Name   string `json:"ResponseName"`
		Active string `json:"ResponseActive"`
	}
	var responseUsers []ResponseUser
	user := User{Name: "jinzhu", Age: 99, Active: true}
	DB.Create(&user)
	defer DB.Delete(&user)

	result := DB.Model(&User{}).Select([]string{"Name", "Active"}).Find(&responseUsers)

	if result.Error != nil {
		t.Fatalf("Epected user not found. Got error: %v", result.Error)
	}

	fmt.Printf("Users found: %v", responseUsers)
}
