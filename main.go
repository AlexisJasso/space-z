package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
	"strings"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

var things []string

func getSession(w http.ResponseWriter, r *http.Request) (*sessions.Session) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}

	return session
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("view handler!")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	for _, thing := range things {
		fmt.Fprintf(w, "%v<br />", thing)
	}
}

func sayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("say handler!")
	session := getSession(w, r)
	if session == nil {
		return
	}

	name := session.Values["name"]
	if name == "" || name == nil {
		name = "Anon"
	}

	parts := strings.Split(r.URL.Path, "/")
	thingSaid := strings.Join(parts[2:], " ")
	if thingSaid != "" {
		things = append(things, fmt.Sprintf("%v: %v", name, thingSaid))
	}

	viewHandler(w, r)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login handler!")
	session := getSession(w, r)

	parts := strings.Split(r.URL.Path, "/")
	name := parts[2]

	session.Values["name"] = name
	session.Save(r, w)

	fmt.Fprintf(w, "Logged you in, %v!", name)
}

func baseHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	base := parts[1]

	fmt.Printf("%v -> %v\n", r.URL.Path, base)

	if strings.ToLower(base) == "login" {
		loginHandler(w, r)
	} else if strings.ToLower(base) == "say" {
		sayHandler(w, r)
	} else {
		viewHandler(w, r)
	}
}

func main() {
	http.HandleFunc("/", baseHandler)
	http.ListenAndServe(":8080", nil)
}
