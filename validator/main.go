package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/validator/bin/handlers"
	"github.com/validator/bin/utility"
)

type activehandler func(w http.ResponseWriter, r *http.Request, usermap handlers.Hashtable)

var usermap = make(handlers.Hashtable)

func main() {

	conf, err := utility.CreatConfig("config.json")

	if err != nil{
		log.Fatal(err.Error())
	}

	router := mux.NewRouter()
	router.HandleFunc("/regestry", handlers.Regestry)
	router.HandleFunc("/unregestry", handlers.UnRegestry)
	router.Handle("/login", activehandler(handlers.LogIn))
	router.Handle("/logout", activehandler(handlers.LogOut))

	fmt.Printf("Запускаем сервер: Хост - "+conf.Host+"; Порт - "+conf.Port)

	httpServ := &http.Server{Addr: conf.Host+":"+conf.Port, Handler: router}

	go httpServ.ListenAndServe()

	chanal := make(chan os.Signal, 1)
	signal.Notify(chanal, os.Interrupt)
	<- chanal

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	httpServ.Shutdown(ctx)

}

func (api activehandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	api(w, r, usermap)

}