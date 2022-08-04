package database

import (
	"encoding/json"
	"fmt"
	"github/go-backend/model"
	"io/ioutil"
	"os"
)

func Get() model.Heroes {
	jsonFile, err := os.Open("heroes.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var heroes model.Heroes
	json.Unmarshal(byteValue, &heroes)
	return heroes
}