package main
import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)
type People struct {
	Id int
	Name string
	Sex string
	Age int
}
var database *sql.DB

//show table
func IndexHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := database.Query("select * from people")
	if err != nil {
		log.Println(err)
	}
defer rows.Close()
	people := []People{}

	for rows.Next(){
		p := People{}
		err := rows.Scan(&p.Id, &p.Name, &p.Sex, &p.Age)
		if err != nil{
			fmt.Println(err)
			continue
		}
		people = append(people, p)
	}

	tmpl, _ := template.ParseFiles("index.html")
	tmpl.Execute(w, people)
}
//add data
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		name := r.FormValue("name")
		sex := r.FormValue("sex")
		age := r.FormValue("age")

		_, err = database.Exec("insert into people (uname, sex, age) values ($1, $2, $3)",
			name, sex, age)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	}else{
		http.ServeFile(w,r, "create.html")
	}
}

//edit.html
func EditPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	row := database.QueryRow("select * from people where uid = $1", id)
	pers := People{}
	err := row.Scan(&pers.Id, &pers.Name, &pers.Sex, &pers.Age)
	if err != nil{
		log.Println(err)
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}else{
		tmpl, _ := template.ParseFiles("edit.html")
		tmpl.Execute(w, pers)
	}
}

//get data and update
func EditHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	id := r.FormValue("id")
	name := r.FormValue("name")
	sex := r.FormValue("sex")
	age := r.FormValue("age")

	_, err = database.Exec("update people set uname = $1, sex = $2, age = $3 where id = $4",
		name, sex, age, id)

	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", 301)
}

func main() {
	connStr := "user=postgres password=people1234 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()

	router :=mux.NewRouter()
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/create", CreateHandler)

	router.HandleFunc("/edit/{id:[0-9]+}", EditPage).Methods("GET")
	router.HandleFunc("/edit/{id:[0-9]+}", EditHandler).Methods("POST")

	http.Handle("/",router)
	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
