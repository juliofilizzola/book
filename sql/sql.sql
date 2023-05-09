CREATE DATABASE IF NOT EXISTS `dev-book`;

USE `dev-book`;

DROP TABLE IF EXISTS user;

DROP TABLE IF EXISTS FOLLOWERS;

CREATE TABLE USER (
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    created_at timestamp default current_timestamp(),
    updated_at timestamp default null
) ENGINE=INNODB;

CREATE TABLE FOLLOWERS(
    user_id int not null,
    FOREIGN KEY (user_id)
    REFERENCES USER(id)
    ON DELETE CASCADE,

    follower_id int not null,
    FOREIGN KEY (follower_id)
    REFERENCES USER(id)
    ON DELETE CASCADE,

    PRIMARY KEY (user_id, follower_id)
) ENGINE=INNODB;