
-- +goose Up
CREATE TABLE `sessions` (
	`id`            INT(11) NOT NULL AUTO_INCREMENT,
	`refresh_token` VARCHAR(50) NOT NULL COLLATE 'utf8mb4_general_ci',
	`expires_at`    DATETIME NOT NULL,
	`user_id`       INT(11) NOT NULL REFERENCES `users`,

	PRIMARY KEY (`id`) USING BTREE,
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

-- +goose Down
DROP TABLE `sessions`;

