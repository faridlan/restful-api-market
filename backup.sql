-- MySQL dump 10.13  Distrib 8.0.28, for Linux (x86_64)
--
-- Host: localhost    Database: e_market
-- ------------------------------------------------------
-- Server version	8.0.28

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `addresses`
--

DROP TABLE IF EXISTS `addresses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `addresses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `name` varchar(100) NOT NULL,
  `handphone_number` varchar(15) NOT NULL,
  `street` varchar(100) NOT NULL,
  `districk` varchar(100) NOT NULL,
  `post_code` int NOT NULL,
  `comment` text,
  PRIMARY KEY (`id`),
  KEY `fk_addresses_user` (`user_id`),
  CONSTRAINT `fk_addresses_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `addresses`
--

LOCK TABLES `addresses` WRITE;
/*!40000 ALTER TABLE `addresses` DISABLE KEYS */;
INSERT INTO `addresses` VALUES (1,1,'Faridlan nul hakim','089746526342','Jl. Asgard No 60','cipedes',46717,'RT 09 RW 99'),(2,2,'udin','08567641234','Jl.Leuwidahu','Indihiang',46152,'Gang Melati Belakang'),(5,6,'jonathan','0897635421','Jl.Asgard','indihiang',55632,'samping patung THOR');
/*!40000 ALTER TABLE `addresses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `blacklist`
--

