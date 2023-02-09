package User

import (
	

	DBUtilities "imperialFleetInventory/Common"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

/*
*	Struct : User
*	To model the User entity
 */

type User struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type UserResponse struct {
	User string `json:"user"`
}

type Response struct {
	Success bool `json:"success"`
}



func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/*
**	Creates and inserts a User record to User table
**	Input Params
**	@user : user type struct which contain the user and password and used as user identifier
**	Return Param
**  @Response : returns response success
**	@error : returns error if there is any else returns nil
**
**	This Method is based on the following understanding
**
 */

func AddUser(user User) (Response, error) {

	config, _ := DBUtilities.LoadConfiguration("ApplicationConfigProperties/config.json")
	var dbconn = DBUtilities.GetDbConnection()
	defer DBUtilities.CloseDbConnection(dbconn)

	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		 panic(err.Error())
	}
	
	//	Insert the new entry into the database 
	_, dberr := dbconn.Query("INSERT into "+config.Database.AuthenticateUserTableName+"(user, hash) VALUES (?, ?)", user.User, string(bytes))
	//	we need to close the db handle
	
	var spaceData Response
	spaceData.Success = true
	return spaceData, dberr
}

// Check the user are authorized or not 
func AuthenticateUser(user string, password string) bool {
	var userHash = ""
	config, _ := DBUtilities.LoadConfiguration("ApplicationConfigProperties/config.json")
	var dbconn = DBUtilities.GetDbConnection()

	results, err := dbconn.Query("SELECT hash FROM "+config.Database.AuthenticateUserTableName+" WHERE user = ?", user)

	for results.Next() {
		err = results.Scan(&userHash)
		DBUtilities.CheckDbErrorIsNotNull(err)
	}
	if CheckPasswordHash(password, userHash) {
		return true
	}
	if err != nil {
		return false
	}
	//	we need to close the db handle
	DBUtilities.CloseDbConnection(dbconn)
	return false
}

// Get All users from database 
func GetUsers() ([]UserResponse, error) {

	config, _ := DBUtilities.LoadConfiguration("ApplicationConfigProperties/config.json")
	var dbconn = DBUtilities.GetDbConnection()
	var userlist []UserResponse

	//	Get list of user from the database 
	results, err := dbconn.Query("SELECT user FROM " + config.Database.AuthenticateUserTableName)
	if err != nil {
		return []UserResponse{}, err
	}
	for results.Next() {
		var user UserResponse
		err = results.Scan(&user.User)
		DBUtilities.CheckDbErrorIsNotNull(err)
		userlist = append(userlist, user)
	}
	//	we need to close the db handle
	DBUtilities.CloseDbConnection(dbconn)
	return userlist, err
}
