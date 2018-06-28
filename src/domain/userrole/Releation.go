package userrole

import (
	"toolkit"
	"strings"
	"github.com/jinzhu/gorm"
)

var(
	RoleIdJoinChar=","
)

//Record user and role list
type Releation struct {
	gorm.Model
	UserId string
	RoleIds string
}


func(releation *Releation) JoinRoleId(roleIds []int)string{

	joinString:=""

	for _,id:=range roleIds{
		joinString+=toolkit.ToString(id)+RoleIdJoinChar
	}

	return joinString
}


func(releation *Releation) SplitRoleId()[]int{

	idStringCol:=strings.Split(releation.RoleIds,RoleIdJoinChar)

	idIntCol:=[]int{}

	for _,id:=range idStringCol{
		idIntCol=append(idIntCol,toolkit.ToInt(id))
	}

	return idIntCol
}




