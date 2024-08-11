package main

// Model for the file/database
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice float32 `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fname   string `json:"fname"`
	Website string `json:"website"`
}

// Fake database
var courses []Course

// Middleware, helper files
func IsEmpty(c *Course) bool {
	return c.CourseId == "" && c.CourseName == ""
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
}
