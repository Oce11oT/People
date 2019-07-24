package main
import (
	"database/sql"
	"fmt"
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

		_, err = database.Exec("insert into people (uname, sex, age) values (?, ?, ?)",
			name, sex, age)

		if err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/", 301)
	}else{
		http.ServeFile(w,r, "create.html")
	}
}

func main() {
	connStr := "user=postgres password=people1234 dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/create", CreateHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
