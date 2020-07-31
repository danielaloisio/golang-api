package methods

import (
	"context"
	"encoding/json"
	"github.com/golang-api/connection"
	"github.com/golang-api/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	personCollection := connection.PersonCollection()

	ctx := context.Background()

	cur, err := personCollection.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	if err = cur.All(ctx, &People); err != nil {
		panic(err)
	}

	defer cur.Close(ctx)

	jsonReturn, err := json.Marshal(People)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseJSON(jsonReturn, w, http.StatusOK)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	personCollection := connection.PersonCollection()

	params := mux.Vars(r)

	filter := bson.M{"Id": params["id"]}

	err := personCollection.FindOne(context.Background(), filter).Decode(&PeopleOne)

	if err != nil {
		jsonReturn, err := json.Marshal(ErrorMessage{Message: append(ErrorMessage{}.Message, "Id not found")})

		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(jsonReturn)
		return
	}

	jsonReturn, err := json.Marshal(PeopleOne)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseJSON(jsonReturn, w, http.StatusOK)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	personCollection := connection.PersonCollection()

	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = utils.NewGuid()

	checkValid := PersonIsValid(person)

	if len(checkValid.Message) > 0 {
		jsonReturn, err := json.Marshal(checkValid)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ResponseJSON(jsonReturn, w, http.StatusBadRequest)
		return
	}

	_, err := personCollection.InsertOne(context.Background(), bson.M{"Id": person.ID, "Firstname": person.FirstName, "Lastname": person.LastName, "Address": person.Address})

	if err != nil {
		log.Fatal(err)
	}

	jsonResponse, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseJSON(jsonResponse, w, http.StatusOK)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	personCollection := connection.PersonCollection()
	filter := bson.M{"Id": params["id"]}
	personCollection.DeleteOne(context.Background(), filter)

	jsonReturn, err := json.Marshal(MessageApi{Message: "success"})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ResponseJSON(jsonReturn, w, http.StatusOK)
}

func ResponseJSON(b []byte, w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(b)
}
