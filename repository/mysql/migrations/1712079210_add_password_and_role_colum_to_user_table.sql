-- +migrate Up
ALTER TABLE `user` ADD COLUMN `password` VARCHAR(191) NOT NULL;
ALTER TABLE `user` ADD COLUMN `role` VARCHAR(191) NOT NULL;

-- +migrate Down
ALTER TABLE `user` DROP COLUMN `password`;
ALTER TABLE `user` DROP COLUMN `role`;