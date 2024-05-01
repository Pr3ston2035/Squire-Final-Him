CREATE TABLE users (
                       id INT PRIMARY KEY,
                       nickname VARCHAR(50) NOT NULL,
                       email VARCHAR(100) NOT NULL,
                       password VARCHAR(255) NOT NULL


);

CREATE TABLE mainpage (

                          url VARCHAR(255) NOT NULL,
                          user VARCHAR(255) NOT NULL,
                          password VARCHAR(255) NOT NULL
)

