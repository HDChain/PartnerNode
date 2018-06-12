package httpfile

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main3() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	server := &http.Server{
		Addr:    ":4040",
		Handler: http.DefaultServeMux,
	}

	quitChan := make(chan os.Signal)
	signal.Notify(quitChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	go func() {
		fmt.Println(<-quitChan)
		server.Close()
	}()

	go server.ListenAndServe()

	time.Sleep(2 * time.Second)

	quitChan <- syscall.SIGINT

	time.Sleep(1 * time.Second)
}
