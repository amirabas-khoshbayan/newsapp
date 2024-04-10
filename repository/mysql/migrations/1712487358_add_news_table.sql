-- +migrate Up
CREATE TABLE news(
                     id INT PRIMARY KEY AUTO_INCREMENT,
                     title VARCHAR(191) NOT NULL,
                     short_description VARCHAR(191) NOT NULL,
                     description VARCHAR(191) NOT NULL,
                     image_file_name VARCHAR(191),
                     creator_user_id INT NOT NULL,
                     visit_count INT,
                     like_count INT,
                     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down
DROP TABLE news;