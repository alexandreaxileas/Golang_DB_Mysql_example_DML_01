package dbfullexercise

/*
This source is responsible for load data at the tables Students and Subjects.
*/
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// File 2
// (DML) -> This func receive a pointer to the interface *sql.DB and others two strings
// The first contain the string used for "Prepare" with the sql instruction to insert into table
// The second, has the parameters/data that are going to be used to fill in the table referred at the First string parameter
type student struct {
	id   int
	name string
}
type student_vs_subject struct {
	id_student int
	id_subjetc int
}

// FnInsert -- insert data into any table
func FnInsert(db *sql.DB, sInsert string, sParams ...string) error {
	stm, err := db.Prepare(sInsert)
	if err != nil {
		return err
	}
	defer stm.Close()
	vTotalInsert := 0
	for _, param := range sParams {
		//fmt.Println(param)
		_, err := stm.Exec(param)
		vTotalInsert++
		if err != nil {
			return err
		}
	}
	fmt.Println("Total of Inserts....: ", vTotalInsert)
	return nil
}

// FnLoadData_STUDENTS_and_SUBJECTS -- Load tables STUDENTS and SUBJECTS with the basic data using FnInsert
func FnLoadData_STUDENTS_and_SUBJECTS() {
	db, err := FnConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//
	fmt.Println("Hi ! Database connected !")
	//
	vInsertStudents := "INSERT INTO students (name) VALUES (?)"
	err = FnInsert(db, vInsertStudents, "Maria", "Alexandre", "Flavio", "Edilaine", "Cristiano", "Aline", "Sandro", "Denis", "Kaue", "Vivian", "Ricardo", "Eduardo", "Monica", "Ana", "Nathalia", "Rita", "João Paulo", "Angelo", "Marcia", "Marcus", "Olivia", "Livia")
	if err != nil {
		panic(err) // We don't want to treat the error, just stop the program
	}
	fmt.Println("Data load on Table STUDENTS completed!")
	//
	vInsertSubjeccts := "INSERT INTO subjects (name) VALUES (?)"
	err = FnInsert(db, vInsertSubjeccts, "Biology", "Art", "Geography", "History", "Literature", "Mathematics Basic", "Physical", "English", "Economy", "Programming", "Mathematics Advanced")
	if err != nil {
		panic(err) // We don't want to treat the error, just stop the program
	}
	fmt.Println("Data load on Table SUBJECTS completed!")
}

// FnLoadData_STUDENTS_vs_SUBJECTS -- Load tables STUDENTS and SUBJECTS with the basic data using FnInsert
func FnLoadData_STUDENTS_vs_SUBJECTS() {
	// Other way to make inserts, simpler than than the function FnInsert used above.
	// Have fun trying to rebuilt the func FnInsert for using in the function below
	db, err := FnConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//
	fmt.Println("Hello ! Database connected again in order to fill out the table STUDENTS_vs_SUBJECTS")
	// Now we are going to use "select statement" to cross the data each row from table STUDENTS with each row from SUBJECTS
	iCount := 0
	rowsStudent, err1 := db.Query("select id, name from students")
	if err1 != nil {
		log.Fatal(err1)
	}
	for rowsStudent.Next() {
		var sd student
		rowsStudent.Scan(&sd.id, &sd.name) //Passa a referencia porque aqui vai vir na ordem determinada na query
		//For each row from tabel Student we get all rows from table Subjects
		rowsSubject, err2 := db.Query("select id from subjects")
		if err2 != nil {
			log.Fatal(err1)
		} // For each stundent lets connect each subject
		for rowsSubject.Next() {
			var sdj student_vs_subject
			sdj.id_student = sd.id
			//
			rowsSubject.Scan(&sdj.id_subjetc)
			if err2 != nil {
				log.Fatal(err1)
			}
			// inserting into rel_students_subjects affert the loop above where the code bring the data to struct student_vs_subject
			stm3, err3 := db.Prepare("INSERT INTO rel_students_subjects (id_student, id_subject, grade) VALUES (?,?,?)")
			if err3 != nil {
				panic(err3)
			}
			//
			_, err4 := stm3.Exec(sdj.id_student, sdj.id_subjetc, "0")
			if err4 != nil {
				panic(err4)
			}
		}
	}
	fmt.Println("Total of inserts made by STUDENTS_vs_SUBJECTS..:", iCount)
	fmt.Println("table STUDENTS_vs_SUBJECTS loaded successfully")
}

// FnInserStudent Insert one student
func FnInserStudent(pName string) {
	db, err1 := FnConnect()
	if err1 != nil {
		panic(err1)
	}
	defer db.Close()
	//
	vInsertStudents := "INSERT INTO students (name) VALUES (?)"
	err2 := FnInsert(db, vInsertStudents, pName)
	if err2 != nil {
		panic(err2) // We don't want to treat the error, just stop the program
	}

	fmt.Println("Created student: ", pName)
}
