package main

import (
	"fmt"
	"time"

	"github.com/alexandreaxileas/dbfullexercise"
)

func main() {

	horaInicio := time.Now().Format("02/01/2006 03:04:05")
	fmt.Println(horaInicio)
	now := time.Now()
	dbfullexercise.CreateDataBase()
	//
	dbfullexercise.FnLoadData_STUDENTS_and_SUBJECTS()
	dbfullexercise.FnLoadData_STUDENTS_vs_SUBJECTS()
	//
	dbfullexercise.FnLoadUpdateGrades()
	//
	dbfullexercise.PrintAllDataToScreen()
	//

	diff := time.Until(now) //now.Sub(time.Now())
	fmt.Println("Elapsed time in Milliseconds: ", diff.Milliseconds())
	horaFim := time.Now().Format("02/01/2006 03:04:05")
	fmt.Println(horaFim)

	// Delete transaction. But before delete, lets insert a new record.
	dbfullexercise.FnInserStudent("Henrique Oliveira")
	var vId int
	var vName string
	vId, vName = dbfullexercise.FnGetStudentByName("Henrique Oliveira")
	//
	if vId > 0 {
		fmt.Printf("\n Student-ID: %v \n Student NAME: %v", vId, vName)
		dbfullexercise.FnDeleteStudent(vId)
	}
	//
	vId = 0
	vName = ""
	vId, vName = dbfullexercise.FnGetStudentByName("Henrique Oliveira")
	fmt.Printf("\n Student-ID: %v \n Student NAME: %v", vId, vName)
}
