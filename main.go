package main
import (
	"Appointy-API/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Handling API endpoints
	http.HandleFunc("/users", handler.CreateUserEndpoint)
	http.HandleFunc("/users/", handler.GetUserByIDEndpoint)
	http.HandleFunc("/posts", handler.CreatePostEndpoint)
	http.HandleFunc("/posts/", handler.GetPostByIDEndpoint)
	http.HandleFunc("/posts/users/", handler.GetUsersPostByIdEndpoint)
  
	//Run server on localhost
	fmt.Println("running server on localhost port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
