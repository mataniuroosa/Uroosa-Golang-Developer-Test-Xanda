package UserController

import (
	"encoding/json"
	"fmt"
	"imperialFleetInventory/User"
	"io/ioutil"
	"net/http"
	"strings"
)

// insert data into the user table using API And display the response
func CreateUser(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	bodyBytes, err := ioutil.ReadAll(r.Body)
	fmt.Println(bodyBytes, parts)
	defer r.Body.Close()
	if err != nil {
		
	}
	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("need content-type 'application/json', but got '%s'", ct)))
		return
	}

	var user User.User
	inputJsonModel := user
	err = json.Unmarshal(bodyBytes, &inputJsonModel)

	var insertedData User.Response

	if len( strings.Trim(inputJsonModel.Password, " ") ) == 0 {
		jsonbytes, _ := json.Marshal("Password not valid")
		w.Header().Add("content-type", "application/json")
		_, _ = w.Write(jsonbytes)
		w.WriteHeader(http.StatusBadRequest)
		return
	}	

	insertedData, err = User.AddUser(inputJsonModel)

	if err != nil {
		jsonbytes, _ := json.Marshal("Username Already Exist")
		w.Header().Add("content-type", "application/json")
		_, _ = w.Write(jsonbytes)
		w.WriteHeader(http.StatusAlreadyReported)
		return
	}

	jsonBytes, err := json.Marshal(insertedData)
	w.Header().Add("content-type", "application/json")
	_, _ = w.Write(jsonBytes)
	w.WriteHeader(http.StatusOK)
	return
}


// Get List of spaceship records
func GetUsersList(writer http.ResponseWriter, request *http.Request) {
	userdata, err := User.GetUsers()
	if err != nil {
		_, _ = writer.Write([]byte(fmt.Sprintf("Error Details %s", err.Error())))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	userlist := make(map[string][]User.UserResponse)
	userlist["data"] = userdata
	jsonBytes, err := json.Marshal(userlist)

	if err != nil {
		_, _ = writer.Write([]byte(fmt.Sprintf("Error Details %s", err.Error())))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Add("content-type", "application/json")
	_, _ = writer.Write(jsonBytes)
	writer.WriteHeader(http.StatusOK)
	return
}
