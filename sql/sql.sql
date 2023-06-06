CREATE DATABASE IF NOT EXISTS `book`;

USE `book`;

DROP TABLE IF EXISTS FOLLOWERS;
DROP TABLE IF EXISTS PUBLICATION;
DROP TABLE IF EXISTS USER;

CREATE TABLE USER (
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    password varchar(100) not null,
    created_at timestamp default current_timestamp(),
    updated_at timestamp default null
) ENGINE=INNODB;

CREATE TABLE FOLLOWERS (
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

CREATE TABLE PUBLICATION (
    id int auto_increment primary key,
    title varchar(50) not null,
    auth_id int not null,
    FOREIGN KEY (auth_id) REFERENCES USER(id) ON DELETE CASCADE,
    description VARCHAR(300),
    content varchar(500),
    `like` int default 0,
    created_at timestamp default current_timestamp(),
    updated_at timestamp default null
) ENGINE=INNODB;
