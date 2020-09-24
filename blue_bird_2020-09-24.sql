# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.5.5-10.4.13-MariaDB)
# Database: blue_bird
# Generation Time: 2020-09-24 09:46:55 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table car
# ------------------------------------------------------------

DROP TABLE IF EXISTS `car`;

CREATE TABLE `car` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `brand` varchar(255) DEFAULT NULL,
  `price` int(11) DEFAULT NULL,
  `kondisi` varchar(255) DEFAULT NULL,
  `quantity` int(11) DEFAULT NULL,
  `description` text DEFAULT NULL,
  `specification` text DEFAULT NULL,
  `image` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `car` WRITE;
/*!40000 ALTER TABLE `car` DISABLE KEYS */;

INSERT INTO `car` (`id`, `name`, `brand`, `price`, `kondisi`, `quantity`, `description`, `specification`, `image`, `created_at`, `updated_at`)
VALUES
	(26,'Avanza','Toyota',255000000,'New',1,'Harga Toyota Rush terbaru September 2020 mulai dari Rp 255.2 Juta. Sebelum beli, dapatkan info spesifikasi, konsumsi BBM, promo diskon dan simulasi ...','Dimensi: 4.435 mm P x 1.695 mm L x 1.705 mm T','https://tsoimageprod.azureedge.net/sys-master-hybrismedia/h47/h46/8807946092574/Rush_1_Large.jpg','2020-09-24 16:13:20','2020-09-24 16:13:20'),
	(27,'Rush','Toyota',255000000,'New',1,'Harga Toyota Rush terbaru September 2020 mulai dari Rp 255.2 Juta. Sebelum beli, dapatkan info spesifikasi, konsumsi BBM, promo diskon dan simulasi ...','Dimensi: 4.435 mm P x 1.695 mm L x 1.705 mm T','https://tsoimageprod.azureedge.net/sys-master-hybrismedia/h47/h46/8807946092574/Rush_1_Large.jpg','2020-09-24 16:19:42','2020-09-24 16:19:42'),
	(28,'Calya','Toyota',255000000,'New',1,'Harga Toyota Rush terbaru September 2020 mulai dari Rp 255.2 Juta. Sebelum beli, dapatkan info spesifikasi, konsumsi BBM, promo diskon dan simulasi ...','Dimensi: 4.435 mm P x 1.695 mm L x 1.705 mm T','https://tsoimageprod.azureedge.net/sys-master-hybrismedia/h47/h46/8807946092574/Rush_1_Large.jpg','2020-09-24 16:30:48','2020-09-24 16:30:48');

/*!40000 ALTER TABLE `car` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
