package pgmodel

type Column struct {
	name        string
	dataType    DataType
	constraints []ColumnConstraint
	isQuoted    bool
}

func (c *Column) toMap() map[string]any {
	return map[string]any{
		"Name":        c.name,
		"DataType":    c.dataType,
		"Constraints": c.constraints,
	}
}
func (c *Column) SQLColumn() (string, error) {
	m := c.toMap()
	m["Name"] = c.SQLName()
	return fromTemplate(tmplCreateColumn, m)
}

func NewColumn(name string, dataType DataType) *Column {
	return &Column{
		name:        name,
		dataType:    dataType,
		constraints: make([]ColumnConstraint, 0),
		isQuoted:    shouldQuoteIdentifier(name),
	}
}

func (c *Column) SQLName() string {
	if c.isQuoted {
		return quoteIdentifier(c.name)
	}
	return c.name
}

func (c *Column) Constrain(constraint ...ColumnConstraint) *Column {
	c.constraints = append(c.constraints, constraint...)
	return c
}

func (c *Column) PrimaryKey() *Column {
	return c.Constrain(NewPrimaryKey())
}
