package main

import (
	"infrastructure/repository/sqllite"
	"domain/userrole"
	"os"
	"domain/role"
	"github.com/jinzhu/gorm"
	"fmt"
)

// User describes a user
type User struct {
	gorm.Model
	Id   int64
	Name string
	BillingAddress Address
	ShippingAddress Address
	Emails []Email
	Languages []Language
}
type Address struct{
	gorm.Model
	Address1 string
	Address2 string
}

type Email struct{
	gorm.Model
	Address string
}

type Language struct{
	gorm.Model
	Name string
}

func main() {
	os.Remove(sqllite.DBFileName)
	//test1()


	engine:=sqllite.GetDbEngine().Engine

	engine.CreateTable(&userrole.Releation{})
	engine.CreateTable(&role.Role{})
	engine.CreateTable(&role.Permission{})

	//engine.Create(&userrole.Releation{UserId:"1",RoleIds:[]*userrole.RoleIdVal{&userrole.RoleIdVal{1},&userrole.RoleIdVal{2}}})
	//engine.Create(&userrole.Releation{UserId:"1",RoleIds:[]int{1,2,3}})
	engine.Create(&userrole.Releation{UserId:"1",RoleIds:"1,2,3,4"})

	releation:=&userrole.Releation{UserId:"1"}

	newRole:=&role.Role{
		RoleName:"role1",
		Desc:"desc1",
		Permissions:[]*role.Permission{
			&role.Permission{
				Name:"permissino1",
			},
			&role.Permission{
				Name:"permissino2",
			},
		},
	}

	newRole2:=&role.Role{
		RoleName:"role2",
		Desc:"desc1",
		Permissions:[]*role.Permission{
			&role.Permission{
				Name:"permissino1",
			},
			&role.Permission{
				Name:"permissino2",
			},
		},
	}

	engine.Create(newRole)
	//engine.Save(newRole)
	retrieveRole:=&role.Role{RoleName:"role1"}
	engine2:=sqllite.GetDbEngine().Engine
	engine2.Create(newRole2)

	engine.Preload("Permissions").Find(retrieveRole)

	engine.First(retrieveRole)

	engine.First(releation)

	test2()

}

func test1(){
	engine:=sqllite.GetDbEngine().Engine

	user := User{
		Name:            "jinzhu",
		BillingAddress:  Address{Address1: "Billing Address - Address 1"},
		ShippingAddress: Address{Address1: "Shipping Address - Address 1"},
		Emails:          []Email{
			{ Address:"jinzhu@example.com"},
			{ Address:"jinzhu-2@example@example.com"},
		},
		Languages:       []Language{
			{Name: "ZH"},
			{Name: "EN"},
		},
	}

	engine.Create(&user)
	engine.Save(&user)
}

func test2(){
	engine:=sqllite.GetDbEngine().Engine

	engine.CreateTable(&Language{})

	success:=engine.Create(&Language{Name:"123"})

	fmt.Println(success)

}