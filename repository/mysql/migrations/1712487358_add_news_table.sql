-- +migrate Up
CREATE TABLE news(
                     id INT PRIMARY KEY AUTO_INCREMENT,
                     title VARCHAR(191) NOT NULL,
                     short_description VARCHAR(191) NOT NULL,
                     description VARCHAR(191) NOT NULL UNIQUE,
                     image_file_name VARCHAR(191) NOT NULL,
                     creator_user_id INT NOT NULL,
                     visit_count INT NOT NULL,
                     like_count INT NOT NULL,
                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE news;