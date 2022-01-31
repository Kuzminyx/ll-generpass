package utility

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Структура для инициализации конфига из фйла конфигурации...
type Config struct {
	Logpath string `json:"logpath"`
	Host string `json:"host"`
	Port string `json:"port"`
	Livetimetohour int `json:"livetimetohour"`
}

func CreatConfig(path string) (Config, error) {

	Conf := Config{}

	err := readConfigFile(path, &Conf)

	if err != nil{
		return Conf, err
	}

	return Conf, nil

}

func readConfigFile(path string, conf *Config) error {

	jsdata, openerr := ioutil.ReadFile(path)

	if openerr != nil {
		return openerr
	}

	errunmarshal := json.Unmarshal(jsdata, &conf)

	if errunmarshal != nil {
		return errunmarshal
	}

	return nil

}

func SendJSON(w http.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	encode := json.NewEncoder(w)
	encode.Encode(msg)
}