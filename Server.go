package main

import (
	"fmt"
	"html/template"
	//"io/ioutil"
	"log"
	"net/http"

	"./ethrpc"
	"./ipfs"
)

//----------------------------------------------------------------

type Handlers struct {
}

//http://localhost:8081/eth/?id=1&name=abc
func ethfunc(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	if req.Method == "GET" {
		if len(req.Form["id"]) > 0 {
			val := req.Form["id"][0]
			if val == "0" {
				// 获取函数名 和 参数
				valfunc := req.Form["func"][0]
				valparam1 := req.Form["param1"][0]
				valparam2 := req.Form["param2"][0]
				log.Println(valfunc + " " + valparam1 + " " + valparam2 + " ")
				result, err := ethrpc.JRpcCall(valfunc)
				if err != nil {
					log.Println(err)
					log.Println(result)
				}
				//fmt.Fprintln(w, "调用成功 ", result)

				EthOutputJson(w, 1, "调用成功", result, nil)
			}
		}
	}
}

func ipfsfunc(w http.ResponseWriter, req *http.Request) {
	//http://localhost:8080/ipfs/QmaEswFNsf3D3Sjb2Qy9kstEZvWyL7bSknoudaxLgdFK4J
	req.ParseForm()
	if req.Method == "GET" {
		if len(req.Form["id"]) > 0 {
			//fmt.Fprintln(w, req.Form["key"][0])
			key := req.Form["id"][0]
			if key == "0" {
				//httpGet()
				//fmt.Fprintf(w, "httpGet()")
				//url := "http://localhost:8080/ipfs/QmaEswFNsf3D3Sjb2Qy9kstEZvWyL7bSknoudaxLgdFK4J"
				//http.Redirect(w, req, url, http.StatusFound)

				//ipfs.Ipfsmain()

				w.Write([]byte("SUCCESS"))
			}
			if key == "1" {
				hash := req.Form["hash"][0]

				//httpGet()
				//fmt.Fprintf(w, "httpGet()")
				//url := "http://localhost:8080/ipfs/QmaEswFNsf3D3Sjb2Qy9kstEZvWyL7bSknoudaxLgdFK4J"
				//http.Redirect(w, req, url, http.StatusFound)

				ipfs.IpfsGet(hash, "./upload/a12.txt")

				//w.Header().Set("Content-Type", "text/html; charset=UTF-8")
				//w.Write([]byte("SUCCESS,   <h3> <font color='red'><a href='/download'> 下载 </a></font></h3> <h3> <font color='red'><a href='/files/a12.txt'>  显示 </a></font></h3><a href='/showipfs'>   返回 </a> "))

				w.Header().Set("Content-Type", "text/html; charset=UTF-8")
				t, err := template.ParseFiles("template/html/retipfsdownload.html")
				if err != nil {
					fmt.Fprintf(w, "parse template error: %s", err.Error())
					return
				}
				t.Execute(w, nil)
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

	webmain()

	log.Println("Listening...")
	log.Print("Server started on localhost:8081, use /upload for uploading files and /files/{fileName} for downloading files.")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	select {} //阻塞进程

}
