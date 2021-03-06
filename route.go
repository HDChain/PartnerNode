package main

import (
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func adminHandler(w http.ResponseWriter, r *http.Request) {
	// 获取cookie
	cookie, err := r.Cookie("admin_name")
	if err != nil || cookie.Value == "" {
		http.Redirect(w, r, "/login/index", http.StatusFound)
	}

	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	admin := &adminController{}
	controller := reflect.ValueOf(admin)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	userValue := reflect.ValueOf(cookie.Value)
	method.Call([]reflect.Value{responseValue, requestValue, userValue})
}

func ajaxHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("ajaxHandler")
	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	ajax := &ajaxController{}
	controller := reflect.ValueOf(ajax)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}

	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("loginHandler")
	pathInfo := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}

	login := &loginController{}
	controller := reflect.ValueOf(login)
	method := controller.MethodByName(action)
	if !method.IsValid() {
		method = controller.MethodByName(strings.Title("index") + "Action")
	}
	requestValue := reflect.ValueOf(r)
	responseValue := reflect.ValueOf(w)
	method.Call([]reflect.Value{responseValue, requestValue})
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		//http.Redirect(w, r, "/showeth", http.StatusFound)
		w.Header().Set("Content-Type", "text/html; charset=gb2312")
		t, err := template.ParseFiles("template/html/index.html")
		if err != nil {
			log.Println(err)
		}
		t.Execute(w, nil)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=gb2312")
	t, err := template.ParseFiles("template/html/404.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}
