// Test generated by RoostGPT for test go-sample using AI Type Vertex AI and AI Model code-bison-32k

/*
	**Unit test scenarios for the createTable function:**

	1. **Test that the function creates a table with the correct name.**
	- Create a new database connection.
	- Call the createTable function.
	- Query the database to see if the table was created.
	- Assert that the table was created with the correct name.

2. **Test that the function creates a table with the correct schema.**
  - Create a new database connection.
  - Call the createTable function.
  - Query the database to get the schema of the table.
  - Assert that the schema of the table is correct.

3. **Test that the function creates a table with the correct constraints.**
  - Create a new database connection.
  - Call the createTable function.
  - Query the database to get the constraints of the table.
  - Assert that the constraints of the table are correct.

4. **Test that the function creates a table with the correct data type.**
  - Create a new database connection.
  - Call the createTable function.
  - Query the database to get the data type of the table.
  - Assert that the data type of the table is correct.

5. **Test that the function creates a table with the correct indexes.**
  - Create a new database connection.
  - Call the createTable function.
  - Query the database to get the indexes of the table.
  - Assert that the indexes of the table are correct.
*/
package main

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/tannergabriel/learning-go/advanced-programs/FiberPostgresCRUD/database"
)

var db *gorm.DB

func TestCreateTable_0224d0f13a(t *testing.T) {
	// Test that the function creates a table with the correct name.
	db = database.DBConn
	createTable()
	query := `SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'items');`
	var exists bool
	db.Raw(query).Scan(&exists)
	if !exists {
		t.Errorf("Table 'items' was not created.")
	}

	// Test that the function creates a table with the correct schema.
	query = `SELECT column_name, data_type FROM information_schema.columns WHERE table_name = 'items';`
	var columns []struct {
		ColumnName string
		DataType   string
	}
	db.Raw(query).Scan(&columns)
	expectedColumns := []struct {
		ColumnName string
		DataType   string
	}{
		{"id", "serial"},
		{"Title", "character varying"},
		{"Owner", "character varying"},
		{"Rating", "integer"},
		{"created_at", "date"},
		{"updated_at", "date"},
		{"deleted_at", "date"},
	}
	for i, column := range columns {
		if column.ColumnName != expectedColumns[i].ColumnName || column.DataType != expectedColumns[i].DataType {
			t.Errorf("Column '%s' has the wrong data type. Expected '%s', got '%s'", column.ColumnName, expectedColumns[i].DataType, column.DataType)
		}
	}

	// Test that the function creates a table with the correct constraints.
	query = `SELECT constraint_name, constraint_type FROM information_schema.table_constraints WHERE table_name = 'items';`
	var constraints []struct {
		ConstraintName string
		ConstraintType string
	}
	db.Raw(query).Scan(&constraints)
	expectedConstraints := []struct {
		ConstraintName string
		ConstraintType string
	}{
		{"pk_books", "PRIMARY KEY"},
	}
	for i, constraint := range constraints {
		if constraint.ConstraintName != expectedConstraints[i].ConstraintName || constraint.ConstraintType != expectedConstraints[i].ConstraintType {
			t.Errorf("Constraint '%s' has the wrong type. Expected '%s', got '%s'", constraint.ConstraintName, expectedConstraints[i].ConstraintType, constraint.ConstraintType)
		}
	}

	// Test that the function creates a table with the correct data type.
	query = `SELECT data_type FROM information_schema.columns WHERE table_name = 'items' AND column_name = 'id';`
	var dataType string
	db.Raw(query).Scan(&dataType)
	if dataType != "serial" {
		t.Errorf("Column 'id' has the wrong data type. Expected 'serial', got '%s'", dataType)
	}

	// Test that the function creates a table with the correct indexes.
	query = `SELECT index_name, column_name FROM information_schema.indexes WHERE table_name = 'items';`
	var indexes []struct {
		IndexName  string
		ColumnName string
	}
	db.Raw(query).Scan(&indexes)
	expectedIndexes := []struct {
		IndexName  string
		ColumnName string
	}{
		{"pk_books", "id"},
	}
	for i, index := range indexes {
		if index.IndexName != expectedIndexes[i].IndexName || index.ColumnName != expectedIndexes[i].ColumnName {
			t.Errorf("Index '%s' has the wrong column. Expected '%s', got '%s'", index.IndexName, expectedIndexes[i].ColumnName, index.ColumnName)
		}
	}
}
