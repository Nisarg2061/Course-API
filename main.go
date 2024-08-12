package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Model for the file/database
type Course struct {
	CourseId    int     `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice float32 `json:"courseprice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fname   string `json:"fname"`
	Website string `json:"website"`
}

// Fake database
var authors []Author = []Author{
	{Fname: "Nisarg", Website: "bynisarg.in"},
	{Fname: "Bhupendra", Website: "bybhupendra.in"},
}

var courses []Course = []Course{
	{CourseId: 1, CourseName: "Golang", CoursePrice: 399.99, Author: &authors[1]},
	{CourseId: 2, CourseName: "Docker", CoursePrice: 399.99, Author: &authors[0]},
	{CourseId: 4, CourseName: "Python", CoursePrice: 399.99, Author: &authors[0]},
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func IsEmpty(c *Course) bool {
	return c.CourseName == ""
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by me!<h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	course, exist := loadCourses()[r.PathValue("id")]

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(course)
}

func loadCourses() map[string]Course {
	res := make(map[string]Course, len(courses))

	for _, x := range courses {
		res[strconv.Itoa(x.CourseId)] = x
	}

	return res
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Enter valid data.")
	}
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("GET /", serveHome)
	router.HandleFunc("GET /courses", getAllCourses)
	router.HandleFunc("GET /course/{id}", getCourse)

	server := http.Server{
		Addr:    ":4000",
		Handler: router,
	}

	fmt.Println("Listening on port 4000...")

	err := server.ListenAndServe()
	check(err)
}
