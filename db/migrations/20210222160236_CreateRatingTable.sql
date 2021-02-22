
-- +goose Up
CREATE TABLE `ratings` (
	`id`         INT(11) NOT NULL AUTO_INCREMENT,
	`score`      FLOAT(10) NOT NULL,
	`photo_id`   INT(11) NOT NULL REFERENCES `photos`,
    `user_id`    INT(11) NOT NULL REFERENCES `users`,
    `created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NOT NULL,
	`deleted_at` DATETIME NULL DEFAULT NULL,

	PRIMARY KEY (`id`) USING BTREE,
    FOREIGN KEY (`photo_id`) REFERENCES `photos`(`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

-- +goose Down
DROP TABLE `ratings`
