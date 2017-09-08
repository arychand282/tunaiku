-- MySQL dump 10.16  Distrib 10.1.26-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: tunaiku
-- ------------------------------------------------------
-- Server version	10.1.26-MariaDB-1~xenial

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `data`
--

DROP TABLE IF EXISTS `data`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `data` (
  `datedata` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `open` int(11) DEFAULT NULL,
  `high` int(11) DEFAULT NULL,
  `low` int(11) DEFAULT NULL,
  `close` int(11) DEFAULT NULL,
  `adjclose` int(11) DEFAULT NULL,
  `volume` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `data`
--

LOCK TABLES `data` WRITE;
/*!40000 ALTER TABLE `data` DISABLE KEYS */;
INSERT INTO `data` VALUES ('2017-06-16',336,346,328,338,338,480537400),('2017-06-19',342,362,338,356,356,652163600),('2017-06-20',360,364,356,356,356,340340400),('2017-06-21',356,358,346,352,352,190819700),('2017-06-22',348,356,336,340,340,234361000),('2017-06-23',340,340,340,340,340,0),('2017-06-26',340,340,340,340,340,0),('2017-06-27',340,340,340,340,340,0),('2017-06-28',340,340,340,340,340,0),('2017-06-29',340,340,340,340,340,0),('2017-06-30',340,340,340,340,340,0),('2017-07-03',360,394,360,374,374,692977900),('2017-07-04',378,380,356,360,360,392925200),('2017-07-05',362,364,344,344,344,376202400),('2017-07-06',346,352,340,342,342,452544700),('2017-07-07',344,346,328,332,332,370934100),('2017-07-10',340,340,322,328,328,153485700),('2017-07-11',328,336,320,324,324,190702100),('2017-07-12',326,332,324,332,332,147953800),('2017-07-13',336,336,328,328,328,107949700),('2017-07-14',330,346,328,346,346,236236300),('2017-07-17',350,354,334,350,350,384016200),('2017-07-18',350,352,342,344,344,189522100),('2017-07-19',348,358,346,350,350,341492600),('2017-07-20',350,356,344,344,344,221387800),('2017-07-21',346,348,334,334,334,163228500),('2017-07-24',336,336,330,332,332,96372600),('2017-07-25',332,340,330,332,332,133683200),('2017-07-26',334,338,326,328,328,115154900),('2017-07-27',326,334,322,328,328,122207200),('2017-07-28',332,360,332,358,358,620423200),('2017-07-31',360,362,342,342,342,320322700),('2017-08-01',350,354,338,340,340,258064500),('2017-08-02',338,346,336,338,338,164324100),('2017-08-03',342,344,332,334,334,107018100),('2017-08-04',334,340,328,330,330,201458900),('2017-08-07',330,334,322,324,324,223815700),('2017-08-08',326,328,300,306,306,644036000),('2017-08-09',308,308,286,296,296,628935000),('2017-08-10',296,302,280,282,282,525483900),('2017-08-11',280,288,270,276,276,423930100),('2017-08-14',280,286,276,280,280,245859500),('2017-08-15',282,296,280,294,294,360328800),('2017-08-16',296,302,290,292,292,287693500);
/*!40000 ALTER TABLE `data` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `even_odd_number`
--

DROP TABLE IF EXISTS `even_odd_number`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `even_odd_number` (
  `id` mediumint(9) NOT NULL AUTO_INCREMENT,
  `number_evens` int(11) DEFAULT NULL,
  `number_odds` int(11) DEFAULT NULL,
  `total` int(11) DEFAULT NULL,
  `name_total` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `even_odd_number`
--

LOCK TABLES `even_odd_number` WRITE;
/*!40000 ALTER TABLE `even_odd_number` DISABLE KEYS */;
/*!40000 ALTER TABLE `even_odd_number` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `nominal_amount`
--

DROP TABLE IF EXISTS `nominal_amount`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `nominal_amount` (
  `id` mediumint(9) NOT NULL AUTO_INCREMENT,
  `nominal` double DEFAULT NULL,
  `additional_number` int(11) DEFAULT NULL,
  `total` double DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `nominal_amount`
--

LOCK TABLES `nominal_amount` WRITE;
/*!40000 ALTER TABLE `nominal_amount` DISABLE KEYS */;
/*!40000 ALTER TABLE `nominal_amount` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `prime_number`
--

DROP TABLE IF EXISTS `prime_number`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `prime_number` (
  `id` mediumint(9) NOT NULL AUTO_INCREMENT,
  `numbers` int(11) DEFAULT NULL,
  `name_number` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `prime_number`
--

LOCK TABLES `prime_number` WRITE;
/*!40000 ALTER TABLE `prime_number` DISABLE KEYS */;
/*!40000 ALTER TABLE `prime_number` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-09-09  1:12:21
