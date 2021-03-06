//Generated by gopages from html/login.html, do not edit
//This file will be overwritten during build

package pages

import (
	"code.google.com/p/gopages/pkg"
	"fmt"
	"github.com/abiosoft/gopages-sample/store"
	"net/http"
)

func Renderhtmlloginhtml(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	writer.Header().Set("Content-Type", "text/html")
	print := func(toPrint ...interface{}) {
		fmt.Fprint(writer, toPrint...)
	}
	formValue := func(keyToGet string) string {
		return request.FormValue(keyToGet)
	}
	_ = print
	_ = formValue // prevent initialization runtime error

	err := ""
	if request.Method == "POST" {
		username := formValue("username")
		password := formValue("password")
		if !store.Validate(username, password) {
			err = "invalid username or password"
		} else {
			store.LoginUser(writer, request, username)
			http.Redirect(writer, request, "/", http.StatusFound)
		}
	}

	fmt.Fprint(writer, `
<html>
	<head> <title>gopages Sample Blog</title></head>
	<body>
		<h3> Login </h3>
		`)
	print(err)
	fmt.Fprint(writer, `
		<form action="/login" method="post">
			Username
			<br />
			<input type="text" name="username" />
			<br />
			Password
			<br />
			<input type="password" name="password" />
			<br />
			<input type="submit" value="Login" />
		</form>
	</body>
</html>`)

}
func init() {
	gopages.ParsedPaths["pages/htmlloginhtml.go"] = "html/login.html"
}
