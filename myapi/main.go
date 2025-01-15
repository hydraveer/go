package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId       string  `json:courseid`
	CourseName     string  `json:coursename`
	CoursePrice    int     `json:price`
	CoursePlatform string  `json:website`
	Author         *Author `json:Instructor`
}
type Author struct {
	FullName string `json:name`
	website  string `json:website`
}

var courses []Course

func (c *Course) IsEmpty() bool {

	return c.CourseName == ""
}

func main() {
	fmt.Println("APIs")
	r := mux.NewRouter()
	courses = append(courses, Course{
		CourseId:"2",
		CourseName: "ReactJs",
		CoursePrice: 299,
		CoursePlatform: "www.udemy.com",
		Author: &Author{
			FullName: "Harkirat Singh",
			website: "www.udemy.com",
		},
	})
	courses = append(courses, Course{
		CourseId:"3",
		CourseName: "NextJs",
		CoursePrice: 299,
		CoursePlatform: "www.udemy.com",
		Author: &Author{
			FullName: "Code with harry",
			website: "www.udemy.com",
		},
	})
	//routing
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses",GetAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",GetOneCourse).Methods("GET")
	r.HandleFunc("/course",CreateCourse).Methods("POST")
	r.HandleFunc("/course/{id}",UpdateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}",DeleteCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Home"))
}
func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no find the course")
	return
}
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Creating Course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("No data is sent")
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func UpdateCourse(w http.ResponseWriter, r *http.Request){

		fmt.Println("Update Course")
		w.Header().Set("Content-Type","application/json")
		param := mux.Vars(r)

		if r.Body==nil {
			json.NewEncoder(w).Encode("Please send write data in body")
		}
		for id, course := range courses{
			if course.CourseId==param["id"] {
				courses = append(courses[:id],courses[id+1:]...)
				var val Course 
				_ = json.NewDecoder(r.Body).Decode(&val)
				val.CourseId = param["id"]
				courses = append(courses, val)
				json.NewEncoder(w).Encode(val)
				return
			}
		}
		json.NewEncoder(w).Encode("there is some error while updating the course")
		return
}
func DeleteCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete one course")
	param := mux.Vars(r)
	if r.Body==nil {
		json.NewEncoder(w).Encode("Please select correct course")
	}
	for id, course := range courses{
		if course.CourseId==param["id"] {
			courses = append(courses[:id],courses[id+1:]...)
			json.NewEncoder(w).Encode("The course deleted successfully")
			return
		}
	}
	json.NewEncoder(w).Encode("The course is already deleted")
	return
}
