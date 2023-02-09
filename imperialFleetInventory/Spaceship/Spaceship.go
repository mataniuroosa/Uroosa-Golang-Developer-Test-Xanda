package Spaceship

import (
	"fmt"
	DBUtilities "imperialFleetInventory/Common"

	_ "github.com/go-sql-driver/mysql"
)

/*
*	Struct : Spaceship
*	To model the SpaceCraft entity
 */

type Spaceship struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	Class     string     `json:"class"`
	Crew      int64      `json:"crew"`
	Image     string     `json:"image"`
	Value     float64    `json:"value"`
	Status    string     `json:"status"`
	Spaceship []Armament `json:"armament"`
}

/*
*	Struct : SpaceshipDetailResponse
*	To model the SpaceCraft and Armament entity
*/
type SpaceshipDetailResponse struct {
	Id        int64              `json:"id"`
	Name      string             `json:"name"`
	Class     string             `json:"class"`
	Crew      int64              `json:"crew"`
	Image     string             `json:"image"`
	Value     float64            `json:"value"`
	Status    string             `json:"status"`
	Spaceship []ArmanentResponse `json:"armament"`
}

/*
*	Struct : SpaceshipResponse
*	To model the SpaceCraft 
*/

type SpaceshipResponse struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}

/*
*	Struct : Response
*	To  return the response of success
*/

type Response struct {
	Success bool `json:"success"`
}

/*
*	Struct : Armament
*	To model the Armament entity
*/

type Armament struct {
	Title   string `json:"title"`
	Qty     int64  `json:"quality"`
	SpaceId int64  `json:"spaceid"`
}

/*
*	Struct : ArmanentResponse
*	To Reponse of Armament entity
*/
type ArmanentResponse struct {
	Title string `json:"title"`
	Qty   int64  `json:"quality"`
}
/*
*	CreateNewSpaceCraft(spaceship Spaceship) (Response, error)
*	Input params
*	@spaceship	:	spaceship details wrapped in a spaceship struct
*	Output params
*	@Response	:	display the response success true .
*	@error	:	error in case of any failure else nil.
*
*	Flow
*	1.	Create a DB connection and defer the call to close the connection
*	2.	Insert the record in spaceCraft table and return the last record by using LastInsertId() 
*	3.	Insert the record in Armament by using space_id from lastInsertedId method
*	4.	In case of any error, return error
*	5.	In case of success, return response
*
*/
func CreateNewSpaceCraft(spaceship Spaceship) (Response, error) {

	config, _ := DBUtilities.LoadConfiguration("ApplicationConfigProperties/config.json")
	var dbconn = DBUtilities.GetDbConnection()
	//	we need to close the db handle
	defer DBUtilities.CloseDbConnection(dbconn)

	query, _ := dbconn.Prepare("INSERT into " + config.Database.SpaceCraptTableName + "(space_name, space_class, space_crew, space_image, space_value, space_status) VALUES (?, ?, ?, ?, ?, ?)")
	result, err := query.Exec(spaceship.Name, spaceship.Class, spaceship.Crew, spaceship.Image, spaceship.Value, spaceship.Status)
	if err != nil {
		panic(err.Error())
	}

	lastId, err := result.LastInsertId()
	DBUtilities.CheckDbErrorIsNotNull(err)

	for i := 0; i < len(spaceship.Spaceship); i++ {
		_, err := dbconn.Query("INSERT into "+config.Database.ArmamentTableName+"(ship_id,title,quality) VALUES (?,?, ?)", lastId, spaceship.Spaceship[i].Title, spaceship.Spaceship[i].Qty)
		DBUtilities.CheckDbErrorIsNotNull(err)
	}
	var spacecraftResponse Response
	spacecraftResponse.Success = true
	return spacecraftResponse, err
}

// Update Spaceship record in SpaceCraft Table

func UpdateSpaceCraft(spaceship Spaceship) (Response, error) {
	config, _ := DBUtilities.LoadConfiguration("ApplicationConfigProperties/config.json")
	var dbconn = DBUtilities.GetDbConnection()
	fmt.Println(config.Database.SpaceCraptTableName)

	// Query to update SpaceCraptTableName with new values. Values got from paramter spaceship
	query, err := dbconn.Query("UPDATE "+config.Database.SpaceCraptTableName+" SET space_name = ?, space_class = ?, space_crew = ?, space_image = ?, space_value = ?, space_status = ?  where id = ?", spaceship.Name, spaceship.Class, spaceship.Crew, spaceship.Image, spaceship.Value, spaceship.Status, spaceship.Id)
	DBUtilities.CheckDbErrorIsNotNull(err)
	fmt.Println("jwjw ", query)

	_, err = dbconn.Query("DELETE FROM "+config.Database.ArmamentTableName+" WHERE ship_id = ?", spaceship.Id)
	DBUtilities.CheckDbErrorIsNotNull(err)

	for i := 0; i < len(spaceship.Spaceship); i++ {
		quy, err := dbconn.Query("INSERT into "+config.Database.ArmamentTableName+"(ship_id,title,quality) VALUES (?,?, ?)", spaceship.Id, spaceship.Spaceship[i].Title, spaceship.Spaceship[i].Qty)
		DBUtilities.CheckDbErrorIsNotNull(err)
		fmt.Println(quy)
	}
	DBUtilities.CloseDbConnection(dbconn)

	var spaceData Response
	spaceData.Success = true
	return spaceData, err
}

/*
*	SpaceshipList() ([]SpaceshipResponse, error)
*	Input params
*	Output params
*	@[]AbstractGuest	:	guest list wrapped in AbstractGuest containing all the relevant details only
*	@error	:	error in case of any failure else nil.
*
*	Flow
*	1.	Create a DB connection and defer the call to close the connection
*	2.	Using the db handle run select query to get all the records where id equal to id from url but only desired fields from SpaceCraft table and Armament Table
*	3.	In case of any error, return (empty slice of SpaceshipResponse,error)
*	4.	In case of success, return (valid slice of SpaceshipResponse, nil)
*
 */
