CREATE TABLE IF NOT EXISTS users
  (
     `id`        INT UNSIGNED NOT NULL auto_increment,
     `firstname` VARCHAR(255) NOT NULL,
     `lastname`  VARCHAR(255) NOT NULL,
     `email`     VARCHAR(255) NOT NULL,
     `password`  VARCHAR(255) NOT NULL,
     `createdat` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
     PRIMARY KEY (`id`),
     UNIQUE KEY `email_unique` (`email`)
  );
