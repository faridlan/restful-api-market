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
  `id_address` varchar(36) DEFAULT NULL,
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
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `addresses`
--

LOCK TABLES `addresses` WRITE;
/*!40000 ALTER TABLE `addresses` DISABLE KEYS */;
INSERT INTO `addresses` VALUES (1,'36159d6ee3cb11ec989c0242ac110002',1,'Faridlan nul hakim','089746526342','Jl. Asgard No 60','cipedes',46717,'RT 09 RW 99'),(2,'37cace3de3cb11ec989c0242ac110002',2,'udin','08567641234','Jl.Leuwidahu','Indihiang',46152,'Gang Melati Belakang'),(9,'38d107a1e3cb11ec989c0242ac110002',2,'Udin Jhon','08972635212','Jl.Mitra','Cipedes',45326,'Depan ATM');
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
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `blacklist`
--

LOCK TABLES `blacklist` WRITE;
/*!40000 ALTER TABLE `blacklist` DISABLE KEYS */;
INSERT INTO `blacklist` VALUES (1,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoiWFZsQnpnYmFpQ01SQWpXd2hUSGMiLCJleHAiOjE2NTA5MjE0MTd9.uUoZqo_cqw3XWGakBUjX_gv96MZt6Yt_besEFQC79cY'),(2,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJmYXJpZGxhbiIsImVtYWlsIjoiZmFyaWRsYW5AZ21haWwuY29tIiwicm9sZV9pZCI6MiwidG9rZW4iOiJYVmxCemdiYWlDTVJBald3aFRIYyIsImV4cCI6MTY1MTU5Njk4MX0.5J7-uKZdpo42DanjwtZ-UOnbgEYCgrm6gfoV-nVXtkE'),(3,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidXNlcm5hbWUiOiJqb25hdGhhbiBudWwiLCJyb2xlX2lkIjoyLCJ0b2tlbiI6IlhWbEJ6Z2JhaUNNUkFqV3doVEhjIiwiZXhwIjoxNjUyNjQwNDUwfQ.anGrFiLMmkdXoiqFUgYwIHCAs5HdC1DnG_NWXvNNydg'),(4,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoidGN1QXhoeEtRRkRhRnBMU2pGYmMiLCJleHAiOjE2NTI2NDA0NTB9.8n-RR4D8qH2x7DBDPrZmf-gAo4FOxN_VLkuf2HdIMl8'),(5,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoiWFZsQnpnYmFpQ01SQWpXd2hUSGMiLCJleHAiOjE2NTI3MjA4NTJ9.-iZe5OnYRrXjkdtCVX29IxcGuCvBgcUPKWaksEjYvek'),(6,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoiWFZsQnpnYmFpQ01SQWpXd2hUSGMiLCJleHAiOjE2NTMyNDA5MzB9.qVgmZ0JVSwgGQI7jgJ1jyUq1LgXGBVQqkvUBq3fqavE'),(7,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NywidXNlcm5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIiwicm9sZV9pZCI6MSwidG9rZW4iOiJYVmxCemdiYWlDTVJBald3aFRIYyIsImV4cCI6MTY1MzI0MTEwNn0.Sx1dgSwTwfVpVFN-GlVdjWjzMLR3Yf2s0u8GcXig4NI'),(8,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidXNlcm5hbWUiOiJqb25hdGhhbiBudWwiLCJyb2xlX2lkIjoyLCJ0b2tlbiI6IlhWbEJ6Z2JhaUNNUkFqV3doVEhjIiwiZXhwIjoxNjU0MDE1NTE2fQ.rcwBgEYahO3-FcMIuvWtd_Gmds1ogANqimGE_17x3EU'),(9,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NywidXNlcm5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIiwicm9sZV9pZCI6MSwidG9rZW4iOiJYVmxCemdiYWlDTVJBald3aFRIYyIsImV4cCI6MTY1NDAxOTg0MX0.APmyh2-E3tazw60GC1gI88yHWRYQJMdEEZwDWB3AXFY'),(10,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidXNlcm5hbWUiOiJqb25hdGhhbiBudWwiLCJyb2xlX2lkIjoyLCJ0b2tlbiI6IlhWbEJ6Z2JhaUNNUkFqV3doVEhjIiwiZXhwIjoxNjU0MDIyNDk1fQ.0m8ZJ8qhxfk6hktIXkB_5RlfaV2AIzCLRfLkmaqmt7Q'),(11,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NiwidXNlcm5hbWUiOiJqb25hdGhhbiBudWwiLCJyb2xlX2lkIjoyLCJ0b2tlbiI6IlhWbEJ6Z2JhaUNNUkFqV3doVEhjIiwiZXhwIjoxNjU0MDIyNjE0fQ.n0l9k2S8yoelQPHM4m5drROVUBePyT4LuCqhEIMr8W4'),(12,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoidGN1QXhoeEtRRkRhRnBMU2pGYmMiLCJleHAiOjE2NTQwMjI2MTR9.jkm7aLN6EAIGm_zmzwt6qEx55dN3O_bnjrkkVjpC8WU'),(13,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NywidXNlcm5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIiwicm9sZV9pZCI6MSwidG9rZW4iOiJYVmxCemdiYWlDTVJBald3aFRIYyIsImV4cCI6MTY1NDA4NTIyM30.cnZl1SqLBsH_ECiIPcDOuYi7DHCLWVijUjBbixEbu_0'),(14,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoiWFZsQnpnYmFpQ01SQWpXd2hUSGMiLCJleHAiOjE2NTQwODU5NTJ9.mZaXOkXyk_S4ICkZnk5mGubB4p-qLoJxRDGcqCdj-yI'),(15,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NywidXNlcm5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIiwicm9sZV9pZCI6MSwidG9rZW4iOiJYVmxCemdiYWlDTVJBald3aFRIYyIsImV4cCI6MTY1NDA4OTMxM30.MkN5CJYBGGMdNEbqUurWjX0-0jo8tM6SheJtFzWezyc'),(16,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoiWFZsQnpnYmFpQ01SQWpXd2hUSGMiLCJleHAiOjE2NTQxODYzNTZ9.Ah4zyozjPwOLRzjayPqAkacTFyQbA5h6xoHsWWyAy3Y'),(17,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NywidXNlcm5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIiwicm9sZV9pZCI6MSwidG9rZW4iOiJYVmxCemdiYWlDTVJBald3aFRIYyIsImV4cCI6MTY1NDE4Nzk2M30.KqogrkVGsRABsxI1kP4DfV-JlhT7pMMOWJrhiBEtBjc'),(18,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiaWRfdXNlciI6IjE0YTE4NGE0ZTNjYjExZWM5ODljMDI0MmFjMTEwMDAyIiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoiWFZsQnpnYmFpQ01SQWpXd2hUSGMiLCJleHAiOjE2NTQzNTY5NjJ9.wMq-UMzMFEzzBFmK06zj5CXwEIZDb0xmoybs-90A2-k'),(19,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NywiaWRfdXNlciI6IjE5Mzk0ZjU3ZTNjYjExZWM5ODljMDI0MmFjMTEwMDAyIiwidXNlcm5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIiwicm9sZV9pZCI6MSwidG9rZW4iOiJYVmxCemdiYWlDTVJBald3aFRIYyIsImV4cCI6MTY1NDM1ODM3OX0.sMfyRndddENtk11AEuaZku6VzBXy91cTJrhUoQ1jBsw'),(20,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiaWRfdXNlciI6IjE0YTE4NGE0ZTNjYjExZWM5ODljMDI0MmFjMTEwMDAyIiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoiWFZsQnpnYmFpQ01SQWpXd2hUSGMiLCJleHAiOjE2NTQ2MTc1NDl9.Q1MS0GUmnc7Q5Gu96vIG1QiS4k1XogYlnzsbscRzDCY'),(21,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NywiaWRfdXNlciI6IjE5Mzk0ZjU3ZTNjYjExZWM5ODljMDI0MmFjMTEwMDAyIiwidXNlcm5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIiwicm9sZV9pZCI6MSwidG9rZW4iOiJYVmxCemdiYWlDTVJBald3aFRIYyIsImV4cCI6MTY1NDYxODk2OH0.r9GLWbcztWvmc8byDMImFW2vLBkxpYedV6V8Yi0WSyE'),(22,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MTcsImlkX3VzZXIiOiI3Yzc2YzZmY2U1YmUxMWVjYWM3MTAyNDJhYzEzMDAwMiIsInVzZXJuYW1lIjoiTWFpbWFpMTIzNDUiLCJlbWFpbCI6Im1haW1haUBnbWFpbC5jb20iLCJyb2xlX2lkIjoyLCJ0b2tlbiI6InRjdUF4aHhLUUZEYUZwTFNqRmJjIiwiZXhwIjoxNjU0NjIzMDIzfQ.kvos24SOAtrBlMtv8GlsEoowrB-1ERDYgZFUsh2Fe4k'),(23,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NywiaWRfdXNlciI6IjE5Mzk0ZjU3ZTNjYjExZWM5ODljMDI0MmFjMTEwMDAyIiwidXNlcm5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIiwicm9sZV9pZCI6MSwidG9rZW4iOiJYVmxCemdiYWlDTVJBald3aFRIYyIsImV4cCI6MTY1NDY5NDU3N30.2E-BPMVRBKImnxGa1Zy2JeIBEFfxACUD8-r5xxKBb2k'),(24,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MiwiaWRfdXNlciI6IjE0YTE4NGE0ZTNjYjExZWM5ODljMDI0MmFjMTEwMDAyIiwidXNlcm5hbWUiOiJ1ZGluIiwiZW1haWwiOiJ1ZGluQGdtYWlsLmNvbSIsInJvbGVfaWQiOjIsInRva2VuIjoiWFZsQnpnYmFpQ01SQWpXd2hUSGMiLCJleHAiOjE2NTQ2OTQ5MDd9.dqDj4a9rb-KxdiqWdrBVo0FGHJ3U8ElqTJKMtHcufpc'),(25,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NywiaWRfdXNlciI6IjE5Mzk0ZjU3ZTNjYjExZWM5ODljMDI0MmFjMTEwMDAyIiwidXNlcm5hbWUiOiJhZG1pbiIsImVtYWlsIjoiYWRtaW5AYWRtaW4uY29tIiwicm9sZV9pZCI6MSwidG9rZW4iOiJYVmxCemdiYWlDTVJBald3aFRIYyIsImV4cCI6MTY1NDY5NzE4MH0.QltVr3P0dTbAcz1q0GilZryqQ9dBRbakNCrjcCcQ07Y'),(26,'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZF91c2VyIjoiMTkzOTRmNTdlM2NiMTFlYzk4OWMwMjQyYWMxMTAwMDIiLCJ1c2VybmFtZSI6ImFkbWluIiwiZW1haWwiOiJhZG1pbkBhZG1pbi5jb20iLCJyb2xlX2lkIjoxLCJ0b2tlbiI6IkpqUGp6cGZSRkVnbW90YUZldEhzIiwiZXhwIjoxNjU0Njk4NjMxfQ.55KdKc24nGEj-FWJp1wvIbF2GdSMrNjAWcAT_RAlJEs');
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
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `carts`
--

LOCK TABLES `carts` WRITE;
/*!40000 ALTER TABLE `carts` DISABLE KEYS */;
INSERT INTO `carts` VALUES (21,2,2,2),(26,2,4,2),(28,2,6,2);
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
  `id_category` varchar(36) DEFAULT NULL,
  `category_name` varchar(150) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `categories`
--

LOCK TABLES `categories` WRITE;
/*!40000 ALTER TABLE `categories` DISABLE KEYS */;
INSERT INTO `categories` VALUES (1,'54111de7e26711ec93820242ac110002','Gadget'),(2,'552c3350e26711ec93820242ac110002','Sport'),(3,'561684f9e26711ec93820242ac110002','Vape'),(4,'56e194e2e26711ec93820242ac110002','Bag'),(6,'7058dda8e27611ec93820242ac110002','Shoes'),(7,'e3f156dde27611ec93820242ac110002','Watch'),(11,'dbc4833be41b11ec989c0242ac110002','Rope');
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
  `id_order` varchar(36) DEFAULT NULL,
  `user_id` int NOT NULL,
  `address_id` int NOT NULL,
  `total` int NOT NULL DEFAULT '0',
  `order_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `status_id` int NOT NULL DEFAULT '1',
  `payment` varchar(150) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_orders_user` (`user_id`),
  KEY `fk_status_order` (`status_id`),
  KEY `fk_orders_address` (`address_id`),
  CONSTRAINT `fk_orders_address` FOREIGN KEY (`address_id`) REFERENCES `addresses` (`id`),
  CONSTRAINT `fk_orders_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `fk_status_order` FOREIGN KEY (`status_id`) REFERENCES `status_order` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;
INSERT INTO `orders` VALUES (54,'78b356dbe28711ec93820242ac110002',2,2,1050000,'2022-06-02 15:20:01',2,'https://olshop.sgp1.digitaloceanspaces.com/payments/tcuAxhxKQF.png'),(55,'7542f029e28a11ec93820242ac110002',7,2,1050000,'2022-06-02 15:41:24',1,NULL),(88,'ed3c7323e41e11ec989c0242ac110002',2,2,1650000,'2022-06-04 15:56:42',1,NULL);
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
INSERT INTO `orders_detail` VALUES (54,10,1,450000,450000),(54,11,1,600000,600000),(55,10,1,450000,450000),(55,11,1,600000,600000),(88,10,1,450000,450000),(88,11,2,600000,1200000);
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
  `id_product` varchar(50) DEFAULT NULL,
  `product_name` varchar(200) NOT NULL,
  `category_id` int NOT NULL,
  `price` int NOT NULL,
  `quantity` int NOT NULL,
  `image_url` text NOT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_products_category` (`category_id`),
  CONSTRAINT `fk_products_category` FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `products`
--

LOCK TABLES `products` WRITE;
/*!40000 ALTER TABLE `products` DISABLE KEYS */;
INSERT INTO `products` VALUES (1,'4fa6503fe26411ec93820242ac110002','Laptop',1,5000000,1,'http://image.com/default'),(2,'50a10cb2e26411ec93820242ac110002','Handphone',1,2000000,12,'http://image.com/default'),(3,'5158d097e26411ec93820242ac110002','Tab',1,3000000,13,'http://image.com/default'),(4,'5219bceae26411ec93820242ac110002','Sport Shoes',2,1000000,4,'http://image.com/default'),(5,'539d750fe26411ec93820242ac110002','Cotton Bacon',3,50000,9,'http://image.com/default'),(6,'5481fc47e26411ec93820242ac110002','Vape Thelema',3,500000,22,'http://image.com/default'),(7,'560cc7a7e26411ec93820242ac110002','Liquid 100 ml',3,100000,44,'http://image.com/default'),(10,'49028a36e26411ec93820242ac110002','Carrier 40L',4,450000,7,'https://olshop.sgp1.digitaloceanspaces.com/products/XVlBzgbaiC.png'),(11,'0fc1f235e26311ec93820242ac110002','Carrier 60L',4,600000,6,'https://olshop.sgp1.digitaloceanspaces.com/products/XVlBzgbaiC.png'),(12,'89593b52e5a011eca5000242ac110002','Carrier 100L',4,1000000,5,'http://image.com/default');
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
  `id_role` varchar(36) DEFAULT NULL,
  `role_name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,'7f055956e27711ec93820242ac110002','admin'),(2,'7ff69d58e27711ec93820242ac110002','user'),(3,'80eb20a9e27711ec93820242ac110002','merchant'),(5,'f18226cbe27a11ec93820242ac110002','manager');
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
  `id_status_order` varchar(36) DEFAULT NULL,
  `status_name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `status_order`
--

LOCK TABLES `status_order` WRITE;
/*!40000 ALTER TABLE `status_order` DISABLE KEYS */;
INSERT INTO `status_order` VALUES (1,'99acfec1e27711ec93820242ac110002','Menunggu Pembayaran'),(2,'9a89f7ffe27711ec93820242ac110002','Menunggu Konfirmasi'),(3,'9ba3e4b2e27711ec93820242ac110002','Pembayaran Tervalidasi');
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
  `id_user` varchar(36) DEFAULT NULL,
  `username` varchar(100) NOT NULL,
  `email` varchar(200) NOT NULL,
  `password` varchar(225) NOT NULL,
  `image_url` varchar(150) DEFAULT NULL,
  `role_id` int NOT NULL DEFAULT '2',
  PRIMARY KEY (`id`),
  KEY `fk_user_role` (`role_id`),
  CONSTRAINT `fk_user_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'1323c6c4e3cb11ec989c0242ac110002','faridlan','faridlan@gmail.com','rahasia123',NULL,2),(2,'14a184a4e3cb11ec989c0242ac110002','udin','udin@gmail.com','rahasia123','https://olshop.sgp1.digitaloceanspaces.com/profiles/anjing-goblog1.png',2),(4,'16b81d6ee3cb11ec989c0242ac110002','poi','poi@gmail.com','rahasia123',NULL,2),(5,'17996767e3cb11ec989c0242ac110002','poilahhh','poii@gmail.com','rahasia123',NULL,2),(7,'19394f57e3cb11ec989c0242ac110002','admin','admin@admin.com','admin',NULL,1),(8,'1a158c5be3cb11ec989c0242ac110002','Faridlan123','faridlannulhakim@gmail.com','rahasia123',NULL,2),(9,'1aecd4bde3cb11ec989c0242ac110002','jhondoe','jhon@gmail.com','rahasia123',NULL,2),(10,'1d27af1ce3cb11ec989c0242ac110002','jane','jane@gmail.com','rahasia123',NULL,2),(11,'1d68e49ce5b711eca5000242ac110002','kaneki','kaneki@gmail.com','rahasia123',NULL,2),(12,'93d23635e5b711eca5000242ac110002','kanekiKen','kaneki@gmail.com','rahasia123',NULL,2),(13,'40efc363e5b811eca5000242ac110002','kanekiKenKen','kaneki@gmail.com','rahasia123',NULL,2),(14,'83d76aeae5b811eca5000242ac110002','kanekiKenKen','kaneki@gmail.com','rahasia123',NULL,2),(15,'62517760e5b911eca5000242ac110002','kanekiKenKen123','kaneki@gmail.com','rahasia123',NULL,2),(16,'7fcc8dcbe5ba11eca5000242ac110002','kanekiKenKen12345','kaneki@gmail.com','rahasia123',NULL,2),(17,'7c76c6fce5be11ecac710242ac130002','Maimai12345','maimai@gmail.com','rahasia123',NULL,2),(18,'9e1c4458e66d11eca1720242ac130002','anjay','anjay@gmail.com','anjay123',NULL,2),(19,'fa4c63d4e66d11eca1720242ac130002','anjay','anjay@gmail.com','anjay123456',NULL,2),(20,'856ba6d2e66e11eca1720242ac130002','fooo','fooo@gmail.com','fooo123456',NULL,2),(21,'f20c5320e67411eca1720242ac130002','hihi','hihi@gmail.com','hihiihihihihihi',NULL,2);
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

-- Dump completed on 2022-06-07 15:26:40
