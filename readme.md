# Todo App
The todo app is a Golang application designed to be a example of Go, Docker and Kubernetes with a React Typrescript frontend.

## !!!! DO NOT USE README !!!!
Currently I am changiung to use Kubernetes as well as create a good developer environment, I will update the readme when everything is working again
[Skaffold](https://skaffold.dev/docs/)

## Getting Setup
Currently to get started, you will need to run go through the following
1. [Download Docker](https://docs.docker.com/v17.12/install/)
2. [Download Go](https://golang.org/doc/install)
3. Create an .env file with the following
```
DB_ROOT_USER=root
DB_ROOT_PW=root
DB_NAME=todolist
DB_USER=user
DB_PW=password
DB_URLuser:password@/todolist
```
4. Run the following base code
```bash
# create database and create bash shell
docker-compose up database
docker-compose exec database /bin/bash

# login to mysql using default login
mysql -u root -p
# default password is root

# not needed for local development, but always good practice,
# obviously use your own password
ALTER USER 'root'@'localhost' IDENTIFIED BY 'PASSWORD';

# create database and user
CREATE DATABASE IF NOT EXISTS todolist;
USE todolist;

CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON todolist TO 'user'@'localhost';

# exit mysql
exit
# exit container
exit
```
4. Stop the database container using ctrl + c or
```bash
docker stop [CONTAINER_ID]
```

## Starting backend server
In order to start the backend server, run the command
```bash
docker-compose build && docker-compose up
```
This will build images of the database and the server and then start them

## Starting the frontend server
Coming soon, probably when there is a frontend server
