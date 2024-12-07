# Paymentengine

## 启动

设置IDEA环境变量，30秒保活

```
ENGINE_KEEP_LIVE_SCAN_INTERVAL=30;ENGINE_KEEP_LIVE_TIME_INTERVAL=10;LOG_LEVEL=debug
```

## Swagger
### 更新接口文档
swag init -g cmd/main.go

## SQL

### 创建引擎数据库表

```sql
/*
 Navicat Premium Data Transfer

 Source Server         : 本地数据库
 Source Server Type    : MySQL
 Source Server Version : 80039 (8.0.39)
 Source Host           : localhost:3306
 Source Schema         : paymentengine

 Target Server Type    : MySQL
 Target Server Version : 80039 (8.0.39)
 File Encoding         : 65001

 Date: 07/12/2024 00:35:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for bank_account
-- ----------------------------
DROP TABLE IF EXISTS `bank_account`;
CREATE TABLE `bank_account` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `account` varchar(32) COLLATE utf8mb4_general_ci NOT NULL COMMENT '银行账户名',
  `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL COMMENT '银行账户密码',
  `phone` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '银行账户手机',
  `provider` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商业银行标识',
  `status` tinyint unsigned NOT NULL DEFAULT '2' COMMENT '1. 启用 2.禁用 4.异常',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for bank_account_card
-- ----------------------------
DROP TABLE IF EXISTS `bank_account_card`;
CREATE TABLE `bank_account_card` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `bank_account_id` bigint NOT NULL COMMENT '银行账户ID',
  `card_number` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '银行卡号',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '银行卡密码',
  `currency_code` varchar(32) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '银行卡币种',
  `remark` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '备注',
  `status` tinyint unsigned NOT NULL DEFAULT '2' COMMENT '1. 启用 2.禁用 4.异常',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_bank_number` (`bank_account_id`,`card_number`),
  KEY `idx_bank_account_id` (`bank_account_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for bank_account_tag
-- ----------------------------
DROP TABLE IF EXISTS `bank_account_tag`;
CREATE TABLE `bank_account_tag` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `tag` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签',
  `provider` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '商业银行标识',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for bank_account_tag_relations
-- ----------------------------
DROP TABLE IF EXISTS `bank_account_tag_relations`;
CREATE TABLE `bank_account_tag_relations` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `bank_account_id` bigint NOT NULL,
  `tag_id` bigint NOT NULL COMMENT '标签',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_unique_relation` (`bank_account_id`,`tag_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
```

```sql
INSERT INTO `paymentengine`.`bank_account` (`id`, `account`, `password`, `phone`, `provider`, `status`, `created_at`, `updated_at`) VALUES (1, 'natpon576509', 'Abc116688@', '', 'thailand_scb', 1, 1733315279, 1733315279);
```

### 测试方法

在 Swagger 页面中，[登录银行账户](http://localhost:7068/swagger/index.html#/%E9%93%B6%E8%A1%8C%E8%B4%A6%E6%88%B7/post_api_v1_bank_account_login) （account_id = 1）

![image-20241207003837371](./docs/image-20241207003837371.png)

登录后观察`bank_account_card`是否新增银行卡列表，并且定时打印KeepLive日志。
