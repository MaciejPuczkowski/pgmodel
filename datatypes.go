package pgmodel

import "fmt"

type integer struct {
}

func (i integer) SQLDataType() string {
	return "INTEGER"
}

var Integer DataType = integer{}

type varchar struct {
	length int
}

func (v varchar) SQLDataType() string {
	if v.length > 0 {
		return fmt.Sprintf("VARCHAR(%d)", v.length)
	}
	return "VARCHAR"
}

func Varchar(length int) DataType {
	return varchar{length: length}
}
