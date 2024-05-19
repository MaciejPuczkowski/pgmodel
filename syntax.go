package pgmodel

type DataType interface {
	SQLDataType() string
}

type ColumnConstraint interface {
	Name() string
	SQLColumnConstraint() (string, error)
}

type TableConstraint interface {
	Name() string
	SQLTableConstraint() (string, error)
}

type AlterTable struct {
	TableName string
	Action    AlterTableAction
}

type AlterTableAction interface {
	SQLAlterTableAction() string
}

type AlterColumn struct {
	ColumnName string
	Action     AlterColumnAction
}

type AlterColumnAction interface {
	SQLAlterColumnAction() string
}

type Migration struct {
	Actions []MigrationAction
}

type MigrationAction interface {
	SQLMigrationAction() string
}
