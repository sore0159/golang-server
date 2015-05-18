package main

import (
	"fmt"
	"html/template"
	"mule/apps/auth"
	"net/http"
)

var (
	mainT = mixTem("frame", "main")
	userT = mixTem("frame", "user")
)

const TEMP_TEMP = "templates/main/%s.html"

func mixTem(tmpls ...string) *template.Template {
	new_list := make([]string, 0)
	for _, val := range tmpls {
		new_list = append(new_list, fmt.Sprintf(TEMP_TEMP, val))
	}
	t, err := template.New("t").ParseFiles(new_list...)
	if err == nil {
		return t
	}
	panic(err)
}

func MainPage(d *auth.Data) {
	d.ExeT(mainT, "frame", nil)
}

func UserPage(d *auth.Data) {
	if d.HomeURL != d.R.URL.Path {
		d.GoHome()
		return
	}
	d.ExeT(userT, "frame", nil)
}

func FavIcon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon.ico")
}
