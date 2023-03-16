CREATE DATABASE IF NOT EXISTS `dev-book`;
USE `dev-book`;
DROP TABLE IF EXISTS user;
CREATE TABLE USER (
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique ,
    password varchar(20)not null unique ,
    created_at timestamp default current_timestamp(),
    updated_at timestamp default null
) ENGINE=INNODB;