package main

import (
	"imperialFleetInventory/ApplicationController"
	DBUtilities "imperialFleetInventory/Common"
	"net/http"
)

/*
* start Function Definition:
* All the handle function's and specific route
* if any error occur thrown a panic error
 */

func initilize() {
	/*
	*	Routing the requests for different endpoints to Different handlers via Controller
	*	UserHandler	: To handle requests for '/user' requests
	*	SpaceshipHandler	: To handle requests for '/spaceship', '/spaceship_list', '/spaceship_d' and '/spaceship_u' requests
	*	GetSpaceShipDetailHandler		: To handle requests for '/spaceships/'.
	*	And started to listen on 'localhost:8080'
	 */
	http.HandleFunc("/user", ApplicationController.UserHandler)
	http.HandleFunc("/spaceship", ApplicationController.SpaceshipHandler)
	http.HandleFunc("/spaceship_list", ApplicationController.SpaceshipHandler)
	http.HandleFunc("/spaceships/", ApplicationController.GetSpaceShipDetailHandler)
	http.HandleFunc("/spaceship_d", ApplicationController.SpaceshipHandler)
	http.HandleFunc("/spaceship_u", ApplicationController.SpaceshipHandler)

	
	config, _ := DBUtilities.LoadConfiguration("ApplicationConfigProperties/config.json")
	anyError := http.ListenAndServe(config.SERVER_HOST+":"+config.SERVER_PORT, nil)
	if anyError != nil {
		panic(anyError.Error())
	}
}
