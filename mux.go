package main

import (
	"github.com/gorilla/mux"
	"mule/apps/auth"
	"mule/apps/menagerie"
	"mule/apps/roomgame"
	"mule/mylog"
	"net/http"
)

var (
	ErrF   = mylog.ErrF
	Log    = mylog.Log
	ErrLog = mylog.Err
)

func SetupMux() *mux.Router {
	r := mux.NewRouter()
	r.StrictSlash(true)
	// Static Stuff //
	r.HandleFunc("/favicon.ico", FavIcon)
	r.PathPrefix("/_static/").Handler(http.StripPrefix("/_static/", http.FileServer(http.Dir("static/"))))
	// Not Static //
	r.HandleFunc("/", auth.DataWrap(MainPage)).Name("main")
	auth.SetupMux(r)
	r.HandleFunc("/{user}", auth.DataWrap(UserPage)).Name("user")
	// APPS //
	s1 := r.PathPrefix("/{user}/" + menagerie.APPNAME).Subrouter()
	menagerie.SetupMux(s1)
	s2 := r.PathPrefix("/{user}/" + roomgame.APPNAME).Subrouter()
	roomgame.SetupMux(s2)
	// //
	r.HandleFunc("/{user}/{non-app:.*}", auth.DataWrap(auth.RedirHome))
	return r
}
