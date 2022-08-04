package actions

import (
	"encoding/json"
	"github/go-backend/model"
	"net/http"

	//"strconv"
	"github/go-backend/database"

	"github.com/gorilla/mux"
)

func Healthy(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Healthy"))
}

func GetHeroes(w http.ResponseWriter, r *http.Request) {

	var heroes model.Heroes = database.Get()

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// w.Header = http.Header{
	// 	"Content-Type":  {"application/json"},
	// 	"Access-Control-Allow-Origin": {"*"},
	// 	"Access-Control-Allow-Methods": {"GET,POST,PUT,PATCH,DELETE"},
	// 	//"Access-Control-Allow-Methods": {"Content-Type','Authorization"},
	// }

	json.NewEncoder(w).Encode(heroes.GetAllBHeroes().Heroes)
}

func GetHeroByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var heroes model.Heroes = database.Get()
	strVar := params["id"]
	book := heroes.GetHeroByID(strVar)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(book)
}

// func AddHero(w http.ResponseWriter, r *http.Request) {
// 	var book model.Book
// 	var heroes model.Heroes = database.Get()
// 	var hero model.Hero
// 	_ = json.NewDecoder(r.Body).Decode(&hero)
// 	if hero.SuperHero != "" {
// 		newBook := model.AddHero(ID, SuperHero, Publisher, Alter_Ego, First_apparance, Characters)
// 		json.NewEncoder(w).Encode(newBook)
// 	}
// }

// func UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	var book model.Book
// 	intVar, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		json.NewEncoder(w).Encode(&model.Book{})
// 		return
// 	}
// 	_ = json.NewDecoder(r.Body).Decode(&book)
// 	book.ID = intVar
// 	newBook := model.UpdateBook(book.ID, book.Title, book.Author)
// 	json.NewEncoder(w).Encode(newBook)
// }

// func DeleteBook(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	intVar, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		json.NewEncoder(w).Encode(&model.Book{})
// 		return
// 	}
// 	book := model.DeleteBook(intVar)
// 	bookAux := model.Book{}
// 	if book != bookAux {
// 		w.WriteHeader(200)
// 	} else {
// 		w.WriteHeader(404)
// 	}
// }

// func strToInt(w http.ResponseWriter, r *http.Request) int {
// 	params := mux.Vars(r)
// 	intVar, err := strconv.Atoi(params["int"])
// 	if err != nil {
// 		json.NewEncoder(w).Encode(&model.Book{})
// 		return -1
// 	}
// 	return intVar
// }
