package main
import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/gorilla/mux"


)

type Article struct {
	Id int `json:"Id"`
	Title string `json:"Title"`
	Desc string  `json:"Desc"`
	Content string `json:"Content"`

}
//defining an global Articles array or slice to populate in our main function
var Articles []Article

//endpoints
func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage")
	fmt.Println("Endpoint hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Fprintf(w, "Key: " + key)
	//loop over all articles and if articles.Id equals
	//the key we pass in it will return article encoded as JSON
	for _,article := range Articles{
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}
// nije mi skroz jasno sve
func createNewArticle(w http.ResponseWriter, r *http.Request){
	//get the body of our POST request
	//return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}
func deleteArticle(w http.ResponseWriter, r *http.Request){
	vars := mux(r)
	id := vars["id"]
	for index, article := range Articles{
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

//handlers
//Handleri uzimaju polozaj i funkciju koja radi get ili post
func requestHandlers(){
	//new router
	myRouter := mux.NewRouter()
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	http.HandleFunc("/article/{id}", returnSingleArticle)
	http.HandleFunc("/article", createNewArticle).Methods("POST")
	http.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")


	/* log.Fatal(http.ListenAndServe(":8080", nil)) */
	log.Fatal(http.ListenAndServe(":8080", myRouter ))
	
}

func main(){
	//Articles array populated by Article structs
	//scope of this variable is global scope
	Articles = []Article{
		Article{Id: 1, Title:"Title1", Desc:"Desc1", Content:"Content1"},
		Article{Id: 2, Title:"Title2", Desc:"Desc2", Content:"Content2"},
	}
	requestHandlers()
}
