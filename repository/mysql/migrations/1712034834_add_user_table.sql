-- +migrate Up
CREATE TABLE user(
    id INT PRIMARY KEY AUTO_INCREMENT,
    first_name VARCHAR(191) NOT NULL,
    last_name VARCHAR(191) NOT NULL,
    phone_number VARCHAR(191) NOT NULL UNIQUE,
    email VARCHAR(191) NOT NULL,
    register_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE user;