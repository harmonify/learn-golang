CREATE DATABASE belajar_golang_restful_api;

CREATE DATABASE belajar_golang_restful_api_test;

CREATE TABLE belajar_golang_restful_api.category (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
) ENGINE = InnoDB;

CREATE TABLE belajar_golang_restful_api_test.category (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
) ENGINE = InnoDB;