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
		fmt.Fprintf(w, "%v<hr /><br />", thing)
	}

	form := `
  <hr />
  <form action="/say" method="POST">
  Say What?
  <input type="text" name="say" value="???"><br>
  <input type="submit" value="Submit">
</form>
`

  fmt.Fprint(w, form)
}

func sayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("say handler!")
	session := getSession(w, r)
	if session == nil {
		return
	}

	fmt.Println(fmt.Printf("%#v", r))
	r.ParseForm()
	fmt.Println(fmt.Printf("%#v", r.Form))

	name := session.Values["name"]
	if name == "" || name == nil {
		name = "Anon"
	}

	thingSaid := r.Form["say"][0]
	if thingSaid != "" {
		things = append(things, fmt.Sprintf("%v: %v", name, thingSaid))
	}

  http.Redirect(w, r, "/", 303)
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
