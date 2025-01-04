package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	pass   = "tsMSSb0B9c"
	dbname = "postgres"
)

func main() {

	//	var (
	//		 ch string
	//		 c string
	//		 ID string
	//		 NAME string
	//		 MARKS string
	//	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, pass, dbname)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		fmt.Println("Enter the correct details")
	} else {
		fmt.Println("Connection established")
	}

	//func show() {

	//	rows,err := db.Query("Select * from student")
	//	for rows.next() {
	//    rows.Scan(&ID,&NAME,&MARKS)
	// fmt.Println(ID,NAME,MARKS)
	//}
	//}

	//	for c {

	//	fmt.Println("ENTER C/R/U/D")

	//	fmt.Scanln(ch)

	//	switch ch {
	//    case "C" :
	//createStmt := `CREATE TABLE IF NOT EXISTS student(id int primary key , name VARCHAR(20),marks int); `
	//db.Exec(createStmt)
	//	case "R" :
	insertStmt := `INSERT INTO student VALUES(1,'John',20); `
	db.Exec(insertStmt)
	insertStmts := `INSERT INTO student VALUES(2,'wick',21); `
	db.Exec(insertStmts)
	//show()
	//	case "U" :
	updateStmt := `UPDATE student SET name = 'Jala' where id = 1;`
	db.Exec(updateStmt)
	//show()
	//	case "D" :
	deleteStmt := `DELETE from student where id = 2;`
	db.Exec(deleteStmt)
	//	show()
	//	}

	//	fmt.Println("Do you want to continue ? yes/no")
	//	fmt.Scanln(c)
	//}
}
