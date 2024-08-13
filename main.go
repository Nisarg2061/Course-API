package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

// Model for the file/database
type Book struct {
	BookId    int     `json:"bookid"`
	BookName  string  `json:"bookname"`
	BookPrice float32 `json:"bookprice"`
	Author    *Author `json:"author"`
}

type Author struct {
	Fname   string `json:"fname"`
	Website string `json:"website"`
}

// Fake database
var authors []Author

var books []Book

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func IsEmpty(c *Book) bool {
	return c.BookName == ""
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by me!<h1>"))
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	book, exist := loadBooks()[r.PathValue("id")]

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func loadBooks() map[string]Book {
	res := make(map[string]Book, len(books))

	for _, x := range books {
		res[strconv.Itoa(x.BookId)] = x
	}

	return res
}

func createBook(w http.ResponseWriter, r *http.Request, param string) {
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	if IsEmpty(&book) {
		json.NewEncoder(w).Encode("Enter valid data.")
		return
	}
	if param == "" {
		book.BookId = rand.Intn(100)
	}
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

func createBookRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Enter valid data.")
	}

	createBook(w, r, "")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param, err := strconv.Atoi(r.PathValue("id"))
	p := r.PathValue("id")
	check(err)
	for i, j := range books {
		if j.BookId == param {
			books = append(books[:i], books[i+1:]...)
			createBook(w, r, p)
		}
	}
}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "json/application")
	param, err := strconv.Atoi(r.PathValue("id"))
	check(err)
	for i, j := range books {
		if j.BookId == param {
			books = append(books[:i], books[i+1:]...)
			json.NewEncoder(w).Encode("Book Deleted")
			break
		}
	}
}

func main() {
	router := http.NewServeMux()

	authors = append(authors, Author{Fname: "Nisarg", Website: "bynisarg.in"})
	authors = append(authors, Author{Fname: "Bhupendra", Website: "bybhupendra.in"})

	books = append(books, Book{BookId: 1, BookName: "Introduction to Golang", BookPrice: 399.99, Author: &authors[0]})
	books = append(books, Book{BookId: 2, BookName: "Beginners Guide to JavaScript", BookPrice: 399.99, Author: &authors[1]})
	books = append(books, Book{BookId: 3, BookName: "Getting Started with Docker", BookPrice: 399.99, Author: &authors[0]})

	router.HandleFunc("GET /", serveHome)
	router.HandleFunc("GET /books", getAllBooks)
	router.HandleFunc("GET /book/{id}", getBook)
	router.HandleFunc("POST /book", createBookRoute)
	router.HandleFunc("PUT /book/{id}", updateBook)
	router.HandleFunc("DELETE /book/{id}", deleteBook)

	server := http.Server{
		Addr:    ":4000",
		Handler: router,
	}

	fmt.Println("Listening on port 4000...")

	err := server.ListenAndServe()
	check(err)
}