DROP TABLE IF EXISTS `blacklist`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `blacklist` (
  `id` int NOT NULL AUTO_INCREMENT,
  `token` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blacklist`
--

LOCK TABLES `blacklist` WRITE;
/*!40000 ALTER TABLE `blacklist` DISABLE KEYS */;
INSERT INTO `blacklist` VALUES (1,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoiWFZsQnpnYmFpQ01SQWpXd2hUSGMiLCJleHAiOjE2NTA5MjE0MTd9.uUoZqo_cqw3XWGakBUjX_gv96MZt6Yt_besEFQC79cY'),(2,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJmYXJpZGxhbiIsImVtYWlsIjoiZmFyaWRsYW5AZ21haWwuY29tIiwicm9sZV9pZCI6MiwidG9rZW4iOiJYVmxCemdiYWlDTVJBald3aFRIYyIsImV4cCI6MTY1MTU5Njk4MX0.5J7-uKZdpo42DanjwtZ-UOnbgEYCgrm6gfoV-nVXtkE'),(3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidXNlcm5hbWUiOiJqb25hdGhhbiBudWwiLCJyb2xlX2lkIjoyLCJ0b2tlbiI6IlhWbEJ6Z2JhaUNNUkFqV3doVEhjIiwiZXhwIjoxNjUyNjQwNDUwfQ.anGrFiLMmkdXoiqFUgYwIHCAs5HdC1DnG_NWXvNNydg'),(4,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoidGN1QXhoeEtRRkRhRnBMU2pGYmMiLCJleHAiOjE2NTI2NDA0NTB9.8n-RR4D8qH2x7DBDPrZmf-gAo4FOxN_VLkuf2HdIMl8'),(5,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoiWFZsQnpnYmFpQ01SQWpXd2hUSGMiLCJleHAiOjE2NTI3MjA4NTJ9.-iZe5OnYRrXjkdtCVX29IxcGuCvBgcUPKWaksEjYvek');
/*!40000 ALTER TABLE `blacklist` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `carts`
--

DROP TABLE IF EXISTS `carts`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `carts` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `product_id` int NOT NULL,
  `quantity` int NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_carts_user` (`user_id`),
  KEY `fk_carts_product` (`product_id`),
  CONSTRAINT `fk_carts_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  CONSTRAINT `fk_carts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `carts`
--

LOCK TABLES `carts` WRITE;
/*!40000 ALTER TABLE `carts` DISABLE KEYS */;
INSERT INTO `carts` VALUES (20,2,1,2),(21,2,2,2),(23,2,5,0),(24,2,5,0);
/*!40000 ALTER TABLE `carts` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `categories`
--

DROP TABLE IF EXISTS `categories`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `categories` (
  `id` int NOT NULL AUTO_INCREMENT,
  `category_name` varchar(150) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'Gadget'),(2,'Sport'),(3,'Vape'),(4,'Bag');
/*!40000 ALTER TABLE `categories` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders`
--

DROP TABLE IF EXISTS `orders`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `address_id` int NOT NULL,
  `total` int NOT NULL DEFAULT '0',
  `order_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status_id` int NOT NULL DEFAULT '1',
  `payment_id` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `fk_payment_id_unique` (`payment_id`),
  KEY `fk_orders_user` (`user_id`),
  KEY `fk_status_order` (`status_id`),
  KEY `fk_orders_address` (`address_id`),
  CONSTRAINT `fk_orders_address` FOREIGN KEY (`address_id`) REFERENCES `addresses` (`id`),
  CONSTRAINT `fk_orders_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_payment_order` FOREIGN KEY (`payment_id`) REFERENCES `payments` (`id`),
  CONSTRAINT `fk_status_order` FOREIGN KEY (`status_id`) REFERENCES `status_order` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (1,2,2,10000000,'2022-05-05 15:52:36',1,1),(2,2,2,19000000,'2022-05-05 16:08:02',1,NULL),(3,2,2,9000000,'2022-05-05 16:42:04',1,NULL),(4,2,2,9000000,'2022-05-05 17:51:31',1,NULL),(8,2,2,4000000,'2022-05-08 04:19:23',1,NULL),(19,2,2,0,'2022-05-09 13:07:21',1,NULL),(24,2,2,0,'2022-05-09 13:26:49',1,NULL),(25,2,2,0,'2022-05-09 13:39:49',1,NULL),(27,2,2,11100000,'2022-05-09 14:51:46',1,NULL),(28,2,2,11100000,'2022-05-09 15:35:11',1,NULL),(29,2,2,4000000,'2022-05-10 04:33:11',1,NULL),(30,2,2,1050000,'2022-05-10 04:34:45',1,NULL),(38,2,2,4000000,'2022-05-15 18:13:26',1,NULL),(40,6,5,2000000,'2022-05-15 18:46:27',1,NULL),(42,2,2,10000000,'2022-05-16 06:40:18',1,NULL);
/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `orders_detail`
--

DROP TABLE IF EXISTS `orders_detail`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders_detail` (
  `order_id` int NOT NULL,
  `product_id` int NOT NULL,
  `quantity` int NOT NULL,
  `unit_price` int DEFAULT NULL,
  `total_price` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`order_id`,`product_id`),
  KEY `fk_orders_detail_product` (`product_id`),
  CONSTRAINT `fk_orders_detail_order` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
  CONSTRAINT `fk_orders_detail_product` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders_detail`
--

LOCK TABLES `orders_detail` WRITE;
/*!40000 ALTER TABLE `orders_detail` DISABLE KEYS */;
INSERT INTO `orders_detail` VALUES (1,1,2,5000000,10000000),(2,1,2,5000000,10000000),(2,2,3,2000000,6000000),(2,3,1,3000000,3000000),(3,1,1,5000000,5000000),(3,2,2,2000000,4000000),(4,3,2,3000000,6000000),(4,4,3,1000000,3000000),(8,2,2,2000000,4000000),(19,1,2,5000000,0),(19,2,2,2000000,0),(19,3,2,3000000,0),(24,2,2,2000000,0),(24,3,2,3000000,0),(24,4,2,1000000,0),(25,1,2,5000000,0),(25,5,2,50000,0),(25,6,2,500000,0),(27,1,2,5000000,10000000),(27,5,2,50000,100000),(27,6,2,500000,1000000),(28,1,2,5000000,10000000),(28,5,2,50000,100000),(28,6,2,500000,1000000),(29,2,2,2000000,4000000),(30,4,1,1000000,1000000),(30,5,1,50000,50000),(38,2,2,2000000,4000000),(40,4,2,1000000,2000000),(42,1,2,5000000,10000000);
/*!40000 ALTER TABLE `orders_detail` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payments`
--

DROP TABLE IF EXISTS `payments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `payments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `image_url` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `payments`
--

LOCK TABLES `payments` WRITE;
/*!40000 ALTER TABLE `payments` DISABLE KEYS */;
INSERT INTO `payments` VALUES (1,'image.com/zzz');
/*!40000 ALTER TABLE `payments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `products`
--

DROP TABLE IF EXISTS `products`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_name` varchar(200) NOT NULL,
  `category_id` int NOT NULL,
  `price` int NOT NULL,
  `quantity` int NOT NULL,
  `image_url` text NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_products_category` (`category_id`),
  CONSTRAINT `fk_products_category` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'Laptop',1,5000000,3,'http://image.com/default'),(2,'Handphone',1,2000000,12,'http://image.com/default'),(3,'Tab',1,3000000,13,'http://image.com/default'),(4,'Sport Shoes',2,1000000,4,'http://image.com/default'),(5,'Cotton Bacon',3,50000,15,'http://image.com/default'),(6,'Vape Thelema',3,500000,26,'http://image.com/default'),(7,'Liquid 100 ml',3,100000,50,'http://image.com/default');
/*!40000 ALTER TABLE `products` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `roles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `role_name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'admin'),(2,'user');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `status_order`
--

DROP TABLE IF EXISTS `status_order`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `status_order` (
  `id` int NOT NULL AUTO_INCREMENT,
  `status_name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `status_order`
--

LOCK TABLES `status_order` WRITE;
/*!40000 ALTER TABLE `status_order` DISABLE KEYS */;
INSERT INTO `status_order` VALUES (1,'Menunggu Pembayaran');
/*!40000 ALTER TABLE `status_order` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(100) NOT NULL,
  `email` varchar(200) NOT NULL,
  `password` varchar(225) NOT NULL,
  `image_url` varchar(225) NOT NULL DEFAULT 'image.com/defaultProfile.jpg',
  `role_id` int NOT NULL DEFAULT '2',
  PRIMARY KEY (`id`),
  KEY `fk_user_role` (`role_id`),
  CONSTRAINT `fk_user_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'faridlan','faridlan@gmail.com','rahasia123','image.com/defaultProfile.jpg',2),(2,'udin','udin@gmail.com','rahasia123','image.com/defaultProfile.jpg',2),(4,'poi','poi@gmail.com','rahasia123','image.com/defaultProfile.jpg',2),(5,'poilahhh','poii@gmail.com','rahasia123','',2),(6,'jonathan nul','','rahasia123','',2);
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-05-17  5:05:23
