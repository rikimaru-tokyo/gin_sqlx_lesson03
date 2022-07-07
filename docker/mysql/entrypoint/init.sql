DROP TABLE IF EXISTS targets;


CREATE TABLE IF NOT EXISTS `targets` (
    `id` tinyint AUTO_INCREMENT NOT NULL,
    `name` varchar(50) NOT NULL,
    `birthday` DATE NOT NULL,
    PRIMARY KEY(`id`)
);


SET @SAMPLE_TIMESTAMP = "2022-06-07 12:34:56";
INSERT INTO targets VALUES (1,'alpha','2022-01-01');
INSERT INTO targets VALUES (2,'bravo','2022-01-22');
INSERT INTO targets VALUES (3,'charlie','2022-02-22');


