package SpaceshipController

import (

	DBUtilities "imperialFleetInventory/Common"
	"imperialFleetInventory/Spaceship"
	"imperialFleetInventory/User"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"encoding/json"
	"fmt"
)

/*
*	API : POST /spaceship
*
*	CreateSpaceship(w http.ResponseWriter, r *http.Request)
*	Input params
*	@http.ResponseWriter	:	standard response writer object
*	@http.Request			:	standard http request object
*
*	Flow
*	1. Valid the user is authenticate or not 
*	2. Check for content-type of the request body
*	3. Check for the request body validity
*	4. Call the Guest.CreateNewSpaceCraft() with the InputModel params to create a record in the table's
*
*/

func CreateSpaceship(w http.ResponseWriter, r *http.Request) {
	username, password, _ := r.BasicAuth()
	isUserAuthenticated := User.AuthenticateUser(username, password)

	if !isUserAuthenticated {
		_, _ = w.Write([]byte(fmt.Sprintf("Unauthenticated User due to %s",isUserAuthenticated)))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

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

	var spaceship Spaceship.Spaceship
	inputJsonModel := spaceship
	err = json.Unmarshal(bodyBytes, &inputJsonModel)

	var insertedData Spaceship.Response
	//  Calling the Create method for creating the record and return response
	insertedData, err = Spaceship.CreateNewSpaceCraft(inputJsonModel)

	if err != nil {
		errResponse := make(map[string]string)
		jsonbytes, _ := json.Marshal(errResponse)
		w.Write(jsonbytes)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Display the response and error 
	jsonBytes, err := json.Marshal(insertedData)
	if err != nil {
		_, _ = w.Write([]byte(fmt.Sprintf("Failed to insert the records due to %s", err.Error())))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("content-type", "application/json")
	_, _ = w.Write(jsonBytes)
	w.WriteHeader(http.StatusOK)
	return 
}

/*
*	API : GET /spaceship_list
*
*	GetSpaceshipList(w http.ResponseWriter, r *http.Request)
*	Input params
*	@http.ResponseWriter	:	standard response writer object
*	@http.Request			:	standard http request object
*
*	Flow
*	1. Valid user is authenticated or not 
*	2. Fetch all Spaceship records using Spaceship.SpaceshipList()
*	3. create a map to create response in desired format
*	4. set the content-type as application/json
*	5. return with statusOK in success
*
 */

func GetSpaceshipList(writer http.ResponseWriter, request *http.Request) {
	username, password, _ := request.BasicAuth()

	isUserAuthenticated := User.AuthenticateUser(username, password)

	if !isUserAuthenticated {
		_, _ = writer.Write([]byte(fmt.Sprintf("Unauthenticated User due to %s",isUserAuthenticated)))
		writer.WriteHeader(http.StatusUnauthorized)
		return
	
	}

	spaceshiplistdata, err := Spaceship.SpaceshipList()
	if err != nil {
		_, _ = writer.Write([]byte(fmt.Sprintf("Error Details %s", err.Error())))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	spaceshiplist := make(map[string][]Spaceship.SpaceshipResponse)
	spaceshiplist["data"] = spaceshiplistdata
	jsonBytes, err := json.Marshal(spaceshiplist)

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

/*
*	API : Get /spaceships/id
*
*	GetSpaceship(w http.ResponseWriter, r *http.Request)
*	Input params
*	@http.ResponseWriter	:	standard response writer object
*	@http.Request			:	standard http request object
*
*	Flow
*	1. Valid user is authenticated or not 
*	2. Get id from the the Url and assign to GetSpaceshipDetails method
*	3. Get all the records from the SpaceCraft and Armament table, where id = urlid.
*	4. create a response in the desired format
*	5. return the response with StatusOK
*
 */

func GetSpaceship(writer http.ResponseWriter, request *http.Request) {

	username, password, _ := request.BasicAuth()

	isUserAuthenticated := User.AuthenticateUser(username, password)

	if !isUserAuthenticated {
		_, _ = writer.Write([]byte(fmt.Sprintf("Unauthenticated User due to %s",isUserAuthenticated)))
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}

	id := request.URL.Query().Get("id")

	intId, err := strconv.Atoi(id)
	spaceShipDetail, err := Spaceship.GetSpaceshipDetails(intId)

	if err != nil {
		_, _ = writer.Write([]byte(fmt.Sprintf("Error Details %s", err.Error())))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(spaceShipDetail)

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

/*
*	API : DELETE /spaceship_d/{id}
*
*	DeleteSpaceship(writer http.ResponseWriter, request *http.Request)
*	Input params
*	@http.ResponseWriter	:	standard response writer object
*	@http.Request			:	standard http request object
*
*	Flow
*	1. Valid user is authenticated or not 
*	2. Fetch the 'id' param from the URL
*	3. Delete record with particular id from the table's
*	4. return the response with StatusOK
*
 */


func DeleteSpaceship(writer http.ResponseWriter, request *http.Request) {

	username, password, _ := request.BasicAuth()

	isUserAuthenticated := User.AuthenticateUser(username, password)

	if !isUserAuthenticated {
		_, _ = writer.Write([]byte(fmt.Sprintf("Unauthenticated User due to %s",isUserAuthenticated)))
		writer.WriteHeader(http.StatusUnauthorized)
		return
	}
	id := request.URL.Query().Get("id")

	intId, err := strconv.Atoi(id)

	deletespaceShipDetail, err := Spaceship.DeleteSpaceShipRecord(intId)

	if err != nil {
		_, _ = writer.Write([]byte(fmt.Sprintf("Error Details %s", err.Error())))
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(deletespaceShipDetail)

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

/*
*	API : PUT /spaceship_u/{id}
*
*	UpdateSpaceShip(w http.ResponseWriter, r *http.Request)
*	Input params
*	@http.ResponseWriter	:	standard response writer object
*	@http.Request			:	standard http request object
*
*	Flow
*	1. Valid the user is authenticate or not
*	2. Check for content-type of the request body
*	3. Check for the request body validity
*	4. Check if the id is valid id using IsValidId()
*	5. If Valid id, update the SpaceCraft details and delete the Armament table id by particular id.
*	6. Check if table can accommodate any change in the SpaceCraft table along with the armament table. If yes, allow and create/update record in SpaceCraft, else return error with response.
*	7. return statusOK if success else return error with message.
*
 */
func UpdateSpaceShip(w http.ResponseWriter, r *http.Request) {

	username, password, _ := r.BasicAuth()

	isUserAuthenticated := User.AuthenticateUser(username, password)

	if !isUserAuthenticated {
		jsonBytes, _ := json.Marshal("Unauthenticated User")
		_, _ = w.Write(jsonBytes)
		w.WriteHeader(http.StatusUnauthorized)

		return
	}
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

	var spaceship Spaceship.Spaceship
	inputJsonModel := spaceship
	err = json.Unmarshal(bodyBytes, &inputJsonModel)

	isValidId := Spaceship.IsValidId(inputJsonModel.Id)
	if isValidId {
		insertedData, err := Spaceship.UpdateSpaceCraft(inputJsonModel)
		if err != nil {
			errResponse := make(map[string]string)
			jsonbytes, _ := json.Marshal(errResponse)
			w.Write(jsonbytes)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonBytes, err := json.Marshal(insertedData)
		w.Header().Add("content-type", "application/json")
		_, _ = w.Write(jsonBytes)
		w.WriteHeader(http.StatusOK)

		return

	} else {
		var spaceData Spaceship.Response
		spaceData.Success = false
		jsonBytes, _ := json.Marshal(spaceData)
		w.Header().Add("content-type", "application/json")
		_, _ = w.Write(jsonBytes)
		w.WriteHeader(http.StatusOK)

		return
	}
}


// Get SpaceShip Details by Id
func getGuestDetails(id int) (Spaceship.SpaceshipDetailResponse, error) {
	return Spaceship.GetSpaceshipDetails(id)
}

// CheckValidSpaceShip record
func CheckValidSpaceShip(id int) bool {
	_, err := getGuestDetails(id)
	DBUtilities.CheckDbErrorIsNotNull(err)
	return true
}