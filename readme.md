docker-compose exec database /bin/bash

CREATE DATABASE IF NOT EXISTS todolist;
USE todolist;

CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON todolist TO 'user'@'localhost';