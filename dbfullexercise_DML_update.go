package dbfullexercise

import (
	"fmt"
)

// File 3
// FnLoadUpdateGrades -- This function will set grade for all subjects that is connected to all students. In order to do that, lets use random procedures
func FnLoadUpdateGrades() {
	db, err1 := FnConnect() //
	if err1 != nil {
		panic(err1)
	}
	defer db.Close() // Ensuring to close connection
	//
	fmt.Println("Database connected to set grades. Maximum 10.00 and Minimum 0.00")
	//
	rowStudetSubject, err2 := db.Query("select id from rel_students_subjects")
	if err2 != nil {
		panic(err2)
	}
	var vId int
	iCount := 0
	for rowStudetSubject.Next() {
		rowStudetSubject.Scan(&vId)
		rowUpdate, err3 := db.Prepare("update rel_students_subjects set grade = ? where id = ?")
		if err3 != nil {
			panic(err3)
		}
		iCount++
		rowUpdate.Exec(fnRetGrade(), vId)
	}
	fmt.Println("total of Grades updated....:", iCount)
}
