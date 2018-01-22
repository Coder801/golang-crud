package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	//"github.com/rs/cors"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Person struct {
	ID       string `json:"id,omitemply"`
	Username string `json:"username,omitemply"`
	Position string `json:"position,omitemply"`
}

var people []Person
var tmpl = template.Must(template.ParseGlob("template/*"))

func main() {
	r := httprouter.New()
	//handler := cors.Default().Handler(r)

	// Add data in array
	people = append(people, Person{ID: "1", Username: "Lokom", Position: "Frontend"})
	people = append(people, Person{ID: "2", Username: "Kilmer", Position: "Frontend"})
	people = append(people, Person{ID: "3", Username: "Tatsiak", Position: "QA"})
	people = append(people, Person{ID: "4", Username: "Rostik", Position: "Backend"})
	people = append(people, Person{ID: "5", Username: "Kreo", Position: "Frontend"})
	people = append(people, Person{ID: "6", Username: "Trueman", Position: "QA"})
	people = append(people, Person{ID: "7", Username: "Derry", Position: "Frontend"})
	people = append(people, Person{ID: "8", Username: "Mirkin", Position: "Team Lead"})
	people = append(people, Person{ID: "9", Username: "Savok", Position: "QA"})
	people = append(people, Person{ID: "10", Username: "Firebreath", Position: "QA"})
	people = append(people, Person{ID: "11", Username: "Nevill", Position: "Frontend"})

	// Template Routes
	r.GET("/", renderList)
	r.GET("/user", renderList)
	r.GET("/createUser", createUser)
	r.GET("/updateUser", updateUser)
	r.GET("/deleteUser", deleteUser)

	// Get Routes
	r.GET("/list", read)
	r.GET("/list/:key", read)
	r.POST("/create", create)
	r.POST("/update", update)
	r.POST("/delete", del)

	// Static files
	r.ServeFiles("/static/*filepath", http.Dir("static"))

	// Run server
	log.Println("Listening...")
	log.Fatalln(http.ListenAndServe(":8080", r))
}

func renderList(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	queryValues := r.URL.Query()
	key := queryValues.Get("id")

	if key != "" {
		for _, item := range people {
			if strings.ToLower(item.ID) == key {
				tmpl.ExecuteTemplate(w, "User", item)
				return
			}
		}
		tmpl.ExecuteTemplate(w, "User", "Not found")
	}
	tmpl.ExecuteTemplate(w, "Index", people)
}

func createUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tmpl.ExecuteTemplate(w, "Create", "string")
}

func updateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tmpl.ExecuteTemplate(w, "Update", people)
}

func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	tmpl.ExecuteTemplate(w, "Delete", people)
}

// API
func read(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	key := p.ByName("key")

	if key != "" {
		for _, item := range people {
			if strings.ToLower(item.Username) == key {
				fmt.Printf("Get person %v - %v\n", item.Username, item.Position)
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		fmt.Println("Person not found\n")
	}

	fmt.Println("Get all people\n")
	json.NewEncoder(w).Encode(people)
}

func create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	var person Person
	person.ID = r.FormValue("id")
	person.Username = r.FormValue("username")
	person.Position = r.FormValue("position")
	_ = json.NewDecoder(r.Body).Decode(&person)
	people = append(people, person)

	fmt.Printf("Create new person: %v - %v\n", person.Username, person.Position)
	json.NewEncoder(w).Encode(person)
	return
}

func update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	for key, item := range people {
		if strings.ToLower(item.ID) == r.FormValue("id") {
			people[key].Username = r.FormValue("username")
			people[key].Position = r.FormValue("position")

			fmt.Printf("Person %v was update\n", item.Username)
			json.NewEncoder(w).Encode(people[key])
			return
		}
	}
}

func del(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	id := r.FormValue("id")
	for key, item := range people {
		if item.ID == id {
			people = append(people[:key], people[key+1:]...)
			fmt.Printf("Person with 'id' - %v was remove\n", item.ID)
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}
