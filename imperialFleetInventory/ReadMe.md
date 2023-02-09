## Step's to run the project on machine 
##SETUP
1. Install golang on the machine
2. Download/clone this project to the _/go/src_ 
3. Install mysql 5.7.x on the local machine to setup db server (Ignore this step if you are planning to place the DB on the cloud)
4. Run _dbQueriesConsoleSetup.sql_ script file on mysql dashboard or mysql terminal
5. Run go build command then exec file is created.
6. Run ./imperialFleetInventory then it's running.
5. You can query at localhost:8080/ to access the APIs (OR change the configuration if hosting the code on application server)

## Assignment Task
You are R3-D3 and were just appointed the general of the imperial fleet. Your first action as the new general is to digitalise the imperial fleet inventory.

You know that each spacecraft has the following characteristics:

Name
Class
Armament
Crew
Image
Value
Status
You need to create a galactic database (using MySQL) that stores all the spacecraft’s details.

Then create a galactic application programming interface (REST API) GoLang.
Create 
Update 
Delete 

** Bonus : Bonus points if the create, update and delete interface can only be performed by authorised members of the imperial fleet (authenticated users)

#### Solution 

I implement the core functionality of the task along with the bonus point in my Code base.Following are the details of each part

#Data Models
I implement the data model using MySQL

1. SpaceCraft
    { id, ship_name, ship_class, ship_crew, ship_image, ship_value, ship_status }
What : To represent a SpaceCraft entity.

2.  Armament
    { id, ship_id, title, quantity }
What : To represent a Armament entity with particular SpaceCraft.

3.  AuthenticateUser
    { user, hash}
What : To represent a AuthenticateUser entity.

##Workflows
1. Add Spaceship Entity to SpaceCraft.
 API : `POST /spaceship`. Create a record in SpaceCraft Table.
2. Fetch the List of Spaceship.
 API : `GET /spaceship_list`. Read all records from SpaceCraft and return id, name and status.
3. Fetch the Single Record of Spaceship by using Id.
 API : `GET /spaceships/{id}`. Read single record with particular spaceship from SpaceCraft and return Id, Name, Class, Crew, Image, Value Status and Armament[].
4. Delete Particular Spaceship record by using Id.
  API : `DELETE /spaceship_d/{id}`. Mark departure by updating the ispresent field of the record.
5. Update the record of spaceship.
   API : `PUT /spaceship_u/`. Update the record in SpaceCraft table and Armanent table.

###File Significance
1. _routes.go_ : To map the APIs to respective methods.
2. _applicationController.go_ : Common interface that will route the incoming request to respective Controller
3. _UserController_ : This controller will the interface to communicate to user entity.
7. _User.go_ : This is Entity model of AuthenticateUser table . It wraps all the operations done on User Entity.
4. _SpaceShipController_ : This controller will the interface to communicate to SpaceCraft entity and Armament entity .
7. _SpaceShip.go_ : This is Entity model of SpaceCraft table and Armament. It wraps all the operations done on SpaceCraft Entity and Armament Entity.
8. _server.go_ : This the **entry point of the service** containing main() method.


###Database Setup
1. _dbQueriesConsoleSetup.sql_ : This contains mysql scripts to setup the DB for the project to work.

###API's Response 
Sample data and output are attached in pdf file







