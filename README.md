# Appointy-API ✨

This is an API mimicking a basic version of Instagram's backend. It is an HTTP json API capable of operations such as:

## Operations
### Creating a User 🤵
* Is a POST request 
* URL: '/users'
### Fetching userID 🤵
* Is a GET request
* URL: '/users/<id here>'
### Creating Post 📭
* Is a POST request
* URL: '/posts'
### Fetching postID 📭
* Is a GET request
* URL: '/posts/<id here>'
### Show user's all posts 📭
* Is a GET request
* URL: '/posts/users/<id here>'

## Dependencies ⚙
All the direct and indirect dependencies required along with the version is listed in go.mod file. The checksum present in go.sum file is used to validate the checksum of each of direct and indirect dependency to confirm that none of them has been modified. 
Imported packages include:
* [Golang Standard Packages](/https://pkg.go.dev/std)
* [Golang Mongo Driver v1.4.0](/https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.4.0)
