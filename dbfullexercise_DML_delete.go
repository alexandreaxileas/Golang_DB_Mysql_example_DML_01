package dbfullexercise

import (
	"database/sql"
	"fmt"
)

// Attention !

/* This code I will use a different approach from the functions DML (insert and update) before.
   This time lets use transaction controls. Means that the command "delete" will just be commited
   affer explicitly the DataBase receive the command "commit"

*/

// FnDelete -- insert data into any table
func FnDelete(db *sql.DB, sDelete string, sIdParams int) error {

	tx, _ := db.Begin() //This point indicate the begin of the transaction
	stm, err := db.Prepare(sDelete)
	if err != nil {
		return err
	}
	//
	defer stm.Close()
	//
	_, err2 := stm.Exec(sIdParams)
	if err2 != nil {
		tx, _ := db.Begin() //This point indicate the begin of the transaction
		tx.Rollback()       // This point has a different approach, because the comand "begin" above
		fmt.Println("Fail on delete [001]")
		panic(err2)
	} else {
		tx.Commit() // As mentioned above this time it's needed to execute the command "commit" to confirm the transaction
		fmt.Println("Data deleted successful")
	}

	return nil
}

// FnDeleteStudent -- insert data into any table
func FnDeleteStudent(pId int) {
	db, err1 := FnConnect()
	if err1 != nil {
		panic(err1)
	}
	defer db.Close()
	//
	vDeleteStudents := "delete from students where id = ?"
	err2 := FnDelete(db, vDeleteStudents, pId)
	if err2 != nil {
		panic(err2) // We don't want to treat the error, just stop the program
	}
	fmt.Println("student deleted: ", pId)
}
