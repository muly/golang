package main

import "database/sql"

type stageStudent struct {
	StudentID int
	FirstName string
	Address1  string
	Address2  string
	City      string
	State     string
	Zip       string
	Country   string
}

type student struct {
	StudentID int
	FirstName string
	Address   string
}

type studentETL struct {
	db *sql.DB
}

func (e *studentETL) Extract(tx *sql.Tx) ([]stageStudent, error) {
	rows, err := tx.Query("SELECT id, name, address1, address2, city, state, zip, country FROM stage_student")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []stageStudent
	for rows.Next() {
		var s stageStudent
		if err := rows.Scan(&s.StudentID, &s.FirstName, &s.Address1, &s.Address2, &s.City, &s.State, &s.Zip, &s.Country); err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	return students, nil
}

func (e *studentETL) Transform(data []stageStudent) ([]student, error) {
	var students []student
	for _, s := range data {
		address := s.Address1
		if s.Address2 != "" {
			address += ", " + s.Address2
		}
		address += ", " + s.City + ", " + s.State + ", " + s.Country + ", " + s.Zip

		students = append(students, student{
			StudentID: s.StudentID,
			FirstName: s.FirstName,
			Address:   address,
		})
	}

	return students, nil
}

func (e *studentETL) Load(data []student, tx *sql.Tx) error {
	stmt, err := tx.Prepare("INSERT INTO student (id, name, address) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, s := range data {
		if _, err := stmt.Exec(s.StudentID, s.FirstName, s.Address); err != nil {
			return err
		}
	}

	return nil
}
