package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	//task "./task"
	"archive/zip"
	"strconv"

	"crypto/rand"
	"io/ioutil"
	"mime"
	"path/filepath"
)

var staticHandler http.Handler

// 静态文件处理
func StaticServer(w http.ResponseWriter, req *http.Request) {
	//path := "/root/down.txt"
	w.Header().Set("Content-Disposition", "attachment; filename=WHATEVER_YOU_WANT")
	path := "/00Inbox/Inbox.txt"

	http.ServeFile(w, req, path)
}

func zipHandler(rw http.ResponseWriter, r *http.Request) {
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
	url = "http://127.0.0.1:1789/src/qq.exe"
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

func hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello"))
}

type Handlers struct {
}

func (h *Handlers) ResAction(w http.ResponseWriter, req *http.Request) {
	fmt.Println("res")
	w.Write([]byte("res"))
}

func say(w http.ResponseWriter, req *http.Request) {
	pathInfo := strings.Trim(req.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	var action = ""
	fmt.Println(strings.Join(parts, "|"))
	if len(parts) > 1 {
		action = strings.Title(parts[1]) + "Action"
	}
	fmt.Println(action)
	handle := &Handlers{}
	controller := reflect.ValueOf(handle)
	method := controller.MethodByName(action)
	r := reflect.ValueOf(req)
	wr := reflect.ValueOf(w)
	method.Call([]reflect.Value{wr, r})
}

const maxUploadSize = 2 * 1024 * 2014 // 2 MB
const uploadPath = "./tmp"

func uploadFileHandler() http.HandlerFunc {
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

		newPath := filepath.Join(uploadPath, fileName+fileEndings[0])
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

func main() {
	/*
		do := task.NewInter(task.NewTask())
		do.OnInit()
		log.Output(1, "log")
	*/
	http.HandleFunc("/upload", uploadFileHandler())

	fs := http.FileServer(http.Dir(uploadPath))
	http.Handle("/files/", http.StripPrefix("/files", fs))

	http.HandleFunc("/zipdownload", zipHandler)
	http.HandleFunc("/download", StaticServer)
	http.HandleFunc("/hello", hello)
	http.Handle("/handle/", http.HandlerFunc(say))
	log.Println("Listening...")
	log.Print("Server started on localhost:8080, use /upload for uploading files and /files/{fileName} for downloading files.")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	select {} //阻塞进程

}
