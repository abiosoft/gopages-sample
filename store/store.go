package store

import (
	"code.google.com/p/gorilla/sessions"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"net/http"
	"time"
)

var Database *mgo.Database
var Sessions = sessions.NewCookieStore([]byte("some-secret-keys"))

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	Database = session.DB("gopages-sample")
	//create an index for the username field on the users collection
	if err := Database.C("users").EnsureIndex(mgo.Index{
		Key:    []string{"username"},
		Unique: true,
	}); err != nil {
		panic(err)
	}
}

type User struct {
	Username string
	Password string
}

func AddUser(username, password string) error {
	user := &User{
		username,
		password,
	}
	return Database.C("users").Insert(user)
}

func GetUser(username string) *User {
	var u *User
	err := Database.C("users").Find(bson.M{"username": username}).One(&u)
	if err != nil {
		return nil
	}
	return u
}

func Validate(username, password string) bool {
	var u *User
	err := Database.C("users").Find(bson.M{"username": username}).One(&u)
	if err == nil {
		return u.Username == username && u.Password == password
	}
	return false
}

type Post struct {
	Name    string
	Title   string
	Content string
	Time    time.Time
}

func AddPost(name, title, content string) error {
	post := &Post{
		name,
		title,
		content,
		time.Now(),
	}
	return Database.C("posts").Insert(post)
}

func GetPosts() []*Post {
	var posts []*Post
	err := Database.C("posts").Find(nil).Sort("-time").All(&posts)
	if err != nil {
		return nil
	}
	return posts
}

func LoginUser(w http.ResponseWriter, r *http.Request, username string) {
	s, _ := Sessions.Get(r, "gopages-sample")
	s.Values["username"] = username
	s.Save(r, w)
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	s, _ := Sessions.Get(r, "gopages-sample")
	user := CurrentUser(r)
	if user != nil {
		delete(s.Values, "username")
		s.Save(r, w)
	}
}

func CurrentUser(r *http.Request) *User {
	s, _ := Sessions.Get(r, "gopages-sample")
	username, ok := s.Values["username"]
	if ok {
		return GetUser(username.(string))
	}
	return nil
}
