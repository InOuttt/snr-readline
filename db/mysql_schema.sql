#
# SQL Export
# Created by Querious (300050)
# Created: 7 September 2020 21.33.41 GMT+8
# Encoding: Unicode (UTF-8)
#


SET @ORIG_FOREIGN_KEY_CHECKS = @@FOREIGN_KEY_CHECKS;
SET FOREIGN_KEY_CHECKS = 0;

SET @ORIG_UNIQUE_CHECKS = @@UNIQUE_CHECKS;
SET UNIQUE_CHECKS = 0;

SET @ORIG_TIME_ZONE = @@TIME_ZONE;
SET TIME_ZONE = '+00:00';

SET @ORIG_SQL_MODE = @@SQL_MODE;
SET SQL_MODE = 'NO_AUTO_VALUE_ON_ZERO';



CREATE TABLE `twitter_feed` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `feed_id` varchar(200) CHARACTER SET utf8 NOT NULL,
  `published_date` date NOT NULL,
  `link` varchar(255) CHARACTER SET utf8 NOT NULL,
  `content` text COLLATE utf8mb4_unicode_520_ci NOT NULL,
  `author_name` varchar(140) COLLATE utf8mb4_unicode_520_ci NOT NULL DEFAULT '',
  `author_link` varchar(255) CHARACTER SET utf8 NOT NULL,
  `author_avatar` varchar(255) CHARACTER SET utf8 NOT NULL,
  `author_username` varchar(255) CHARACTER SET utf8 NOT NULL,
  `location` varchar(100) CHARACTER SET utf8 DEFAULT NULL,
  `coord_lat` float DEFAULT NULL,
  `coord_lon` float DEFAULT NULL,
  `following` int(11) NOT NULL DEFAULT '0',
  `followers` int(11) NOT NULL DEFAULT '0',
  `language` varchar(5) CHARACTER SET utf8 NOT NULL,
  `num_replies` int(11) DEFAULT '0',
  `num_rts` int(11) DEFAULT '0',
  `published_ts` int(11) NOT NULL DEFAULT '0',
  `status_type` int(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `feed_id` (`feed_id`),
  KEY `author_id` (`author_username`),
  KEY `published_date` (`published_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;




LOCK TABLES `twitter_feed` WRITE;
UNLOCK TABLES;






SET FOREIGN_KEY_CHECKS = @ORIG_FOREIGN_KEY_CHECKS;

SET UNIQUE_CHECKS = @ORIG_UNIQUE_CHECKS;

SET @ORIG_TIME_ZONE = @@TIME_ZONE;
SET TIME_ZONE = @ORIG_TIME_ZONE;

SET SQL_MODE = @ORIG_SQL_MODE;



# Export Finished: 7 September 2020 21.33.41 GMT+8
