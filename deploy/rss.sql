/*
 Navicat Premium Data Transfer

 Source Server         : 本地mariadb
 Source Server Type    : MariaDB
 Source Server Version : 100511
 Source Host           : localhost:3306
 Source Schema         : rss

 Target Server Type    : MariaDB
 Target Server Version : 100511
 File Encoding         : 65001

 Date: 20/11/2021 21:49:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

create database if not exists `rss` default character set utf8mb4 collate utf8mb4_bin;
use `rss`;

-- ----------------------------
-- Table structure for feed
-- ----------------------------
DROP TABLE IF EXISTS `feed`;
CREATE TABLE `feed` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `name` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '源名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='rss 源表';

-- ----------------------------
-- Table structure for news
-- ----------------------------
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` datetime NOT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  `title` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '新闻名称',
  `link` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '链接地址',
  `publish_time` datetime NOT NULL COMMENT '发布时间',
  `feed_id` int(11) NOT NULL DEFAULT 0 COMMENT '源id',
  `feed_name` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '源名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=660 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='新闻表';

SET FOREIGN_KEY_CHECKS = 1;
