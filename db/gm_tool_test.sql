-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Oct 01, 2021 at 05:21 AM
-- Server version: 10.4.18-MariaDB
-- PHP Version: 8.0.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gm_tool_test`
--
CREATE DATABASE IF NOT EXISTS `gm_tool_test` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `gm_tool_test`;

-- --------------------------------------------------------

--
-- Table structure for table `permissions`
--

CREATE TABLE `permissions` (
  `id` bigint(20) NOT NULL,
  `permission_name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `permissions`
--

INSERT INTO `permissions` (`id`, `permission_name`, `description`, `created_at`, `updated_at`) VALUES
(1, 'modify_user', 'Create, edit, delete users', '2021-05-19 07:13:46', '2021-05-19 07:13:46'),
(2, 'shop_control', 'Create, edit, delete shop data', '2021-06-24 07:53:16', '2021-06-24 07:53:16'),
(3, 'player_report_control', 'modify player and view reports', '2021-07-27 08:41:44', '2021-07-27 08:41:44'),
(4, 'send_mail', 'Send mail', '2021-05-20 10:05:16', '2021-05-20 10:05:16'),
(5, 'matches', 'manage matches', '2021-08-19 11:27:03', '2021-08-19 11:27:03'),
(6, 'ksa_rotation', 'control free ksatriya rotation', '2021-07-27 11:05:38', '2021-07-27 11:05:38'),
(7, 'player_reports', 'read player reports and take action', '2021-08-19 11:27:30', '2021-08-19 11:27:30'),
(8, 'blacklist_player_chat', 'blacklist players from chatting', '2021-08-24 09:25:08', '2021-08-24 09:25:08'),
(9, 'voucher_control', 'generate and manage vouchers', '2021-09-17 08:33:00', '2021-09-17 08:33:00'),
(10, 'judge_control', 'manage judge accounts', '2021-09-20 07:39:07', '2021-09-20 07:39:07'),
(11, 'user_statistics', 'see user statistics', '2021-09-21 12:57:08', '2021-09-21 12:57:08');

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` bigint(20) NOT NULL,
  `role_name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `role_name`, `description`, `created_at`, `updated_at`) VALUES
(1, 'super_admin', 'Super admin', '2021-05-18 11:51:13', '2021-05-18 11:51:13'),
(2, 'marketing', 'Marketing', '2021-05-19 07:02:08', '2021-05-19 07:02:08'),
(3, 'gm', 'Game Master', '2021-05-19 07:07:41', '2021-05-19 07:07:41');

-- --------------------------------------------------------

--
-- Table structure for table `roles_permissions`
--

CREATE TABLE `roles_permissions` (
  `id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL,
  `permission_id` bigint(20) NOT NULL,
  `created_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `roles_permissions`
--

INSERT INTO `roles_permissions` (`id`, `role_id`, `permission_id`, `created_at`) VALUES
(1, 1, 1, '2021-05-20 11:28:40'),
(4, 1, 4, '2021-06-09 09:38:39'),
(6, 1, 2, '2021-06-24 07:59:39'),
(7, 1, 3, '2021-07-27 08:44:59'),
(8, 1, 6, '2021-07-27 11:07:57'),
(9, 1, 5, '2021-08-20 13:34:26'),
(10, 1, 7, '2021-08-20 13:34:26'),
(11, 1, 8, '2021-09-06 12:54:52'),
(12, 1, 9, '2021-09-17 08:34:38'),
(13, 1, 10, '2021-09-20 07:39:36'),
(14, 1, 11, '2021-10-01 05:18:20');

-- --------------------------------------------------------

--
-- Table structure for table `t_news_images`
--

CREATE TABLE `t_news_images` (
  `id` int(11) NOT NULL,
  `image_name` varchar(255) NOT NULL,
  `image_checksum` varchar(255) NOT NULL,
  `uploader` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`) VALUES
(1, 'test1', 'test1@email.com', '$2a$10$EE2./3JeiXBb4/.8LikuE.aHuT1czQ4jZi/G8JXGd3iBRu48t9MQO'),
(4, 'test2', 'test2@email.com', '$2a$10$ir2RQ8YF6NyVjXsrFMMVQuZP3AvrM5PBNRRabBvXnmjX.z4ccgkuu');

-- --------------------------------------------------------

--
-- Table structure for table `users_roles`
--

CREATE TABLE `users_roles` (
  `user_id` bigint(20) NOT NULL,
  `role_id` bigint(20) DEFAULT NULL,
  `created_at` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users_roles`
--

INSERT INTO `users_roles` (`user_id`, `role_id`, `created_at`) VALUES
(1, 1, '2021-06-02 06:38:23');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `permissions`
--
ALTER TABLE `permissions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `roles_permissions`
--
ALTER TABLE `roles_permissions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `roles_permissions_ibfk_1` (`role_id`),
  ADD KEY `roles_permissions_ibfk_2` (`permission_id`);

--
-- Indexes for table `t_news_images`
--
ALTER TABLE `t_news_images`
  ADD PRIMARY KEY (`id`),
  ADD KEY `uploader` (`uploader`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indexes for table `users_roles`
--
ALTER TABLE `users_roles`
  ADD PRIMARY KEY (`user_id`),
  ADD KEY `role_id` (`role_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `permissions`
--
ALTER TABLE `permissions`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=12;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `roles_permissions`
--
ALTER TABLE `roles_permissions`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `t_news_images`
--
ALTER TABLE `t_news_images`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `roles_permissions`
--
ALTER TABLE `roles_permissions`
  ADD CONSTRAINT `roles_permissions_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `roles_permissions_ibfk_2` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `t_news_images`
--
ALTER TABLE `t_news_images`
  ADD CONSTRAINT `t_news_images_ibfk_1` FOREIGN KEY (`uploader`) REFERENCES `users` (`id`);

--
-- Constraints for table `users_roles`
--
ALTER TABLE `users_roles`
  ADD CONSTRAINT `users_roles_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `users_roles_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE SET NULL ON UPDATE SET NULL;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

CREATE USER 'anantarupa'@'localhost' IDENTIFIED BY '8n8nt8rup8';
GRANT ALL PRIVILEGES ON * . * TO 'anantarupa'@'localhost';
FLUSH PRIVILEGES;