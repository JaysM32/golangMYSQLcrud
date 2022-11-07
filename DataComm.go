package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type project struct {
	ID       int    `json:"ID"`
	Name     string `json:"Name"`
	Grade    int    `json:"Grade"`
	assignID int    `json:"AID"`
	gradeID  int    `json:"GID`
	feedback string `json:"feed"`
}

func main() {
	fmt.Println("Go MySQL Tutorial")

	deleteByProjID(143)

}

func insertInto(projID int, projName string, projGrade int, assignID int, feedback string) {
	//connect to database
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/intelsysgrading")
	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO `grading` (`projectID`, `projectName`, `projectGrade`, `assignmentID`,`projectFeedback`) VALUES (?, ?, ?, ?, ? )", projID, projName, projGrade, assignID, feedback)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()

}

func readDatabyProjID(projID int) {
	//connect to database
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/intelsysgrading")
	defer db.Close()

	// perform a db.Query read
	read, err := db.Query("SELECT * FROM `grading` WHERE `projectID`=?", projID)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer read.Close()

	for read.Next() {
		var proj project

		err = read.Scan(&proj.ID, &proj.Name, &proj.Grade, &proj.assignID, &proj.gradeID, &proj.feedback)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("project ID =", proj.ID)
		fmt.Println("project Name =", proj.Name)
		fmt.Println("project Grade =", proj.Grade)
		fmt.Println("project Assignment ID =", proj.assignID)
		fmt.Println("project grade ID =", proj.gradeID)
		fmt.Println("project feedback =", proj.feedback)

	}

}

func readDatabyGrdID(gradeID int) {
	//connect to database
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/intelsysgrading")
	defer db.Close()

	// perform a db.Query read
	read, err := db.Query("SELECT * FROM `grading` WHERE `gradeID`=?", gradeID)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer read.Close()

	for read.Next() {
		var proj project

		err = read.Scan(&proj.ID, &proj.Name, &proj.Grade, &proj.assignID, &proj.gradeID, &proj.feedback)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("project ID =", proj.ID)
		fmt.Println("project Name =", proj.Name)
		fmt.Println("project Grade =", proj.Grade)
		fmt.Println("project Assignment ID =", proj.assignID)
		fmt.Println("project grade ID =", proj.gradeID)
		fmt.Println("project feedback =", proj.feedback)

	}

}

func updateFeedbackData(gradeID int, feedback string) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/intelsysgrading")
	defer db.Close()

	// perform a db.Query read
	read, err := db.Query("UPDATE `grading` SET `projectFeedback`=? WHERE `gradeID`=? ", feedback, gradeID)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer read.Close()
}

func deleteByProjID(projID int) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/intelsysgrading")
	defer db.Close()

	// perform a db.Query read
	read, err := db.Query("DELETE FROM `grading` WHERE `projectID`=?", projID)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer read.Close()
}
