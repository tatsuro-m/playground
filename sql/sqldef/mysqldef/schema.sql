CREATE TABLE `user`
(
    `id`   bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(191) DEFAULT 'k0kubun',
    created_at DATETIME NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

ALTER TABLE user ADD INDEX index_name(name);
