package pgmodel

import "github.com/MaciejPuczkowski/errs"

type Table struct {
	name        string
	isTemporary bool
	isUnlogged  bool
	isQuoted    bool
	columns     []*Column
	constraints []TableConstraint
}

func NewTable(name string) *Table {
	return &Table{
		name:        name,
		isQuoted:    shouldQuoteIdentifier(name),
		isTemporary: false,
		isUnlogged:  false,
		columns:     make([]*Column, 0),
		constraints: make([]TableConstraint, 0),
	}
}

func (t *Table) SQLName() string {
	if t.isQuoted {
		return quoteIdentifier(t.name)
	}
	return t.name
}

func (t *Table) Columns(columns ...*Column) *Table {
	t.columns = append(t.columns, columns...)
	return t
}

func (t *Table) Constraints(constraints ...TableConstraint) *Table {
	t.constraints = append(t.constraints, constraints...)
	return t
}

func (t *Table) Temporary() *Table {
	t.isTemporary = true
	return t
}

func (t *Table) Unlogged() *Table {
	t.isUnlogged = true
	return t
}

func (t *Table) IsTemporary() bool {
	return t.isTemporary
}

func (t *Table) IsUnlogged() bool {
	return t.isUnlogged
}

func (t *Table) SQLMigrationAction() (string, error) {
	m := t.toMap()
	m["Name"] = t.SQLName()
	r, err := fromTemplate(tmplCreateTable, m)
	if err != nil {
		return "", errs.Wrap(err)
	}
	return r, nil
}

func (t *Table) toMap() map[string]any {
	return map[string]any{
		"Name":        t.name,
		"IsTemporary": t.isTemporary,
		"IsUnlogged":  t.isUnlogged,
		"IsQuoted":    t.isQuoted,
		"Columns":     t.columns,
		"Constraints": t.constraints,
	}
}
