
-- +goose Up
CREATE TABLE `photos` (
	`id`            INT(11) NOT NULL AUTO_INCREMENT,
	`path`          VARCHAR(65) NOT NULL,
	`user_id`       INT(11) NOT NULL REFERENCES `users`,
    `created_at`    DATETIME NOT NULL,
	`updated_at`    DATETIME NOT NULL,
	`deleted_at`    DATETIME NULL DEFAULT NULL,

	PRIMARY KEY (`id`) USING BTREE,
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

-- +goose Down
DROP TABLE `photos`;
