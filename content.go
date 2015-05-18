package main

import (
	"mule/apps/auth"
	"mule/helpers"
	"net/http"
)

const APPNAME = "main"

var (
	mixTem = helpers.GenTMixer(APPNAME)
	mainT  = mixTem("frame", "main")
	userT  = mixTem("frame", "user")
)

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
