package main
import ("fmt"
		 "net/http"
		 "encoding/json"
		)

type Books []Book 
var books Books


func main(){
	fmt.Println("joke")
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)
}


func home(w http.ResponseWriter, r *http.Request){
	fmt.Println("kill")
	books =  Books{
		Book{Title: "felix", Writer :"messy", Year:"joe"}, 
		Book{Title: "felix", Writer :"messy", Year:"joe"}, 
		Book{Title: "felix", Writer :"messy", Year:"joe"}}
	json.NewEncoder(w).Encode(books)
}


type Book struct{
	Title string `json:"Title"`
	Writer string `json:"writer"`
	Year string `json:"content"`
}
