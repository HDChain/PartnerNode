package main

import (
	"fmt"
	//"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"./ethrpc"
)

//----------------------------------------------------------------

type Handlers struct {
}

//http://localhost:8081/eth/?id=1&name=abc
func ethfunc(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if req.Method == "GET" {
		if len(req.Form["id"]) > 0 {
			//fmt.Fprintln(w, req.Form["id"][0])
			val := req.Form["id"][0]
			if val == "0" {
				ethrpc.Jrpccall()
				fmt.Fprintln(w, "调用成功")
			}
		}
	}
}

//http://localhost:8080/ipfs/QmaEswFNsf3D3Sjb2Qy9kstEZvWyL7bSknoudaxLgdFK4J
func httpGet() {
	resp, err := http.Get("http://localhost:8080/ipfs/QmaEswFNsf3D3Sjb2Qy9kstEZvWyL7bSknoudaxLgdFK4J")
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func ipfsfunc(w http.ResponseWriter, req *http.Request) {
	//http://localhost:8080/ipfs/QmaEswFNsf3D3Sjb2Qy9kstEZvWyL7bSknoudaxLgdFK4J
	req.ParseForm()
	if req.Method == "GET" {
		if len(req.Form["key"]) > 0 {
			//fmt.Fprintln(w, req.Form["key"][0])
			key := req.Form["key"][0]
			if key == "0" {
				//httpGet()
				//fmt.Fprintf(w, "httpGet()")
				url := "http://localhost:8080/ipfs/QmaEswFNsf3D3Sjb2Qy9kstEZvWyL7bSknoudaxLgdFK4J"
				http.Redirect(w, req, url, http.StatusFound)
			}

		}
		/*
			w.Header().Set("Content-Type", "text/html; charset=gb2312")
			t, err := template.ParseFiles("template/html/showpic.html")
			if err != nil {
				fmt.Fprintf(w, "parse template error: %s", err.Error())
				return
			}
			t.Execute(w, nil)

		*/
	}

}

func main() {

	http.Handle("/ipfs/", http.HandlerFunc(ipfsfunc))
	http.Handle("/eth/", http.HandlerFunc(ethfunc))

	log.Println("Listening...")
	log.Print("Server started on localhost:8081, use /upload for uploading files and /files/{fileName} for downloading files.")

	webmain()

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	select {} //阻塞进程

}
