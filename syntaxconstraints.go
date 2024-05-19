package pgmodel

import (
	"fmt"
	"strings"
)

type PrimaryKey struct {
	name    string
	columns []string
}

func NewPrimaryKey(columns ...string) *PrimaryKey {
	return &PrimaryKey{columns: columns}
}

func (pk *PrimaryKey) Name() string {
	return pk.name
}

func (pk *PrimaryKey) SetName(name string) *PrimaryKey {
	pk.name = name
	return pk
}

func (pk *PrimaryKey) SQLTableConstraint() (string, error) {
	if pk.name != "" {
		return fmt.Sprintf("CONSTRAINT %s PRIMARY KEY (%s)", quoteIdentifier(pk.name), strings.Join(pk.columns, ", ")), nil
	}
	return fmt.Sprintf("PRIMARY KEY (%s)", strings.Join(pk.columns, ", ")), nil
}

func (pk *PrimaryKey) SQLColumnConstraint() (string, error) {
	if pk.name != "" {
		return fmt.Sprintf("CONSTRAINT %s PRIMARY KEY", quoteIdentifier(pk.name)), nil
	}
	return "PRIMARY KEY", nil
}
