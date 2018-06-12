package httpfile

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"io/ioutil"
	//"mime"
	"net/http"
	"os"
	//	"path/filepath"

	"archive/zip"
	//	"crypto/md5"
	"io"
	"log"
	"strconv"
	//	"time"

	"../ipfs"
)

func Showpic(w http.ResponseWriter, r *http.Request) {
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
		t, err := template.ParseFiles("template/html/showpic.html")
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

//------------------------------------------------------
const maxUploadSize = 2 * 1024 * 2014 // 2 MB
const UploadPath = "./upload"

func UploadFileHandler() http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validate file size
		/*
			r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
			if err := r.ParseMultipartForm(maxUploadSize); err != nil {
				renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
				return
			}
			// parse and validate file and post parameters
			fileType := r.PostFormValue("type")
			file, handler, err := r.FormFile("uploadfile")
			if err != nil {
				renderError(w, "INVALID_FILE", http.StatusBadRequest)
				return
			}
			defer file.Close()
			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				renderError(w, "INVALID_FILE", http.StatusBadRequest)
				return
			}
		*/
		// check file type, detectcontenttype only needs the first 512 bytes
		/*	filetype := http.DetectContentType(fileBytes)
			if filetype != "image/jpeg" && filetype != "image/jpg" &&
				filetype != "image/gif" && filetype != "image/png" &&
				filetype != "application/pdf" {
				renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
				return
			}
			log.Println(handler.Filename)
			log.Println(fileType)
			log.Println("---")
		*/
		//fileName := randToken(12)
		//fileEndings := ".txt"
		/*
			fileEndings, err := mime.ExtensionsByType(fileType)
			if err != nil {
				renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
				return
			}

			newPath := filepath.Join(UploadPath, fileName+fileEndings[0])
			fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)

			// write file
			newFile, err := os.Create(newPath)
			if err != nil {
				renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
				return
			}
			defer newFile.Close()
			if _, err := newFile.Write(fileBytes); err != nil {
				renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
				return
			}
			w.Write([]byte("SUCCESS"))
		*/

		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在upload目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

		//调用 ipfs api，添加到ipfs，获取 hash 返回
		Filehash, err := ipfs.IpfsAdd("./upload/" + handler.Filename)
		if err != nil {
			w.Header().Set("Content-Type", "text/html; charset=UTF-8")
			w.Write([]byte("添加失败 , <a href='/showipfs'> 返回 </a> "))
		} else {
			//w.Header().Set("Content-Type", "text/html; charset=UTF-8")
			//w.Write([]byte("SUCCESS, 文件hash: <h3> <font color='red'> " + Filehash + "</font></h3><a href='/showipfs'>   返回 </a> "))
			type filedata struct {
				Hash string
			}

			filedata1 := filedata{Hash: Filehash}

			w.Header().Set("Content-Type", "text/html; charset=UTF-8")
			t, err := template.ParseFiles("template/html/retipfsupload.html")
			if err != nil {
				fmt.Fprintf(w, "parse template error: %s", err.Error())
				return
			}
			t.Execute(w, filedata1)

		}
	})
}
func renderError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(message))
}
func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

//----------------------------------------------------

var staticHandler http.Handler

// 静态文件处理
func StaticServer(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Disposition", "attachment; filename=a12.txt")

	path := "./upload/a12.txt"

	http.ServeFile(w, req, path)
}

func ZipHandler(rw http.ResponseWriter, r *http.Request) {
	zipName := "ZipTest.zip"
	rw.Header().Set("Content-Type", "application/zip")
	rw.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipName))
	err := getZip(rw)
	if err != nil {
		log.Fatal(err)
	}
}
func getZip(w io.Writer) error {
	zipW := zip.NewWriter(w)
	defer zipW.Close()

	for i := 0; i < 5; i++ {
		f, err := zipW.Create(strconv.Itoa(i) + ".txt")
		if err != nil {
			return err
		}
		_, err = f.Write([]byte(fmt.Sprintf("Hello file %d", i)))
		if err != nil {
			return err
		}
	}
	return nil
}

var (
	url = "http://127.0.0.1:1789/src/abc.exe"
)

func httpdownfile() {
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("qq.exe")
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
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
