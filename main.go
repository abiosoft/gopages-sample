package main

import (
	"code.google.com/p/gopages/pkg"
	"net/http"

	_ "github.com/abiosoft/gopages-sample/pages" //required for initialization
	"github.com/abiosoft/gopages-sample/store"
)

func main() {
	http.HandleFunc("/", gopages.Handler("html/index.html"))
	http.HandleFunc("/login", gopages.Handler("html/login.html"))
	http.HandleFunc("/register", gopages.Handler("html/register.html"))
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/post", post)
	http.ListenAndServe(":9999", nil)
}

func post(w http.ResponseWriter, r *http.Request) {
	user := store.CurrentUser(r)
	title := r.FormValue("title")
	content := r.FormValue("content")
	if user != nil && title != "" && content != "" {
		store.AddPost(user.Username, title, content)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func logout(w http.ResponseWriter, r *http.Request) {
	store.LogoutUser(w, r)
	http.Redirect(w, r, "/", http.StatusFound)
}
