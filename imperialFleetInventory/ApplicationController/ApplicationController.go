/*
*   ApplicationController is a main controller of the application.
*   It contain's all the Handler, will routing of valid request to specific controller and discard the invalid requests
 */

package ApplicationController

import (
	"imperialFleetInventory/SpaceshipController"
	"imperialFleetInventory/UserController"
	"net/http"
)

func UserHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		UserController.CreateUser(writer, request)
		return
	case "GET":
		UserController.GetUsersList(writer, request)
		return
	}
	writer.WriteHeader(http.StatusMethodNotAllowed)
	_, _ = writer.Write([]byte("This request method (" + request.Method + ") is not implemented."))
	return
}

/*
* SpaceshipHandler is the controller for SpaceCraft
* has requests method like POST, DELETE, GET and PUT to insert, retrive and update entity into table of db
* otherwise thrown error
 */
func SpaceshipHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		SpaceshipController.CreateSpaceship(writer, request)
		return
	case "DELETE":
		SpaceshipController.DeleteSpaceship(writer, request)
		return
	case "GET":
		SpaceshipController.GetSpaceshipList(writer, request)
		return
	case "PUT":
		SpaceshipController.UpdateSpaceShip(writer, request)
		return
	}
	writer.WriteHeader(http.StatusMethodNotAllowed)
	_, _ = writer.Write([]byte("This request method (" + request.Method + ") is not implemented."))
	return
}

func GetSpaceShipDetailHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		SpaceshipController.GetSpaceship(writer, request)
		return
	}
	writer.WriteHeader(http.StatusMethodNotAllowed)
	_, _ = writer.Write([]byte("This request method (" + request.Method + ") is not implemented."))
	return
}
