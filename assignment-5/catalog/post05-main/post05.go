package post05

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// Connection details
var (
	Hostname = ""
	Port     = 2345
	Username = ""
	Password = ""
	Database = ""
)

type MSDSCourse struct {
	CID     string
	CNAME   string
	CPREREQ string
}

func openConnection() (*sql.DB, error) {
	// Connection string
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Hostname, Port, Username, Password, Database)
	// Open database connection
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// AddCourse adds a new course to the msdsCourseCatalog table
func AddCourse(course MSDSCourse) error {
	db, err := openConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	insertQuery := `INSERT INTO msdscoursecatalog (CID, CNAME, CPREREQ) VALUES ($1, $2, $3)`
	_, err = db.Exec(insertQuery, course.CID, course.CNAME, course.CPREREQ)
	if err != nil {
		return err
	}
	return nil
}

// List all courses in the database
func ListCourses() ([]MSDSCourse, error) {
	db, err := openConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Select all courses
	query := `SELECT CID, CNAME, CPREREQ FROM msdscoursecatalog`
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []MSDSCourse
	for rows.Next() {
		var course MSDSCourse
		if err := rows.Scan(&course.CID, &course.CNAME, &course.CPREREQ); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}
