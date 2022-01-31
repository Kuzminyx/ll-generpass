package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/validator/bin/utility"
)

//Структура предназначена для хранения детальной информации о пользователях получивших доступ пройдя валидацию...
type User struct {
	Name         string `json:"name"`
	Publicid     int 
	Secrethashid string `json:"hashid"`
	Seed         string `json:"seed"`
	Createdate time.Time
}

//Мапа для временного хранения данных авторизованных пользователей...
type Hashtable map[string]User

func LogIn(w http.ResponseWriter, r *http.Request, usermap Hashtable) {

	user := User{}

	err := readJson(r, &user)

	if err != nil {
		fmt.Println(err.Error())
		utility.SendJSON(w, err.Error(), 400)
	}

	_, ok := usermap[user.Name]

	if ok {

		utility.SendJSON(w, "Вы были ранее авторизованы", 200)

	}else{

		user.Createdate = time.Now()
		usermap[user.Name] = user

		utility.SendJSON(w, "Вы авторизованы", 200)

	}

}

func LogOut(w http.ResponseWriter, r *http.Request, usermap Hashtable) {

}

func Regestry(w http.ResponseWriter, r *http.Request){

}

func UnRegestry(w http.ResponseWriter, r *http.Request){

}

func readJson(r *http.Request, user *User) error {

	decode := json.NewDecoder(r.Body)

	alert := decode.Decode(&user)

	if alert != nil {
		return alert
	}

	return nil

}