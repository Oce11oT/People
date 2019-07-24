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

func main() {
	connStr := "user=postgres password=people1234 dbname=people sslmode=disable" //error is here!
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	http.HandleFunc("/", IndexHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}
