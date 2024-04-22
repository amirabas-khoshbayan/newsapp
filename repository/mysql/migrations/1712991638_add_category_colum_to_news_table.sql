-- +migrate Up
ALTER TABLE `news` ADD  COLUMN `categories` ENUM ('business','health','sports','politics') NOT NULL;

-- +migrate Down
ALTER TABLE `news` DROP COLUMN `categories`;
