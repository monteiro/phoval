CREATE TABLE `verification` (
    `id` VARCHAR(36) PRIMARY KEY,
    `country_code` VARCHAR(100) NOT NULL,
    `phone_number` VARCHAR(100) NOT NULL,
    `code` int NOT NULL,
    `created_at` DATETIME NULL,
    `verified_at` DATETIME NULL
);
