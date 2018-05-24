package main

import (
	"log"
	"net/http"

	"./test"
)

func webmain() {
	log.Println("main")

	http.HandleFunc("/html/pics/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/showpic", test.Showpic)

	http.HandleFunc("/showeth", showeth)
	http.HandleFunc("/showipfs", showipfs)

	http.Handle("/pics/", http.FileServer(http.Dir("template")))
	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))

	http.HandleFunc("/admin/", adminHandler)
	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/ajax/", ajaxHandler)
	http.HandleFunc("/", NotFoundHandler)
	//http.ListenAndServe(":8888", nil)

	http.HandleFunc("/upload", test.UploadFileHandler())
	fs := http.FileServer(http.Dir(test.UploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	http.HandleFunc("/zipdownload", test.ZipHandler)
	http.HandleFunc("/download", test.StaticServer)

}
