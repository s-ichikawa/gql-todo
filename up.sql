CREATE DATABASE IF NOT EXISTS gql_todo;

USE gql_todo;

CREATE TABLE IF NOT EXISTS users (
  id   VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS todos (
  id      VARCHAR(255) NOT NULL,
  user_id VARCHAR(255) NOT NULL,
  text    VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
)