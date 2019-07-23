package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"testing"
)

type S string

func TestConectDb(t *testing.T) {
	var actual *gorm.DB
	actual = ConectDb("default_database")
	fmt.Println(actual)
}
