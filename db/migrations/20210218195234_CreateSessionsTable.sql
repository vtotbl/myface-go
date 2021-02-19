
-- +goose Up
CREATE TABLE `sessions` (
	`id`            INT(11) NOT NULL AUTO_INCREMENT,
	`refresh_token` VARCHAR(65) NOT NULL COLLATE 'utf8mb4_general_ci',
	`expires_at`    DATETIME NOT NULL,
	`user_id`       INT(11) NOT NULL REFERENCES `users`,
    `created_at`    DATETIME NOT NULL,
	`updated_at`    DATETIME NOT NULL,
	`deleted_at`    DATETIME NULL DEFAULT NULL,

	PRIMARY KEY (`id`) USING BTREE,
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
    UNIQUE INDEX `user_id` (`user_id`)
);

-- +goose Down
DROP TABLE `sessions`;

