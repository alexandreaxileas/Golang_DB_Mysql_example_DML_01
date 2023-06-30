package dbfullexercise

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// File (1)
// This file will create all tables e provide some example of DDL statements
// Have fun !!!!!!!!!!!

// Exec - sql.Result is an interface returned by this function
func Exec(db *sql.DB, sql string) sql.Result { //
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

// CreateDataBase - Create all tables that will be used
func CreateDataBase() {
	// first we must open the connection
	db, err := sql.Open("mysql", "principal:skinner@/")
	if err != nil {
		panic(err)
	}
	defer db.Close() // Closing the connection using "DEFER". This way we avoid to forget to close it
	// Bellow let create the DataBase and tables that going to be use at the next example.
	Exec(db, "create database if not exists high_school")
	// poiting high_school DataBase
	Exec(db, "use high_school")
	// Creating all tables.
	Exec(db, "drop table if exists rel_students_subjects") // This table is here if will be necessery to re-execute all scripts
	//
	Exec(db, "drop table if exists students")
	Exec(db, `create table students (
		id integer auto_increment,
		name varchar(80),
		PRIMARY KEY (id)
		)`) // Its necessary to be together -> ")`)"
	Exec(db, "drop table if exists subjects")
	Exec(db, `create table subjects (
		id integer auto_increment,
		name varchar(80),
		PRIMARY KEY (id)
		)`)
	
	Exec(db, `create table rel_students_subjects (
		id integer auto_increment,
		id_student integer,
		id_subject  integer,
		grade numeric(4,2),
		PRIMARY KEY (id),
		INDEX idx_student (id_student),
		FOREIGN KEY (id_student)
			REFERENCES students(id)
			ON DELETE CASCADE,
		INDEX idx_subjects (id_subject),
		FOREIGN KEY (id_subject)
			REFERENCES  subjects(id)
			ON DELETE CASCADE	
		)`)

	fmt.Println("DataBase and all tables successfully created")
}
