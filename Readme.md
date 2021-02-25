## Entity User
 ID int
 Username string 
 Password string 
 Nama Lengkap string
 Foto string
##  SQL DB 
 CREATE DATABASE test_backend;
 USE test_backend;
 CREATE TABLE users (
   id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
   username varchar(30) NOT NULL,
   password varchar(255) NOT NULL,
   nama_lengkap varchar(255) NOT NULL,
   foto varchar(255) NOT NULL
 )

 ## ENVIRONMENT File
- Create backend.env file
- Values :
APP_PORT=
DB_HOST=
DB_PORT=
DB_NAME=
DB_USERNAME=
DB_PASSWORD=
