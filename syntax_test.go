package pgmodel

import (
	"testing"
)

func TestTable_Usage1(t *testing.T) {
	table := NewTable("User").Columns(
		NewColumn("id", Integer),
		NewColumn("Name", Varchar(255)),
	)
	r, err := table.SQLMigrationAction()
	if err != nil {
		t.Fatal(err)
	}
	expected := `CREATE TABLE "User" (
	id INTEGER,
	"Name" VARCHAR(255)
)`
	if r != expected {
		t.Fatalf("expected %q, got %q", expected, r)
	}
}

func TestTable_Usage2_constraint(t *testing.T) {
	table := NewTable("User").Columns(
		NewColumn("id", Integer),
		NewColumn("Name", Varchar(255)),
	).Constraints(
		NewPrimaryKey("id"),
	)
	r, err := table.SQLMigrationAction()
	if err != nil {
		t.Fatal(err)
	}
	expected := `CREATE TABLE "User" (
	id INTEGER,
	"Name" VARCHAR(255),
	PRIMARY KEY (id)
)`
	if r != expected {
		t.Fatalf("expected %q, got %q", expected, r)
	}
}

func TestTable_Usage_constraint_named(t *testing.T) {
	table := NewTable("User").Columns(
		NewColumn("id", Integer),
		NewColumn("Name", Varchar(255)),
	).Constraints(
		NewPrimaryKey("id").SetName("User_pk"),
	)
	r, err := table.SQLMigrationAction()
	if err != nil {
		t.Fatal(err)
	}
	expected := `CREATE TABLE "User" (
	id INTEGER,
	"Name" VARCHAR(255),
	CONSTRAINT "User_pk" PRIMARY KEY (id)
)`
	if r != expected {
		t.Fatalf("expected %q, got %q", expected, r)
	}
}

func TestTable_Usage3_temporary(t *testing.T) {
	table := NewTable("User").Columns(
		NewColumn("id", Integer),
		NewColumn("Name", Varchar(255)),
	).Temporary()
	r, err := table.SQLMigrationAction()
	if err != nil {
		t.Fatal(err)
	}
	expected := `CREATE TEMPORARY TABLE "User" (
	id INTEGER,
	"Name" VARCHAR(255)
)`
	if r != expected {
		t.Fatalf("expected %q, got %q", expected, r)
	}
}

func TestTable_Usage4_unlogged(t *testing.T) {
	table := NewTable("User").Columns(
		NewColumn("id", Integer),
		NewColumn("Name", Varchar(255)),
	).Unlogged()
	r, err := table.SQLMigrationAction()
	if err != nil {
		t.Fatal(err)
	}
	expected := `CREATE UNLOGGED TABLE "User" (
	id INTEGER,
	"Name" VARCHAR(255)
)`
	if r != expected {
		t.Fatalf("expected %q, got %q", expected, r)
	}
}

func TestTable_Usage5_temporary_unlogged(t *testing.T) {
	table := NewTable("User").Columns(
		NewColumn("id", Integer),
		NewColumn("Name", Varchar(255)),
	).Temporary().Unlogged()
	r, err := table.SQLMigrationAction()
	if err != nil {
		t.Fatal(err)
	}
	expected := `CREATE TEMPORARY UNLOGGED TABLE "User" (
	id INTEGER,
	"Name" VARCHAR(255)
)`
	if r != expected {
		t.Fatalf("expected %q, got %q", expected, r)
	}
}

func TestTable_Usage6_column_constraint(t *testing.T) {
	table := NewTable("User").Columns(
		NewColumn("id", Integer).Constrain(NewPrimaryKey()),
		NewColumn("Name", Varchar(255)),
	)
	r, err := table.SQLMigrationAction()
	if err != nil {
		t.Fatal(err)
	}
	expected := `CREATE TABLE "User" (
	id INTEGER PRIMARY KEY,
	"Name" VARCHAR(255)
)`
	if r != expected {
		t.Fatalf("expected %q, got %q", expected, r)
	}
}

func TestTable_Usage7_column_constraint_named(t *testing.T) {
	table := NewTable("User").Columns(
		NewColumn("id", Integer).Constrain(NewPrimaryKey().SetName("User_pk")),
		NewColumn("Name", Varchar(255)),
	)
	r, err := table.SQLMigrationAction()
	if err != nil {
		t.Fatal(err)
	}
	expected := `CREATE TABLE "User" (
	id INTEGER CONSTRAINT "User_pk" PRIMARY KEY,
	"Name" VARCHAR(255)
)`
	if r != expected {
		t.Fatalf("expected %q, got %q", expected, r)
	}
}

func TestTable_Usage8_column_constraint_shortcut(t *testing.T) {
	table := NewTable("User").Columns(
		NewColumn("id", Integer).PrimaryKey(),
		NewColumn("Name", Varchar(255)),
	)
	r, err := table.SQLMigrationAction()
	if err != nil {
		t.Fatal(err)
	}
	expected := `CREATE TABLE "User" (
	id INTEGER PRIMARY KEY,
	"Name" VARCHAR(255)
)`
	if r != expected {
		t.Fatalf("expected %q, got %q", expected, r)
	}
}
