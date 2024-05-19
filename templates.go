package pgmodel

import "text/template"

var tmplCreateTable = template.Must(template.New("create_table").Parse(
	`CREATE{{ if .IsTemporary }} TEMPORARY{{ end }}{{ if .IsUnlogged }} UNLOGGED{{ end }} TABLE {{ .Name }} (
{{- range $i, $c := .Columns }}
	{{- if $i }},{{ end }}
	{{ $c.SQLColumn }}
{{- end }}
{{- if .Constraints }}
  {{- range .Constraints }},
	{{ .SQLTableConstraint }}
  {{- end }}
{{- end }}
)`))

var tmplCreateColumn = template.Must(template.New("create_column").Parse(
	`{{ .Name }} {{ .DataType.SQLDataType }}{{ if .Constraints }} {{ range .Constraints }}{{ .SQLColumnConstraint }}{{ end }}{{ end }}`),
)

var tmplCreateConstraint = template.Must(template.New("create_constraint").Parse(`
{{ .Name }} {{ .SQLTableConstraint }}
`))

var tmplAlterTable = template.Must(template.New("alter_table").Parse(`
ALTER TABLE {{ .TableName }} {{ .Action.SQLAlterTableAction }}
`))

var tmplAlterTableActionAddColumn = template.Must(template.New("alter_table_action_add_column").Parse(`
ADD COLUMN {{ .Name }} {{ .DataType.SQLDataType }}{{ if .Constraints }} {{- range .Constraints }}{{ .SQLColumnConstraint }}{{ end }}{{ end }}
`))

var tmplAlterTableActionDropColumn = template.Must(template.New("alter_table_action_drop_column").Parse(`
DROP COLUMN {{ .Name }}
`))

var tmplAlterColumn = template.Must(template.New("alter_column").Parse(`
ALTER COLUMN {{ .ColumnName }} {{ .Action.SQLAlterColumnAction }}
`))

var tmplAlterColumnActionSetNotNull = template.Must(template.New("alter_column_action_set_not_null").Parse(`
SET NOT NULL
`))

var tmplAlterColumnActionDropNotNull = template.Must(template.New("alter_column_action_drop_not_null").Parse(`
DROP NOT NULL
`))

var tmplAlterColumnActionSetDefault = template.Must(template.New("alter_column_action_set_default").Parse(`
SET DEFAULT {{ .DefaultValue }}
`))

var tmplAlterColumnActionDropDefault = template.Must(template.New("alter_column_action_drop_default").Parse(`
DROP DEFAULT
`))

var tmplAlterColumnActionSetType = template.Must(template.New("alter_column_action_set_type").Parse(`
TYPE {{ .DataType.SQLDataType }}{{ if .Constraints }} {{- range .Constraints }}{{ .SQLColumnConstraint }}{{ end }}{{ end }}
`))

var tmplAlterColumnActionDropType = template.Must(template.New("alter_column_action_drop_type").Parse(`
TYPE {{ .DataType.SQLDataType }}{{ if .Constraints }} {{- range .Constraints }}{{ .SQLColumnConstraint }}{{ end }}{{ end }}
`))

var tmplAlterColumnActionAddConstraint = template.Must(template.New("alter_column_action_add_constraint").Parse(`
ADD CONSTRAINT {{ .Name }} {{ .SQLColumnConstraint }}
`))

var tmplAlterColumnActionDropConstraint = template.Must(template.New("alter_column_action_drop_constraint").Parse(`
DROP CONSTRAINT {{ .Name }}
`))

var tmplDropTable = template.Must(template.New("drop_table").Parse(`
DROP TABLE {{ .Name }}
`))

var tmplDropColumn = template.Must(template.New("drop_column").Parse(`
ALTER TABLE {{ .TableName }} DROP COLUMN {{ .Name }}
`))

var tmplDropConstraint = template.Must(template.New("drop_constraint").Parse(`
ALTER TABLE {{ .TableName }} DROP CONSTRAINT {{ .Name }}
`))
