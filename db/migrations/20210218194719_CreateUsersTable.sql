
-- +goose Up
CREATE TABLE `users` (
	`id`       INT(11) NOT NULL AUTO_INCREMENT,
	`login`    VARCHAR(50) NOT NULL COLLATE 'utf8mb4_general_ci',
	`password` VARCHAR(255) NOT NULL COLLATE 'utf8mb4_general_ci',
	`sex`      VARCHAR(50) NOT NULL COLLATE 'utf8mb4_general_ci',
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NOT NULL,
	`deleted_at` DATETIME NULL DEFAULT NULL,

	PRIMARY KEY (`id`) USING BTREE,
	UNIQUE INDEX `login` (`login`)
);

-- +goose Down
DROP TABLE `users`;
