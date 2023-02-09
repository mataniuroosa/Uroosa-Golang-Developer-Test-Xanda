/*
* The SQL Script to create database and table's 
* Database name : imperialfleetinventory
* Table Name : SpaceCraft
* Table Name : Armament
*/

create database if not exists imperialfleetinventory; 
use imperialfleetinventory;

create table if not exists SpaceCraft(id int auto_increment primary key,space_name varchar(20) not null,space_class varchar(20) not null, space_crew int not null,space_image varchar(500) not null, space_value float not null, space_status varchar(30));
create table if not exists Armament(id int auto_increment primary key, ship_id varchar(20) not null, title varchar(20) not null, quality int not null);
create table if not exists AuthenticateUser ( user varchar(20) primary key, hash varchar(500) not null );



-- insert into TableInfo (spacename, class, crew,spaceimage, spacevalue, status) values ("Devastator","Star Destroyer",35000, "https:\\url.to.image",1999.99,"operational");

