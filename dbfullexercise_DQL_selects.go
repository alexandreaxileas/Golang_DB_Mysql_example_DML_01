package dbfullexercise

import "fmt"

type GradeStudents struct {
	idStudent   int
	nameStudent string
	idSubject   int
	nameSubject string
	grade       float64
}

// PrintAllDataToScreen Print all students with their grade for each subject with their are connected
func PrintAllDataToScreen() {
	db, err1 := FnConnect() //
	if err1 != nil {
		panic(err1)
	}
	defer db.Close()
	//
	fmt.Println("Database connected. Now we are going to view what we inserted and modified")
	fmt.Println(";)")
	//
	rowGrades, err2 := db.Query("select st.id, st.name, sj.id, sj.name, rss.grade from students st ,subjects sj ,rel_students_subjects rss where st.id = rss.id_student and sj.id = rss.id_subject")
	if err2 != nil {
		panic(err2)
	}
	//
	iCount := 0
	for rowGrades.Next() {
		var grades GradeStudents
		iCount++
		rowGrades.Scan(&grades.idStudent, &grades.nameStudent, &grades.idSubject, &grades.nameSubject, &grades.grade)
		fmt.Printf("Student....: %v / %v - Subject....: %v / %v - Grade....: %.2f \n ", grades.idStudent, grades.nameStudent, grades.idSubject, grades.nameSubject, grades.grade)
	}
	fmt.Println("Total of line printed...: ", iCount)
}

//
// FnGetStudentByName Get one student using the Name
func FnGetStudentByName(pname string) (int, string) {
	db, err1 := FnConnect() //
	if err1 != nil {
		panic(err1)
	}
	defer db.Close()

	rowStudent, err2 := db.Query("select id, name from students where name = ? limit 1", pname)
	if err2 != nil {
		panic(err2)
	}
	var s student
	rowStudent.Next()
	rowStudent.Scan(&s.id, &s.name)
	return s.id, s.name
}
