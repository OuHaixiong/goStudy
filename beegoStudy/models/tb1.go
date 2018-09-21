package models;

import (
	"github.com/astaxie/beego/orm";

	// "github.com/astaxie/beego"
);

// Model Struct
type Tb1 struct {
	Id int
    Name string
}

func init() {
	orm.RegisterModel(new(Tb1)); // register model
}
