package model

// import "myapp/dataStore/postgres"

// type Course struct {
// 	CourseId   string `json:"csid"`
// 	CourseName string `json:"cname"`
// }

// const (
// 	queryCourseInsert  = "INSERT INTO course(courseid, coursename) VALUES($1, $2);"
// 	queryCourseGetUser = "SELECT courseid, coursename FROM course WHERE courseid=$1;"
// 	queryCourseUpdate  = "UPDATE course SET courseid=$1, coursename=$2 WHERE courseid=$3 RETURNING courseid;"
// 	queryCourseDelete  = "DELETE FROM course WHERE courseid=$1 RETURNING courseid;"
// 	queryCourseSelect  = "SELECT * from course;"
// )

// func (c *Course) Create() error {
// 	_, err := postgres.Db.Exec(queryCourseInsert, c.CourseId, c.CourseName)
// 	return err
// }

// func (c *Course) Read() error {
// 	return postgres.Db.QueryRow(queryCourseGetUser, c.CourseId).Scan(&c.CourseId, &c.CourseName)
// }

// func (c *Course) Update(csid string) error {
// 	row := postgres.Db.QueryRow(queryCourseUpdate, c.CourseId, c.CourseName, csid)
// 	err := row.Scan(&c.CourseId)
// 	return err
// }

// func (c *Course) Delete() error {
// 	row := postgres.Db.QueryRow(queryCourseDelete, c.CourseId)
// 	err := row.Scan(&c.CourseId)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func GetAllCourses() ([]Course, error) {
// 	rows, getErr := postgres.Db.Query(queryCourseSelect)
// 	if getErr != nil {
// 		return nil, getErr
// 	}

// 	courses := []Course{}

// 	for rows.Next() {
// 		var c Course
// 		dbErr := rows.Scan(&c.CourseId, &c.CourseName)
// 		if dbErr != nil {
// 			return nil, dbErr
// 		}

// 		courses = append(courses, c)
// 	}
// 	rows.Close()
// 	return courses, nil
// }
