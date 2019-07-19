package main
import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r, "./hello.html" )
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "About Page")
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Index Page")
	})
	fmt.Println("Server* is listening...")
	http.ListenAndServe(":8181", nil)
}
