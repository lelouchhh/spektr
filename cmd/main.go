package main

import (
	"fmt"
	"html/template"
	"net/http"
	_ "net/smtp"
)

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		tmpl, _ := template.ParseFiles("../web/html/kuibishev/error.html")
		tmpl.Execute(w, nil)
	}

}

func barabinsk_index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/barabinsk/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/barabinsk/index.html")
	tmpl.Execute(w, nil)
}

func barabinsk_tariffs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/barabinsk/tariffs" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/barabinsk/tariff.html")
	tmpl.Execute(w, nil)
}

func barabinsk_contacts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/barabinsk/contacts" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/barabinsk/contacts.html")
	tmpl.Execute(w, nil)
}

func barabinsk_news(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/barabinsk/news" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/barabinsk/news.html")
	tmpl.Execute(w, nil)
}



func kuibishev_index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/kuibishev/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/kuibishev/index.html")
	tmpl.Execute(w, nil)
}

func kuibishev_tariffs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/kuibishev/tariffs" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/kuibishev/tariff.html")
	tmpl.Execute(w, nil)
}

func kuibishev_contacts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/kuibishev/contacts" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/kuibishev/contacts.html")
	tmpl.Execute(w, nil)
}

func kuibishev_news(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/kuibishev/news" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/kuibishev/news.html")
	tmpl.Execute(w, nil)
}

func new1(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/news/10-11-2020/1" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/kuibishev/new1.html")
	tmpl.Execute(w, nil)
}

func celinnoe_index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/celinnoe/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/celinnoe/index.html")
	tmpl.Execute(w, nil)
}

func celinnoe_tariffs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/celinnoe/tariffs" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/celinnoe/tariff.html")
	tmpl.Execute(w, nil)
}

func celinnoe_contacts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/celinnoe/contacts" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/celinnoe/contacts.html")
	tmpl.Execute(w, nil)
}

func celinnoe_news(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/celinnoe/news" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/celinnoe/news.html")
	tmpl.Execute(w, nil)
}




func krasnogorskoe_index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/krasnogorskoe/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/krasnogorskoe/index.html")
	tmpl.Execute(w, nil)
}

func krasnogorskoe_tariffs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/krasnogorskoe/tariffs" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/krasnogorskoe/tariff.html")
	tmpl.Execute(w, nil)
}

func krasnogorskoe_contacts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/krasnogorskoe/contacts" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/krasnogorskoe/contacts.html")
	tmpl.Execute(w, nil)
}

func krasnogorskoe_news(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/krasnogorskoe/news" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/krasnogorskoe/news.html")
	tmpl.Execute(w, nil)
}


func chani_index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/chani/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/chani/index.html")
	tmpl.Execute(w, nil)
}

func chani_tariffs(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/chani/tariffs" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/chani/tariff.html")
	tmpl.Execute(w, nil)
}

func chani_contacts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/chani/contacts" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/chani/contacts.html")
	tmpl.Execute(w, nil)
}

func chani_news(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/chani/news" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("../web/html/chani/news.html")
	tmpl.Execute(w, nil)
}
func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	http.Redirect(w, r, "kuibishev/", 301)
}


func newUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	Send("Привет" + r.Form["name"][0] + "\nМы перезвоним вам на " +r.Form["phone"][0])
	fmt.Println(r.Form["name"])
	fmt.Println(r.Form["phone"])
	fmt.Println(r.Form["address"])
	tmpl, _ := template.ParseFiles("../web/html/kuibishev/newuser.html")
	tmpl.Execute(w, r.Form)
}
func main() {
	http.HandleFunc("/kuibishev/", kuibishev_index)
	http.HandleFunc("/kuibishev/tariffs", kuibishev_tariffs)
	http.HandleFunc("/kuibishev/contacts", kuibishev_contacts)
	http.HandleFunc("/kuibishev/news", kuibishev_news)

	http.HandleFunc("/barabinsk/", barabinsk_index)
	http.HandleFunc("/barabinsk/tariffs", barabinsk_tariffs)
	http.HandleFunc("/barabinsk/contacts", barabinsk_contacts)
	http.HandleFunc("/barabinsk/news", barabinsk_news)

	http.HandleFunc("/celinnoe/", celinnoe_index)
	http.HandleFunc("/celinnoe/tariffs", celinnoe_tariffs)
	http.HandleFunc("/celinnoe/contacts",celinnoe_contacts)
	http.HandleFunc("/celinnoe/news", celinnoe_news)

	http.HandleFunc("/krasnogorskoe/", krasnogorskoe_index)
	http.HandleFunc("/krasnogorskoe/tariffs", krasnogorskoe_tariffs)
	http.HandleFunc("/krasnogorskoe/contacts", krasnogorskoe_contacts)
	http.HandleFunc("/krasnogorskoe/news", krasnogorskoe_news)

	http.HandleFunc("/chani/", chani_index)
	http.HandleFunc("/chani/tariffs", chani_tariffs)
	http.HandleFunc("/chani/contacts", chani_contacts)
	http.HandleFunc("/chani/news", chani_news)

	http.HandleFunc("/kuibishev/newuser", newUser)
	http.HandleFunc("/news/10-11-2020/1", new1)

	http.HandleFunc("/", index)




	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("../web/"))))
	server := &http.Server{
		Addr: "localhost:80",
	}
	server.ListenAndServe()
}
