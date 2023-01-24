package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	coll, _, ctx := DbConnection()
	tmpl, err := template.ParseFiles(path.Join("templates", "index.html"))
	if err != nil {
		panic(err)
	}
	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			http.Error(w, " Sorry something is wrong", http.StatusBadRequest)
			return
		}
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := getHash([]byte(r.FormValue("password")))
		gender := r.FormValue("gender")

		fmt.Printf("Name: %s\n Email: %s\n Password:%s\n Gender:%s\n ", name, email, password, gender)
		user := User{
			Name:     name,
			Email:    email,
			Password: password,
			Gender:   gender,
		}
		fmt.Println(user)
		insertResult, err := coll.InsertOne(ctx, &user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Inserted a single document: ", insertResult.InsertedID)
		http.Redirect(w, r, "/login", http.StatusMovedPermanently)
	} else {
		err = tmpl.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	}

}
