package main

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../ui/html/index.html", "../ui/html/header.html", "../ui/html/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/news_portal")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//select data from table
	res, err := db.Query("SELECT * FROM `posts` ORDER BY id DESC LIMIT 10")
	if err != nil {
		panic(err)
	}

	posts = []Article{}
	for res.Next() {
		var article Article
		err = res.Scan(&article.Id, &article.Title, &article.Anons, &article.Content, &article.CategoryId, &article.CreatedDate)
		if err != nil {
			panic(err)
		}

		posts = append(posts, article)
	}

	t.ExecuteTemplate(w, "index", posts)
}

func students(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../ui/html/students.html", "../ui/html/header.html", "../ui/html/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/news_portal")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//select data from table
	res, err := db.Query("SELECT * FROM `posts` WHERE `category_id` = 1 ORDER BY `id` DESC")
	if err != nil {
		panic(err)
	}

	posts = []Article{}
	for res.Next() {
		var article Article
		err = res.Scan(&article.Id, &article.Title, &article.Anons, &article.Content, &article.CategoryId, &article.CreatedDate)
		if err != nil {
			panic(err)
		}

		posts = append(posts, article)
	}

	t.ExecuteTemplate(w, "students", posts)
}

func staff(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../ui/html/staff.html", "../ui/html/header.html", "../ui/html/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/news_portal")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//select data from table
	res, err := db.Query("SELECT * FROM `posts` WHERE `category_id` = 2 ORDER BY `id` DESC")
	if err != nil {
		panic(err)
	}

	posts = []Article{}
	for res.Next() {
		var article Article
		err = res.Scan(&article.Id, &article.Title, &article.Anons, &article.Content, &article.CategoryId, &article.CreatedDate)
		if err != nil {
			panic(err)
		}

		posts = append(posts, article)
	}

	t.ExecuteTemplate(w, "staff", posts)
}

func applicants(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../ui/html/applicants.html", "../ui/html/header.html", "../ui/html/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/news_portal")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//select data from table
	res, err := db.Query("SELECT * FROM `posts` WHERE `category_id` = 3 ORDER BY `id` DESC")
	if err != nil {
		panic(err)
	}

	posts = []Article{}
	for res.Next() {
		var article Article
		err = res.Scan(&article.Id, &article.Title, &article.Anons, &article.Content, &article.CategoryId, &article.CreatedDate)
		if err != nil {
			panic(err)
		}

		posts = append(posts, article)
	}

	t.ExecuteTemplate(w, "applicants", posts)
}

func researches(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../ui/html/researches.html", "../ui/html/header.html", "../ui/html/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/news_portal")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//select data from table
	res, err := db.Query("SELECT * FROM `posts` WHERE `category_id` = 4 ORDER BY `id` DESC ")
	if err != nil {
		panic(err)
	}

	posts = []Article{}
	for res.Next() {
		var article Article
		err = res.Scan(&article.Id, &article.Title, &article.Anons, &article.Content, &article.CategoryId, &article.CreatedDate)
		if err != nil {
			panic(err)
		}

		posts = append(posts, article)
	}

	t.ExecuteTemplate(w, "researches", posts)
}

func contacts(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../ui/html/contacts.html", "../ui/html/header.html", "../ui/html/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "contacts", nil)
}

func saveArticle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	content := r.FormValue("content")
	category_id, _ := strconv.Atoi(r.FormValue("category_id"))

	if title == "" || anons == "" || content == "" || category_id == 0 {
		fmt.Fprintf(w, "Not all data is filled in")
	} else {
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/news_portal")
		if err != nil {
			panic(err)
		}
		defer db.Close()

		//insert data to the table
		insert, err := db.Query(fmt.Sprintf("INSERT INTO `posts` (`title`,`anons`,`content`,`category_id`) "+
			"VAlUES('%s','%s','%s','%d')", title, anons, content, category_id))
		if err != nil {
			panic(err)
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../ui/html/createArticle.html", "../ui/html/header.html", "../ui/html/footer.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "createArticle", nil)
}

func showPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("../ui/html/show.html", "../ui/html/header.html", "../ui/html/footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/news_portal")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	res, err := db.Query(fmt.Sprintf("SELECT * FROM `posts` WHERE `id` = '%s'", vars["id"]))
	if err != nil {
		panic(err)
	}

	singlePost = Article{}
	for res.Next() {
		var article Article
		err = res.Scan(&article.Id, &article.Title, &article.Anons, &article.Content, &article.CategoryId, &article.CreatedDate)
		if err != nil {
			panic(err)
		}

		singlePost = article
	}

	t.ExecuteTemplate(w, "show", singlePost)
}

func handleFunc() {
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/contacts", contacts).Methods("GET")
	router.HandleFunc("/save_article", saveArticle).Methods("POST")
	router.HandleFunc("/createArticle", createArticle).Methods("GET")
	router.HandleFunc("/students", students).Methods("GET")
	router.HandleFunc("/staff", staff).Methods("GET")
	router.HandleFunc("/applicants", applicants).Methods("GET")
	router.HandleFunc("/researches", researches).Methods("GET")
	router.HandleFunc("/post/{id:[0-9]+}", showPost).Methods("GET")

	http.Handle("/", router)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	http.Handle("/static/", http.StripPrefix("/static", fileServer))
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", router)
	log.Fatal(err)
}
