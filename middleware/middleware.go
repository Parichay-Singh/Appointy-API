package middleware

import (
	"Appointy-API/encrypt"
	"Appointy-API/model"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Regex for URL match
var (
	getUserRe     = regexp.MustCompile(`^\/users\/(\w+)$`)
	getPostRe     = regexp.MustCompile(`^\/posts\/(\w+)$`)
	getUserPostRe = regexp.MustCompile(`\/posts/users\/(\w+)$`)
)

// Func to connect to MongoDB
func ConnectDB() (*mongo.Collection, *mongo.Collection) {
	// Setting client options
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	// Connecting to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	// Create two collections
	// Collection_1: Users' Data
	// Collection_2: Posts' data
	collection_1 := client.Database("ps").Collection("users")
	collection_2 := client.Database("ps").Collection("posts")
	return collection_1, collection_2
}

var collection_1, collection_2 = ConnectDB()

// Create user in DB (POST)
func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	// encrypting the password of user and storing the encrypted password
	key := "1122334455667788"
	hashed_password := encrypt.Encrypt(key, user.Password)
	user.Password = hashed_password
	result, err := collection_1.InsertOne(context.TODO(), &user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(result)
}

// Get user by userID (GET)
func GetUserByIDEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user model.User
	Path := getUserRe.FindStringSubmatch(r.URL.Path)
	id := Path[1]
	filter := bson.M{"_id": id}
	err := collection_1.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	json.NewEncoder(w).Encode(user)
}

// Create post in DB (POST)
func CreatePostEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post model.Post
	post.TimeStamp = time.Now()
	_ = json.NewDecoder(r.Body).Decode(&post)
	result, err := collection_2.InsertOne(context.TODO(), &post)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(result)
}

// Get post by postID (GET)
func GetPostByIDEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post model.Post
	Path := getPostRe.FindStringSubmatch(r.URL.Path)
	id := Path[1]
	filter := bson.M{"_id": id}
	err := collection_2.FindOne(context.TODO(), filter).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(post)
}

// Get all posts from a userID (GET)
func GetUsersPostByIdEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []model.Post
	Path := getUserPostRe.FindStringSubmatch(r.URL.Path)
	id := Path[1]
	cur, err := collection_2.Find(context.TODO(), bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	for cur.Next(context.TODO()) {
		var single_post model.Post
		err := cur.Decode(&single_post)
		if err != nil {
			log.Fatal(err)
		}
		if (single_post.UserID) == id {
			posts = append(posts, single_post)
		}
	}
	if err := cur.Err(); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(posts)
}
