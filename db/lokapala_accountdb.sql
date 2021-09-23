-- phpMyAdmin SQL Dump
-- version 5.1.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Sep 06, 2021 at 11:26 AM
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
-- Database: `lokapala_accountdb`
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
(2, 'shop_control', 'Create, edit, delete shop data',  '2021-06-24 07:53:16', '2021-06-24 07:53:16'),
(3, 'player_report_control', 'modify player and view reports', '2021-07-27 08:41:44', '2021-07-27 08:41:44'),
(4, 'send_mail', 'Send mail', '2021-05-20 10:05:16', '2021-05-20 10:05:16'),
(5, 'matches', 'manage matches', '2021-08-19 11:27:03', '2021-08-19 11:27:03'),
(6, 'ksa_rotation', 'control free ksatriya rotation', '2021-07-27 11:05:38', '2021-07-27 11:05:38'),
(7, 'player_reports', 'read player reports and take action', '2021-08-19 11:27:30', '2021-08-19 11:27:30'),
(8, 'blacklist_player_chat', 'blacklist players from chatting', '2021-08-24 09:25:08', '2021-08-24 09:25:08'),
(9, 'voucher', 'generate and manage vouchers', '2021-09-17 08:33:00', '2021-09-17 08:33:00'),
(10, 'judge_control', 'manage judge accounts', '2021-09-17 08:33:00', '2021-09-17 08:33:00'),
(11, 'user_statistics', 'see user statistics', '2021-09-21 12:57:08', '2021-09-21 12:57:08');

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` bigint(20) NOT NULL,
  `role_name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `active` tinyint(4) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `role_name`, `description`, `active`, `created_at`, `updated_at`) VALUES
(1, 'super_admin', 'Super admin', 1, '2021-05-18 11:51:13', '2021-05-18 11:51:13'),
(2, 'marketing', 'Marketing', 1, '2021-05-19 07:02:08', '2021-05-19 07:02:08'),
(3, 'gm', 'Game Master', 1, '2021-05-19 07:07:41', '2021-05-19 07:07:41');

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
(11, 1, 8, '2021-08-20 13:34:26'),
(12, 1, 9, '2021-09-17 08:34:38'),
(13, 1, 10, '2021-09-20 07:39:36'),
(14, 1, 11, '2021-09-21 12:57:08');
-- --------------------------------------------------------

--
-- Table structure for table `t_box`
--

