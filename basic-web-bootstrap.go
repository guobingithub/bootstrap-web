/*
 * @Author: bingo
 * @Date: 2018-11-16 10:55:02
 */
package main

import "fmt"
import "net/http"
import "html/template"

func Hello(response http.ResponseWriter, request *http.Request) {
	type person struct {
		Id      int
		Name    string
		Country string
	}

	liumiaocn := person{Id: 1001, Name: "liumiaocn", Country: "China"}

	tmpl, err := template.ParseFiles("./views/user.tpl")
	if err != nil {
		fmt.Println("Error happened..")
	}
	tmpl.Execute(response, liumiaocn)
}

func First_bts(response http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("./views/first_bts.tpl")
	if err != nil {
		fmt.Println("Error happened..")
	}
	tmpl.Execute(response, nil)
}

func First_web_page(response http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("./views/first_webpage.tpl")
	if err != nil {
		fmt.Println("Error happened..")
	}
	tmpl.Execute(response, nil)
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

	http.HandleFunc("/", Hello)
	http.HandleFunc("/first",First_bts)
	http.HandleFunc("/webpage",First_web_page)

	http.ListenAndServe(":8080", nil)
}
