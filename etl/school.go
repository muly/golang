package main

import "database/sql"

type stageSchool struct {
	SchoolID   int
	SchoolName string
	Address1   string
	Address2   string
	City       string
	State      string
	Zip        string
	Country    string
}

type school struct {
	SchoolID   int
	SchoolName string
	Address    string
}

type schoolETL struct {
	db *sql.DB
}

func (e *schoolETL) Extract(tx *sql.Tx) ([]stageSchool, error) {
	rows, err := tx.Query("SELECT id, name, address1, address2, city, state, zip, country FROM stage_school")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schools []stageSchool
	for rows.Next() {
		var s stageSchool
		if err := rows.Scan(&s.SchoolID, &s.SchoolName, &s.Address1, &s.Address2, &s.City, &s.State, &s.Zip, &s.Country); err != nil {
			return nil, err
		}
		schools = append(schools, s)
	}

	return schools, nil
}

func (e *schoolETL) Transform(data []stageSchool) ([]school, error) {
	var schools []school
	for _, s := range data {
		address := s.Address1
		if s.Address2 != "" {
			address += ", " + s.Address2
		}
		address += ", " + s.City + ", " + s.State + ", " + s.Country + ", " + s.Zip

		schools = append(schools, school{
			SchoolID:   s.SchoolID,
			SchoolName: s.SchoolName,
			Address:    address,
		})
	}

	return schools, nil
}

func (e *schoolETL) Load(data []school, tx *sql.Tx) error {
	stmt, err := tx.Prepare("INSERT INTO school (id, name, address) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, s := range data {
		if _, err := stmt.Exec(s.SchoolID, s.SchoolName, s.Address); err != nil {
			return err
		}
	}

	return nil
}