func SpaceshipList() ([]SpaceshipResponse, error) {
	config, _ := DBUtilities.LoadConfiguration("ApplicationConfigProperties/config.json")
	var dbconn = DBUtilities.GetDbConnection()
	//	we need to close the db handle
	defer DBUtilities.CloseDbConnection(dbconn)

	var spaceshiplist []SpaceshipResponse

	results, err := dbconn.Query("SELECT id,space_name,space_status FROM " + config.Database.SpaceCraptTableName)
	if err != nil {
		return []SpaceshipResponse{}, err
	}
	for results.Next() {
		var spaceshipresponse SpaceshipResponse
		err = results.Scan(&spaceshipresponse.Id, &spaceshipresponse.Name, &spaceshipresponse.Status)
		DBUtilities.CheckDbErrorIsNotNull(err)
		spaceshiplist = append(spaceshiplist, spaceshipresponse)
	}
	return spaceshiplist, nil
}

/*
*	GetSpaceshipDetails(id int) (SpaceshipDetailResponse, error)
*	Input params
*	@id	:	id of the space_id for which details are asked
*	Output params
*	@SpaceshipDetailResponse:	SpaceCraft and Armament details wrapped in SpaceshipDetailResponse struct
*	@error	:	error in case of any failure else nil.
*
*	Flow
*	1.	Create a DB connection and defer the call to close the connection
*	2.	Using the dbconn handle run select query to get the record corresponding to the spaceship from SpaceCraft and Armament table that we are looking for.
*	3.	In case of any error, return (empty SpaceshipDetailResponse struct,error)
*	4.	In case of success, return (valid SpaceshipDetailResponse struct, nil)
*
 */

func GetSpaceshipDetails(id int) (SpaceshipDetailResponse, error) {
	
	config, _ := DBUtilities.LoadConfiguration("ApplicationConfigProperties/config.json")
	var dbconn = DBUtilities.GetDbConnection()
    // we need to close the db handle
	 defer DBUtilities.CloseDbConnection(dbconn)

	var spaceship SpaceshipDetailResponse
	results, err := dbconn.Query("SELECT id, space_name, space_class, space_crew , space_image, space_value, space_status  FROM "+config.Database.SpaceCraptTableName+" WHERE id = ?", id)

	for results.Next() {
		
		err = results.Scan(&spaceship.Id, &spaceship.Name, &spaceship.Class, &spaceship.Crew, &spaceship.Image, &spaceship.Value, &spaceship.Status)
		DBUtilities.CheckDbErrorIsNotNull(err)
	}

	ArmamentAesults, ArmamentErr := dbconn.Query("SELECT title, quality  FROM "+config.Database.ArmamentTableName+" WHERE ship_id = ?", id)

	var ArmanentResponselist []ArmanentResponse
	for ArmamentAesults.Next() {
		var armanentResponse ArmanentResponse
		err = ArmamentAesults.Scan(&armanentResponse.Title, &armanentResponse.Qty)
		ArmanentResponselist = append(ArmanentResponselist, armanentResponse)
		DBUtilities.CheckDbErrorIsNotNull(err)
	}

	spaceship.Spaceship = ArmanentResponselist

	if err != nil {
		return SpaceshipDetailResponse{}, err
	}
	if ArmamentErr != nil {
		return SpaceshipDetailResponse{}, err
	}
	DBUtilities.CloseDbConnection(dbconn)
	return spaceship, nil
}
/*
* 	IsValidId(id int64)
*	Input params
*	id type of int64
*	Output params
*	bool either true or false
*/
func IsValidId(id int64) bool {
	var idCount = 0
	config, _ := DBUtilities.LoadConfiguration("ApplicationConfigProperties/config.json")
	var dbconn = DBUtilities.GetDbConnection()
	//	we need to close the db handle
	defer DBUtilities.CloseDbConnection(dbconn)
	results, err := dbconn.Query("SELECT count(*) FROM "+config.Database.SpaceCraptTableName+" WHERE id = ?", id)

	for results.Next() {
		err = results.Scan(&idCount)
		DBUtilities.CheckDbErrorIsNotNull(err)
	}

	if idCount > 0 {
		return true
	}

	if err != nil {
		return false
	}

	
	return false
}

/*
*	Delete(id int) (Response, error)
*	Input params
*	@id	:	int value depicting id of a spaceship to be deleted
*	Output params
*	@Response	:	Dispaly the response of success .
*	@error	:	error in case of any failure else nil.
*
*	Flow
*	1.	Create a DB connection and defer the call to close the connection
*	2.	Delete data from the Armament then delete data from SpaceCraft
*	3.	In case of any error, return error
*	4.	In case of success, return response
*
 */
func DeleteSpaceShipRecord(id int) (Response, error) {
	config, _ := DBUtilities.LoadConfiguration("ApplicationConfigProperties/config.json")
	var dbconn = DBUtilities.GetDbConnection()
 	//	we need to close the db handle
	 defer DBUtilities.CloseDbConnection(dbconn)

	_, armamenterr := dbconn.Query("DELETE FROM "+config.Database.ArmamentTableName+" WHERE ship_id = ?", id)
	DBUtilities.CheckDbErrorIsNotNull(armamenterr)
	_, err := dbconn.Query("DELETE FROM "+config.Database.SpaceCraptTableName+" WHERE id = ?", id)
	DBUtilities.CheckDbErrorIsNotNull(err)
	var shipDelete Response
	shipDelete.Success = true
	return shipDelete, nil
}
