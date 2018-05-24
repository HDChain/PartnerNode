package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func showeth(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {

		//queryForm, err := url.ParseQuery(r.URL.RawQuery)
		//if err == nil && len(queryForm["id"]) > 0 {
		//    fmt.Fprintln(w, queryForm["id"][0])
		//}
		//if len(r.Form["id"]) > 0 {
		//    fmt.Fprintln(w, r.Form["id"][0])
		// }
		w.Header().Set("Content-Type", "text/html; charset=gb2312")
		t, err := template.ParseFiles("template/html/showeth.html")
		if err != nil {
			fmt.Fprintf(w, "parse template error: %s", err.Error())
			return
		}
		t.Execute(w, nil)
	} else {
		//fmt.Fprintln(w, r.PostFormValue("id"))
		username := r.Form["username"]
		password := r.Form["password"]
		fmt.Fprintf(w, "username = %s, password = %s", username, password)
	}
}

func showipfs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=gb2312")
	t, err := template.ParseFiles("template/html/showipfs.html")
	if err != nil {
		fmt.Fprintf(w, "parse template error: %s", err.Error())
		return
	}
	t.Execute(w, nil)

}