CREATE TABLE `t_box` (
  `box_id` int(11) NOT NULL,
  `box_name` varchar(255) NOT NULL,
  `rand_value` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_box_loot_table`
--

CREATE TABLE `t_box_loot_table` (
  `uid` int(11) NOT NULL,
  `box_id` int(11) NOT NULL,
  `item_id` int(11) NOT NULL,
  `item_type` int(11) NOT NULL,
  `amount` int(11) NOT NULL DEFAULT 1,
  `chance` int(11) NOT NULL,
  `min` int(11) NOT NULL,
  `max` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_chest`
--

CREATE TABLE `t_chest` (
  `duration` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `t_chest`
--

INSERT INTO `t_chest` (`duration`) VALUES
(3),
(6),
(12);

-- --------------------------------------------------------

--
-- Table structure for table `t_country`
--

CREATE TABLE `t_country` (
  `country_id` int(11) NOT NULL,
  `country_code` varchar(2) DEFAULT NULL,
  `country_name` varchar(100) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `t_country`
--

INSERT INTO `t_country` (`country_id`, `country_code`, `country_name`) VALUES
(102, 'ID', 'Indonesia');

-- --------------------------------------------------------

--
-- Table structure for table `t_currency_type`
--

CREATE TABLE `t_currency_type` (
  `currency_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `t_currency_type`
--

INSERT INTO `t_currency_type` (`currency_id`, `name`) VALUES
(1, 'Ori'),
(2, 'Citrine'),
(3, 'Lotus');

-- --------------------------------------------------------

--
-- Table structure for table `t_energy`
--

CREATE TABLE `t_energy` (
  `energy_id` int(11) NOT NULL,
  `description` tinytext NOT NULL,
  `target` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `t_energy`
--

INSERT INTO `t_energy` (`energy_id`, `description`, `target`) VALUES
(1, '10 Energy', 10),
(2, '35 Energy', 35),
(4, '50 Energy', 50),
(8, '75 Energy', 75),
(16, '225 Energy', 225),
(32, '500 Energy', 500);

-- --------------------------------------------------------

--
-- Table structure for table `t_gacha`
--

CREATE TABLE `t_gacha` (
  `gacha_id` int(11) NOT NULL,
  `start_date` datetime NOT NULL,
  `end_date` datetime NOT NULL,
  `random_value` int(11) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_gacha_feature`
--

CREATE TABLE `t_gacha_feature` (
  `gacha_feature_id` int(10) UNSIGNED NOT NULL,
  `gacha_id` int(11) NOT NULL,
  `gacha_item_id` int(11) NOT NULL,
  `priority` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_gacha_item`
--

CREATE TABLE `t_gacha_item` (
  `gacha_item_id` int(11) NOT NULL,
  `item_type` int(11) NOT NULL,
  `item_id` int(11) NOT NULL,
  `amount` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_gacha_loot_table`
--

CREATE TABLE `t_gacha_loot_table` (
  `gacha_id` int(11) NOT NULL,
  `gacha_item_id` int(11) NOT NULL,
  `chance` int(11) NOT NULL,
  `min_value` int(11) NOT NULL,
  `max_value` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_icon_avatar`
--

CREATE TABLE `t_icon_avatar` (
  `avatar_id` int(11) NOT NULL,
  `description` tinytext NOT NULL,
  `release_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `t_icon_avatar`
--

INSERT INTO `t_icon_avatar` (`avatar_id`, `description`, `release_date`) VALUES
(1, 'avatar1', '2021-08-26 09:47:10');

-- --------------------------------------------------------

--
-- Table structure for table `t_icon_frame`
--

CREATE TABLE `t_icon_frame` (
  `frame_id` int(11) NOT NULL,
  `description` tinytext NOT NULL,
  `release_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `t_icon_frame`
--

INSERT INTO `t_icon_frame` (`frame_id`, `description`, `release_date`) VALUES
(0, 'none', '2021-08-26 09:46:05');

-- --------------------------------------------------------

--
-- Table structure for table `t_item_type`
--

CREATE TABLE `t_item_type` (
  `item_type_id` int(11) NOT NULL,
  `item_type_name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `t_item_type`
--

INSERT INTO `t_item_type` (`item_type_id`, `item_type_name`) VALUES
(1, 'currency'),
(2, 'ksatriya'),
(3, 'Skin'),
(4, 'Rune'),
(5, 'Item'),
(6, 'Box'),
(7, 'Chest'),
(8, 'Energy'),
(9, 'Skin Part'),
(10, 'Premium'),
(11, 'Frame'),
(12, 'Avatar'),
(13, 'Vikara'),
(14, 'Vahana'),
(15, 'Ksatriya Fragment'),
(16, 'Skin Fragment');

-- --------------------------------------------------------

--
-- Table structure for table `t_ksatriya`
--

CREATE TABLE `t_ksatriya` (
  `ksatriya_id` smallint(6) NOT NULL,
  `role` varchar(255) DEFAULT NULL,
  `release_date` timestamp NULL DEFAULT current_timestamp(),
  `ksatriya_name` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_ksatriya_fragment`
--

CREATE TABLE `t_ksatriya_fragment` (
  `ksatriya_id` smallint(6) NOT NULL,
  `amount_needed` int(11) NOT NULL,
  `sell_currency_id` int(11) NOT NULL,
  `sell_value` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_ksatriya_rotation`
--

CREATE TABLE `t_ksatriya_rotation` (
  `ksatriya_rotation_id` int(11) NOT NULL,
  `ksatriya_id` smallint(6) NOT NULL,
  `start_date` datetime NOT NULL,
  `end_date` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_ksatriya_skin`
--

CREATE TABLE `t_ksatriya_skin` (
  `ksatriya_skin_id` int(11) NOT NULL,
  `ksatriya_id` smallint(6) NOT NULL,
  `release_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_ksatriya_skin_fragment`
--

CREATE TABLE `t_ksatriya_skin_fragment` (
  `ksatriya_skin_id` int(11) NOT NULL,
  `amount_needed` int(11) NOT NULL,
  `sell_currency_id` int(11) NOT NULL,
  `sell_value` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_ksatriya_skin_part`
--

CREATE TABLE `t_ksatriya_skin_part` (
  `skin_part_id` int(11) NOT NULL,
  `release_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_lotto`
--

CREATE TABLE `t_lotto` (
  `lotto_id` int(11) NOT NULL,
  `start_date` datetime NOT NULL,
  `end_date` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_lotto_feature`
--

CREATE TABLE `t_lotto_feature` (
  `lotto_feature_id` int(11) NOT NULL,
  `lotto_id` int(11) NOT NULL,
  `lotto_item_id` int(11) NOT NULL,
  `priority` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_lotto_item`
--

CREATE TABLE `t_lotto_item` (
  `lotto_item_id` int(11) NOT NULL,
  `item_type` int(11) NOT NULL,
  `item_id` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `color_id` int(11) NOT NULL,
  `default_amount` int(11) NOT NULL,
  `item_name` tinytext NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_lotto_item_color`
--

CREATE TABLE `t_lotto_item_color` (
  `color_id` int(11) NOT NULL,
  `color_name` tinytext NOT NULL,
  `weight` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_lotto_loot_table`
--

CREATE TABLE `t_lotto_loot_table` (
  `lotto_id` int(11) NOT NULL,
  `lotto_item_id` int(11) NOT NULL,
  `amount` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_mail`
--

CREATE TABLE `t_mail` (
  `mail_id` int(11) NOT NULL,
  `mail_type` varchar(255) NOT NULL,
  `sender_id` int(11) DEFAULT NULL,
  `reciever_id` int(11) NOT NULL,
  `send_date` datetime NOT NULL DEFAULT current_timestamp(),
  `mail_template` int(11) DEFAULT NULL,
  `confirm_read` tinyint(4) NOT NULL DEFAULT 0,
  `read_date` datetime DEFAULT NULL,
  `confirn_claim` tinyint(4) NOT NULL DEFAULT 0,
  `claim_date` datetime DEFAULT NULL,
  `parameter` varchar(255) DEFAULT NULL,
  `custom_message_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `t_mail`
--

INSERT INTO `t_mail` (`mail_id`, `mail_type`, `sender_id`, `reciever_id`, `send_date`, `mail_template`, `confirm_read`, `read_date`, `confirn_claim`, `claim_date`, `parameter`, `custom_message_id`) VALUES
(1, 'Friend', NULL, 1, '2021-08-27 14:13:03', 1, 0, '2021-08-27 14:13:03', 0, '2021-08-27 14:13:03', NULL, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `t_mail_attachment`
--

CREATE TABLE `t_mail_attachment` (
  `id` int(11) NOT NULL,
  `template_id` int(11) DEFAULT NULL,
  `item_id` int(11) DEFAULT NULL,
  `item_type` int(11) DEFAULT NULL,
  `amount` int(11) NOT NULL,
  `custom_message_id` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `t_mail_attachment`
--

INSERT INTO `t_mail_attachment` (`id`, `template_id`, `item_id`, `item_type`, `amount`, `custom_message_id`) VALUES
(1, 5, 1, 1, 100000, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `t_mail_custom_message`
--

CREATE TABLE `t_mail_custom_message` (
  `message_id` int(11) NOT NULL,
  `subject` varchar(255) NOT NULL,
  `message` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_mail_login`
--

CREATE TABLE `t_mail_login` (
  `template_id` int(11) NOT NULL,
  `parameter` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL,
  `start_date` timestamp NOT NULL DEFAULT current_timestamp(),
  `end_date` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_mail_template`
--

CREATE TABLE `t_mail_template` (
  `template_id` int(11) NOT NULL,
  `subject` tinytext NOT NULL,
  `message` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `t_mail_template`
--

INSERT INTO `t_mail_template` (`template_id`, `subject`, `message`) VALUES
(1, 'test subject', 'this is a test message'),
(3, 'test subject no 32', 'this is a test message 32'),
(4, 'test subject no 3', 'this is a test message 3'),
(5, 'test subject no 4', 'this is a test message 4');

-- --------------------------------------------------------

--
-- Table structure for table `t_news`
--

CREATE TABLE `t_news` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `release_date` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `type` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_news_detail`
--

CREATE TABLE `t_news_detail` (
  `news_id` int(11) NOT NULL,
  `lang` char(2) NOT NULL,
  `title` varchar(255) NOT NULL,
  `banner` varchar(15) DEFAULT NULL,
  `banner_checksum` varchar(32) DEFAULT NULL,
  `content` varchar(15) NOT NULL,
  `content_checksum` varchar(32) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

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
-- Table structure for table `t_news_type`
--

CREATE TABLE `t_news_type` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_premium`
--

CREATE TABLE `t_premium` (
  `item_id` int(11) NOT NULL,
  `duration` int(11) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_rune`
--

CREATE TABLE `t_rune` (
  `rune_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL DEFAULT 'Default',
  `rune_color` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_shop`
--

CREATE TABLE `t_shop` (
  `shop_id` int(11) NOT NULL,
  `item_id` int(11) DEFAULT NULL,
  `item_type` int(11) DEFAULT NULL,
  `amount` int(11) NOT NULL DEFAULT 1,
  `price_coin` int(11) DEFAULT NULL,
  `price_citrine` int(11) DEFAULT NULL,
  `price_lotus` int(11) DEFAULT NULL,
  `release_date` datetime DEFAULT NULL,
  `description` tinytext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `t_shop`
--

INSERT INTO `t_shop` (`shop_id`, `item_id`, `item_type`, `amount`, `price_coin`, `price_citrine`, `price_lotus`, `release_date`, `description`) VALUES
(1, 2, 2, 1, NULL, 100000, NULL, '2021-07-06 07:23:22', '0');

-- --------------------------------------------------------

--
-- Table structure for table `t_shop_bundle`
--

CREATE TABLE `t_shop_bundle` (
  `shop_id` int(11) NOT NULL,
  `item_type` int(11) NOT NULL,
  `item_id` int(11) NOT NULL,
  `amount` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_shop_lotus`
--

CREATE TABLE `t_shop_lotus` (
  `shop_lotus_period_id` int(11) NOT NULL,
  `shop_lotus_item_id` int(11) NOT NULL,
  `player_limit` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_shop_lotus_item`
--

CREATE TABLE `t_shop_lotus_item` (
  `shop_lotus_item_id` int(11) NOT NULL,
  `item_type` int(11) NOT NULL,
  `item_id` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `price` int(11) NOT NULL,
  `default_limit` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_shop_lotus_period`
--

CREATE TABLE `t_shop_lotus_period` (
  `shop_lotus_period_id` int(11) NOT NULL,
  `start_date` datetime NOT NULL,
  `end_date` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------

--
-- Table structure for table `t_user`
--

CREATE TABLE `t_user` (
  `user_id` int(11) NOT NULL,
  `user_name` varchar(45) NOT NULL,
  `avatar_icon` int(11) NOT NULL DEFAULT 1,
  `karma` tinyint(4) NOT NULL DEFAULT 100,
  `gender` enum('M','F') DEFAULT NULL,
  `country` int(11) NOT NULL DEFAULT 102,
  `role` int(11) NOT NULL DEFAULT 0,
  `playing_time` int(11) NOT NULL DEFAULT 0,
  `frame` int(11) NOT NULL DEFAULT 0
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `t_user`
--

INSERT INTO `t_user` (`user_id`, `user_name`, `avatar_icon`, `karma`, `gender`, `country`, `role`, `playing_time`, `frame`) VALUES
(5, 'player1', 1, 100, NULL, 102, 0, 0, 0),
(6, 'player2', 1, 100, NULL, 102, 0, 0, 0);

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
  `id` bigint(4) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `role_id` bigint(20) NOT NULL,
  `created_at` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users_roles`
--

INSERT INTO `users_roles` (`id`, `user_id`, `role_id`, `created_at`) VALUES
(1, 1, 1, '2021-06-02 06:38:23');

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
  ADD KEY `role_id` (`role_id`),
  ADD KEY `permission_id` (`permission_id`);

--
-- Indexes for table `t_box`
--
ALTER TABLE `t_box`
  ADD PRIMARY KEY (`box_id`);

--
-- Indexes for table `t_box_loot_table`
--
ALTER TABLE `t_box_loot_table`
  ADD PRIMARY KEY (`uid`),
  ADD KEY `box_id` (`box_id`),
  ADD KEY `item_type` (`item_type`);

--
-- Indexes for table `t_chest`
--
ALTER TABLE `t_chest`
  ADD PRIMARY KEY (`duration`);

--
-- Indexes for table `t_country`
--
ALTER TABLE `t_country`
  ADD PRIMARY KEY (`country_id`);

--
-- Indexes for table `t_currency_type`
--
ALTER TABLE `t_currency_type`
  ADD PRIMARY KEY (`currency_id`);

--
-- Indexes for table `t_energy`
--
ALTER TABLE `t_energy`
  ADD PRIMARY KEY (`energy_id`);

--
-- Indexes for table `t_gacha`
--
ALTER TABLE `t_gacha`
  ADD PRIMARY KEY (`gacha_id`);

--
-- Indexes for table `t_gacha_feature`
--
ALTER TABLE `t_gacha_feature`
  ADD PRIMARY KEY (`gacha_feature_id`),
  ADD KEY `gacha_id` (`gacha_id`),
  ADD KEY `gacha_item_id` (`gacha_item_id`);

--
-- Indexes for table `t_gacha_item`
--
ALTER TABLE `t_gacha_item`
  ADD PRIMARY KEY (`gacha_item_id`),
  ADD KEY `item_type` (`item_type`);

--
-- Indexes for table `t_gacha_loot_table`
--
ALTER TABLE `t_gacha_loot_table`
  ADD PRIMARY KEY (`gacha_id`,`gacha_item_id`),
  ADD KEY `gacha_item_id` (`gacha_item_id`);

--
-- Indexes for table `t_icon_avatar`
--
ALTER TABLE `t_icon_avatar`
  ADD PRIMARY KEY (`avatar_id`);

--
-- Indexes for table `t_icon_frame`
--
ALTER TABLE `t_icon_frame`
  ADD PRIMARY KEY (`frame_id`);

--
-- Indexes for table `t_item_type`
--
ALTER TABLE `t_item_type`
  ADD PRIMARY KEY (`item_type_id`);

--
-- Indexes for table `t_ksatriya`
--
ALTER TABLE `t_ksatriya`
  ADD PRIMARY KEY (`ksatriya_id`);

--
-- Indexes for table `t_ksatriya_fragment`
--
ALTER TABLE `t_ksatriya_fragment`
  ADD PRIMARY KEY (`ksatriya_id`);

--
-- Indexes for table `t_ksatriya_rotation`
--
ALTER TABLE `t_ksatriya_rotation`
  ADD PRIMARY KEY (`ksatriya_rotation_id`),
  ADD KEY `ksatriya_id` (`ksatriya_id`);

--
-- Indexes for table `t_ksatriya_skin`
--
ALTER TABLE `t_ksatriya_skin`
  ADD PRIMARY KEY (`ksatriya_skin_id`),
  ADD KEY `ksatriya_id` (`ksatriya_id`);

--
-- Indexes for table `t_ksatriya_skin_fragment`
--
ALTER TABLE `t_ksatriya_skin_fragment`
  ADD PRIMARY KEY (`ksatriya_skin_id`);

--
-- Indexes for table `t_ksatriya_skin_part`
--
ALTER TABLE `t_ksatriya_skin_part`
  ADD PRIMARY KEY (`skin_part_id`);

--
-- Indexes for table `t_lotto`
--
ALTER TABLE `t_lotto`
  ADD PRIMARY KEY (`lotto_id`);

--
-- Indexes for table `t_lotto_feature`
--
ALTER TABLE `t_lotto_feature`
  ADD PRIMARY KEY (`lotto_feature_id`),
  ADD KEY `lotto_id` (`lotto_id`),
  ADD KEY `lotto_item_id` (`lotto_item_id`);

--
-- Indexes for table `t_lotto_item`
--
ALTER TABLE `t_lotto_item`
  ADD PRIMARY KEY (`lotto_item_id`),
  ADD KEY `color_id` (`color_id`),
  ADD KEY `item_type` (`item_type`);

--
-- Indexes for table `t_lotto_item_color`
--
ALTER TABLE `t_lotto_item_color`
  ADD PRIMARY KEY (`color_id`);

--
-- Indexes for table `t_lotto_loot_table`
--
ALTER TABLE `t_lotto_loot_table`
  ADD PRIMARY KEY (`lotto_id`,`lotto_item_id`),
  ADD KEY `lotto_item_id` (`lotto_item_id`);

--
-- Indexes for table `t_mail`
--
ALTER TABLE `t_mail`
  ADD PRIMARY KEY (`mail_id`),
  ADD KEY `mail_template` (`mail_template`),
  ADD KEY `custom_message_id` (`custom_message_id`);

--
-- Indexes for table `t_mail_attachment`
--
ALTER TABLE `t_mail_attachment`
  ADD PRIMARY KEY (`id`),
  ADD KEY `template_id` (`template_id`),
  ADD KEY `custom_message_id` (`custom_message_id`),
  ADD KEY `item_type` (`item_type`);

--
-- Indexes for table `t_mail_custom_message`
--
ALTER TABLE `t_mail_custom_message`
  ADD PRIMARY KEY (`message_id`);

--
-- Indexes for table `t_mail_login`
--
ALTER TABLE `t_mail_login`
  ADD PRIMARY KEY (`template_id`);

--
-- Indexes for table `t_mail_template`
--
ALTER TABLE `t_mail_template`
  ADD PRIMARY KEY (`template_id`);

--
-- Indexes for table `t_news`
--
ALTER TABLE `t_news`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `t_news_detail`
--
ALTER TABLE `t_news_detail`
  ADD PRIMARY KEY (`news_id`,`lang`);

--
-- Indexes for table `t_news_images`
--
ALTER TABLE `t_news_images`
  ADD PRIMARY KEY (`id`),
  ADD KEY `uploader` (`uploader`);

--
-- Indexes for table `t_news_type`
--
ALTER TABLE `t_news_type`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `t_premium`
--
ALTER TABLE `t_premium`
  ADD PRIMARY KEY (`item_id`);

--
-- Indexes for table `t_rune`
--
ALTER TABLE `t_rune`
  ADD PRIMARY KEY (`rune_id`);

--
-- Indexes for table `t_shop`
--
ALTER TABLE `t_shop`
  ADD PRIMARY KEY (`shop_id`),
  ADD KEY `item_type` (`item_type`);

--
-- Indexes for table `t_shop_bundle`
--
ALTER TABLE `t_shop_bundle`
  ADD KEY `shop_id` (`shop_id`),
  ADD KEY `item_type` (`item_type`);

--
-- Indexes for table `t_shop_lotus`
--
ALTER TABLE `t_shop_lotus`
  ADD PRIMARY KEY (`shop_lotus_period_id`,`shop_lotus_item_id`),
  ADD KEY `shop_lotus_period_id` (`shop_lotus_period_id`);

--
-- Indexes for table `t_shop_lotus_item`
--
ALTER TABLE `t_shop_lotus_item`
  ADD PRIMARY KEY (`shop_lotus_item_id`),
  ADD KEY `item_type` (`item_type`);

--
-- Indexes for table `t_shop_lotus_period`
--
ALTER TABLE `t_shop_lotus_period`
  ADD PRIMARY KEY (`shop_lotus_period_id`);

--
-- Indexes for table `t_user`
--
ALTER TABLE `t_user`
  ADD PRIMARY KEY (`user_id`),
  ADD UNIQUE KEY `user_name` (`user_name`),
  ADD KEY `t_user_ibfk_1` (`avatar_icon`),
  ADD KEY `t_user_ibfk_2` (`frame`),
  ADD KEY `country` (`country`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users_roles`
--
ALTER TABLE `users_roles`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `role_id` (`role_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `permissions`
--
ALTER TABLE `permissions`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `roles_permissions`
--
ALTER TABLE `roles_permissions`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `t_box`
--
ALTER TABLE `t_box`
  MODIFY `box_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_box_loot_table`
--
ALTER TABLE `t_box_loot_table`
  MODIFY `uid` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_country`
--
ALTER TABLE `t_country`
  MODIFY `country_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=103;

--
-- AUTO_INCREMENT for table `t_currency_type`
--
ALTER TABLE `t_currency_type`
  MODIFY `currency_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `t_gacha`
--
ALTER TABLE `t_gacha`
  MODIFY `gacha_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_gacha_feature`
--
ALTER TABLE `t_gacha_feature`
  MODIFY `gacha_feature_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_gacha_item`
--
ALTER TABLE `t_gacha_item`
  MODIFY `gacha_item_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_item_type`
--
ALTER TABLE `t_item_type`
  MODIFY `item_type_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;

--
-- AUTO_INCREMENT for table `t_ksatriya`
--
ALTER TABLE `t_ksatriya`
  MODIFY `ksatriya_id` smallint(6) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_ksatriya_rotation`
--
ALTER TABLE `t_ksatriya_rotation`
  MODIFY `ksatriya_rotation_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_ksatriya_skin`
--
ALTER TABLE `t_ksatriya_skin`
  MODIFY `ksatriya_skin_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_lotto`
--
ALTER TABLE `t_lotto`
  MODIFY `lotto_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_lotto_feature`
--
ALTER TABLE `t_lotto_feature`
  MODIFY `lotto_feature_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_lotto_item`
--
ALTER TABLE `t_lotto_item`
  MODIFY `lotto_item_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_lotto_item_color`
--
ALTER TABLE `t_lotto_item_color`
  MODIFY `color_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_mail`
--
ALTER TABLE `t_mail`
  MODIFY `mail_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `t_mail_attachment`
--
ALTER TABLE `t_mail_attachment`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `t_mail_custom_message`
--
ALTER TABLE `t_mail_custom_message`
  MODIFY `message_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_mail_template`
--
ALTER TABLE `t_mail_template`
  MODIFY `template_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `t_news`
--
ALTER TABLE `t_news`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_news_images`
--
ALTER TABLE `t_news_images`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_premium`
--
ALTER TABLE `t_premium`
  MODIFY `item_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_rune`
--
ALTER TABLE `t_rune`
  MODIFY `rune_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_shop`
--
ALTER TABLE `t_shop`
  MODIFY `shop_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `t_shop_lotus_item`
--
ALTER TABLE `t_shop_lotus_item`
  MODIFY `shop_lotus_item_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_shop_lotus_period`
--
ALTER TABLE `t_shop_lotus_period`
  MODIFY `shop_lotus_period_id` int(11) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `t_user`
--
ALTER TABLE `t_user`
  MODIFY `user_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `users_roles`
--
ALTER TABLE `users_roles`
  MODIFY `id` bigint(4) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `roles_permissions`
--
ALTER TABLE `roles_permissions`
  ADD CONSTRAINT `roles_permissions_ibfk_1` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`),
  ADD CONSTRAINT `roles_permissions_ibfk_2` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`);

--
-- Constraints for table `t_box_loot_table`
--
ALTER TABLE `t_box_loot_table`
  ADD CONSTRAINT `t_box_loot_table_ibfk_1` FOREIGN KEY (`box_id`) REFERENCES `t_box` (`box_id`),
  ADD CONSTRAINT `t_box_loot_table_ibfk_2` FOREIGN KEY (`item_type`) REFERENCES `t_item_type` (`item_type_id`);

--
-- Constraints for table `t_gacha_feature`
--
ALTER TABLE `t_gacha_feature`
  ADD CONSTRAINT `t_gacha_feature_ibfk_1` FOREIGN KEY (`gacha_id`) REFERENCES `t_gacha` (`gacha_id`),
  ADD CONSTRAINT `t_gacha_feature_ibfk_2` FOREIGN KEY (`gacha_item_id`) REFERENCES `t_gacha_item` (`gacha_item_id`);

--
-- Constraints for table `t_gacha_item`
--
ALTER TABLE `t_gacha_item`
  ADD CONSTRAINT `t_gacha_item_ibfk_1` FOREIGN KEY (`item_type`) REFERENCES `t_item_type` (`item_type_id`);

--
-- Constraints for table `t_gacha_loot_table`
--
ALTER TABLE `t_gacha_loot_table`
  ADD CONSTRAINT `t_gacha_loot_table_ibfk_1` FOREIGN KEY (`gacha_id`) REFERENCES `t_gacha` (`gacha_id`),
  ADD CONSTRAINT `t_gacha_loot_table_ibfk_2` FOREIGN KEY (`gacha_item_id`) REFERENCES `t_gacha_item` (`gacha_item_id`);

--
-- Constraints for table `t_ksatriya_fragment`
--
ALTER TABLE `t_ksatriya_fragment`
  ADD CONSTRAINT `t_ksatriya_fragment_ibfk_1` FOREIGN KEY (`ksatriya_id`) REFERENCES `t_ksatriya` (`ksatriya_id`);

--
-- Constraints for table `t_ksatriya_rotation`
--
ALTER TABLE `t_ksatriya_rotation`
  ADD CONSTRAINT `t_ksatriya_rotation_ibfk_1` FOREIGN KEY (`ksatriya_id`) REFERENCES `t_ksatriya` (`ksatriya_id`);

--
-- Constraints for table `t_ksatriya_skin`
--
ALTER TABLE `t_ksatriya_skin`
  ADD CONSTRAINT `t_ksatriya_skin_ibfk_1` FOREIGN KEY (`ksatriya_id`) REFERENCES `t_ksatriya` (`ksatriya_id`);

--
-- Constraints for table `t_ksatriya_skin_fragment`
--
ALTER TABLE `t_ksatriya_skin_fragment`
  ADD CONSTRAINT `t_ksatriya_skin_fragment_ibfk_1` FOREIGN KEY (`ksatriya_skin_id`) REFERENCES `t_ksatriya_skin` (`ksatriya_skin_id`);

--
-- Constraints for table `t_lotto_feature`
--
ALTER TABLE `t_lotto_feature`
  ADD CONSTRAINT `t_lotto_feature_ibfk_1` FOREIGN KEY (`lotto_id`) REFERENCES `t_lotto` (`lotto_id`),
  ADD CONSTRAINT `t_lotto_feature_ibfk_2` FOREIGN KEY (`lotto_item_id`) REFERENCES `t_lotto_item` (`lotto_item_id`);

--
-- Constraints for table `t_lotto_item`
--
ALTER TABLE `t_lotto_item`
  ADD CONSTRAINT `t_lotto_item_ibfk_1` FOREIGN KEY (`color_id`) REFERENCES `t_lotto_item_color` (`color_id`),
  ADD CONSTRAINT `t_lotto_item_ibfk_2` FOREIGN KEY (`item_type`) REFERENCES `t_item_type` (`item_type_id`);

--
-- Constraints for table `t_lotto_loot_table`
--
ALTER TABLE `t_lotto_loot_table`
  ADD CONSTRAINT `t_lotto_loot_table_ibfk_1` FOREIGN KEY (`lotto_id`) REFERENCES `t_lotto` (`lotto_id`),
  ADD CONSTRAINT `t_lotto_loot_table_ibfk_2` FOREIGN KEY (`lotto_item_id`) REFERENCES `t_lotto_item` (`lotto_item_id`);

--
-- Constraints for table `t_mail`
--
ALTER TABLE `t_mail`
  ADD CONSTRAINT `t_mail_ibfk_1` FOREIGN KEY (`mail_template`) REFERENCES `t_mail_template` (`template_id`),
  ADD CONSTRAINT `t_mail_ibfk_2` FOREIGN KEY (`custom_message_id`) REFERENCES `t_mail_custom_message` (`message_id`);

--
-- Constraints for table `t_mail_attachment`
--
ALTER TABLE `t_mail_attachment`
  ADD CONSTRAINT `t_mail_attachment_ibfk_1` FOREIGN KEY (`template_id`) REFERENCES `t_mail_template` (`template_id`),
  ADD CONSTRAINT `t_mail_attachment_ibfk_2` FOREIGN KEY (`custom_message_id`) REFERENCES `t_mail_custom_message` (`message_id`),
  ADD CONSTRAINT `t_mail_attachment_ibfk_3` FOREIGN KEY (`item_type`) REFERENCES `t_item_type` (`item_type_id`);

--
-- Constraints for table `t_mail_login`
--
ALTER TABLE `t_mail_login`
  ADD CONSTRAINT `t_mail_login_ibfk_1` FOREIGN KEY (`template_id`) REFERENCES `t_mail_template` (`template_id`);

--
-- Constraints for table `t_news_detail`
--
ALTER TABLE `t_news_detail`
  ADD CONSTRAINT `t_news_detail_ibfk_1` FOREIGN KEY (`news_id`) REFERENCES `t_news` (`id`);

--
-- Constraints for table `t_news_images`
--
ALTER TABLE `t_news_images`
  ADD CONSTRAINT `t_news_images_ibfk_1` FOREIGN KEY (`uploader`) REFERENCES `users` (`id`);

--
-- Constraints for table `t_shop`
--
ALTER TABLE `t_shop`
  ADD CONSTRAINT `t_shop_ibfk_1` FOREIGN KEY (`item_type`) REFERENCES `t_item_type` (`item_type_id`);

--
-- Constraints for table `t_shop_bundle`
--
ALTER TABLE `t_shop_bundle`
  ADD CONSTRAINT `t_shop_bundle_ibfk_1` FOREIGN KEY (`shop_id`) REFERENCES `t_item_type` (`item_type_id`),
  ADD CONSTRAINT `t_shop_bundle_ibfk_2` FOREIGN KEY (`shop_id`) REFERENCES `t_shop` (`shop_id`);

--
-- Constraints for table `t_shop_lotus`
--
ALTER TABLE `t_shop_lotus`
  ADD CONSTRAINT `t_shop_lotus_ibfk_1` FOREIGN KEY (`shop_lotus_item_id`) REFERENCES `t_shop_lotus_item` (`shop_lotus_item_id`),
  ADD CONSTRAINT `t_shop_lotus_ibfk_2` FOREIGN KEY (`shop_lotus_period_id`) REFERENCES `t_shop_lotus_period` (`shop_lotus_period_id`);

--
-- Constraints for table `t_shop_lotus_item`
--
ALTER TABLE `t_shop_lotus_item`
  ADD CONSTRAINT `t_shop_lotus_item_ibfk_1` FOREIGN KEY (`item_type`) REFERENCES `t_item_type` (`item_type_id`);

--
-- Constraints for table `t_user`
--
ALTER TABLE `t_user`
  ADD CONSTRAINT `t_user_ibfk_1` FOREIGN KEY (`avatar_icon`) REFERENCES `t_icon_avatar` (`avatar_id`),
  ADD CONSTRAINT `t_user_ibfk_2` FOREIGN KEY (`frame`) REFERENCES `t_icon_frame` (`frame_id`),
  ADD CONSTRAINT `t_user_ibfk_3` FOREIGN KEY (`country`) REFERENCES `t_country` (`country_id`);

--
-- Constraints for table `users_roles`
--
ALTER TABLE `users_roles`
  ADD CONSTRAINT `users_roles_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `users_roles_ibfk_2` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;

CREATE USER 'anantarupa'@'localhost' IDENTIFIED BY '8n8nt8rup8';
GRANT ALL PRIVILEGES ON * . * TO 'anantarupa'@'localhost';
FLUSH PRIVILEGES;