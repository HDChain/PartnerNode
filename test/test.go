package test

import (
	"crypto/rand"
	"fmt"
	"html/template"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"

	"archive/zip"
	"io"
	"log"
	"strconv"
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
const UploadPath = "./tmp"

func UploadFileHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validate file size
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			renderError(w, "FILE_TOO_BIG", http.StatusBadRequest)
			return
		}
		// parse and validate file and post parameters
		fileType := r.PostFormValue("type")
		file, _, err := r.FormFile("uploadFile")
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
		// check file type, detectcontenttype only needs the first 512 bytes
		filetype := http.DetectContentType(fileBytes)
		if filetype != "image/jpeg" && filetype != "image/jpg" &&
			filetype != "image/gif" && filetype != "image/png" &&
			filetype != "application/pdf" {
			renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
			return
		}

		fileName := randToken(12)
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
	//path := "/root/down.txt"
	w.Header().Set("Content-Disposition", "attachment; filename=WHATEVER_YOU_WANT")
	path := "/00Inbox/Inbox.txt"

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
