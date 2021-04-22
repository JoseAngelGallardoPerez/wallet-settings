-- MySQL dump 10.13  Distrib 5.7.24, for osx10.14 (x86_64)
--
-- Host: localhost    Database: settings
-- ------------------------------------------------------
-- Server version	5.7.24

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
-- Table structure for table `configs`
--

DROP TABLE IF EXISTS `configs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `configs` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'The config id',
  `path` varchar(255) NOT NULL COMMENT 'The path is constructed as: section/group/field.',
  `value` text NOT NULL COMMENT 'The value.',
  `scope` enum('private','public') NOT NULL DEFAULT 'private',
  `root_only` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `path_UNIQUE` (`path`)
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8 COMMENT='Config Data';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `configs`
--

LOCK TABLES `configs` WRITE;
/*!40000 ALTER TABLE `configs` DISABLE KEYS */;
INSERT INTO `configs` VALUES (1,'regional/general/default_timezone','Europe/Amsterdam','public',0),(2,'regional/general/default_date_format','DD/MM/YYYY','public',0),(3,'regional/general/site_name','VELMIE_WALLET','public',0),(4,'regional/general/user_register','enable','public',0),(5,'regional/modules/velmie_wallet_cards','enable','private',0),(6,'regional/modules/velmie_wallet_gdpr','disable','public',0),(7,'regional/login/failed_login_username_use','yes','private',0),(8,'regional/login/failed_login_username_limit','10','private',0),(9,'regional/login/failed_login_username_cleanup','10','private',0),(10,'regional/login/failed_login_user_use','6','private',0),(11,'regional/login/failed_login_user_window','1','private',0),(12,'profile/user-options/field_user_beneficial_owner','enable','private',0),(13,'profile/user-options/dormant','3','private',0),(14,'profile/user-options/field_user_first_name','yes','private',0),(15,'profile/user-options/field_user_last_name','yes','private',0),(16,'profile/user-options/field_user_company_name','yes','private',0),(17,'profile/user-options/field_user_date_of_birth','yes','private',0),(18,'profile/user-options/field_user_document_personal_id','yes','private',0),(19,'profile/user-options/field_user_country_of_residence_iso2','yes','private',0),(20,'profile/user-options/field_user_country_of_citizenship_iso2','yes','private',0),(21,'profile/user-options/field_user_office_phone_number','yes','private',0),(22,'profile/user-options/field_user_home_phone_number','yes','private',0),(23,'profile/user-options/field_user_phone_number','yes','private',0),(24,'profile/user-options/field_user_fax','yes','private',0),(25,'profile/user-options/field_user_pa_address','yes','private',0),(26,'profile/user-options/field_user_pa_address_2nd_line','yes','private',0),(27,'profile/user-options/field_user_pa_city','yes','private',0),(28,'profile/user-options/field_user_pa_country_iso2','yes','private',0),(29,'profile/user-options/field_user_pa_state_prov_region','yes','private',0),(30,'profile/user-options/field_user_pa_zip_postal_code','yes','private',0),(31,'profile/user-options/field_user_ma_as_physical','yes','private',0),(32,'profile/user-options/field_user_ma_name','yes','private',0),(33,'profile/user-options/field_user_ma_address','yes','private',0),(34,'profile/user-options/field_user_ma_city','yes','private',0),(35,'profile/user-options/field_user_ma_address_2nd_line','yes','private',0),(36,'profile/user-options/field_user_ma_state_prov_region','yes','private',0),(37,'profile/user-options/field_user_ma_zip_postal_code','yes','private',0),(38,'profile/user-options/field_user_ma_phone_number','yes','private',0),(39,'profile/user-options/field_user_ma_country_iso2','yes','private',0),(40,'profile/user-options/field_user_bo_full_name','yes','private',0),(41,'profile/user-options/field_user_bo_date_of_birth','yes','private',0),(42,'profile/user-options/field_user_bo_document_type','yes','private',0),(43,'profile/user-options/field_user_bo_document_personal_id','yes','private',0),(44,'profile/user-options/field_user_bo_relationship','yes','private',0),(45,'profile/user-options/field_user_bo_address','yes','private',0),(46,'profile/user-options/field_user_bo_phone_number','yes','private',0),(47,'profile/autologout/status','disable','public',0),(48,'profile/autologout/timeout','10','public',0),(49,'profile/autologout/message','Your session is about to expire.','public',0),(50,'profile/autologout/inactivity_message','You have been logged out due to inactivity.','public',0),(51,'profile/autologout/padding','30','public',0),(52,'regional/general/site_url','http://velmie-wallet.de','public',0),(53,'regional/general/site_incoming_message_path','/messages/incoming/{id}','private',0),(54,'regional/general/site_my_profile_settings_path','/my-profile/settings','private',0),(55,'profile/user-options/field_user_email','yes','private',0),(56,'regional/general/maintenance','disabled','public',1),(57,'regional/general/maintenance_text','This system is temporarily unavailable.','public',1),(58,'profile/user-options/field_sms_phone_number','yes','private',0);
/*!40000 ALTER TABLE `configs` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `schema_migrations`
--

DROP TABLE IF EXISTS `schema_migrations`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `schema_migrations` (
  `version` bigint(20) NOT NULL,
  `dirty` tinyint(1) NOT NULL,
  PRIMARY KEY (`version`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `schema_migrations`
--

LOCK TABLES `schema_migrations` WRITE;
/*!40000 ALTER TABLE `schema_migrations` DISABLE KEYS */;
INSERT INTO `schema_migrations` VALUES (20190430115713,0);
/*!40000 ALTER TABLE `schema_migrations` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-05-03 13:22:53
