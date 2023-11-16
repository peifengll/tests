CREATE TABLE `score` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `stu_id` int NOT NULL,
                         `c_name` varchar(20) DEFAULT NULL,
                         `grade` int DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb3;

