-- -------------------------------------------------------------
-- TablePlus 3.11.0(352)
--
-- https://tableplus.com/
--
-- Database: hiads
-- Generation Time: 2023-04-20 14:00:42.3200
-- -------------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


DROP TABLE IF EXISTS `account_tokens`;
CREATE TABLE `account_tokens` (
  `id` bigint(15) unsigned NOT NULL AUTO_INCREMENT,
  `account_id` bigint(15) unsigned NOT NULL DEFAULT '0' COMMENT '账户ID',
  `advertiser_id` varchar(30) NOT NULL DEFAULT '' COMMENT '广告主账户ID',
  `access_token` varchar(1000) NOT NULL DEFAULT '',
  `refresh_token` varchar(1000) NOT NULL DEFAULT '',
  `expired_at` datetime NOT NULL COMMENT 'access_token 过期时间',
  `created_at` datetime NOT NULL COMMENT '添加时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次修改时间',
  `token_type` varchar(20) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_account_id` (`account_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
  `id` bigint(15) unsigned NOT NULL AUTO_INCREMENT,
  `account_name` varchar(50) NOT NULL DEFAULT '' COMMENT '账户名',
  `parent_id` bigint(15) unsigned NOT NULL DEFAULT '0' COMMENT '所属上级服务商',
  `advertiser_id` varchar(30) NOT NULL DEFAULT '' COMMENT '广告主账户ID',
  `developer_id` varchar(30) NOT NULL DEFAULT '' COMMENT '开发者ID',
  `account_type` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '账户类型',
  `state` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '状态',
  `is_auth` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '是否已认证',
  `client_id` varchar(15) NOT NULL DEFAULT '' COMMENT '客户端ID',
  `secret` varchar(70) NOT NULL DEFAULT '' COMMENT '密钥',
  `created_at` datetime NOT NULL COMMENT '添加时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `app_accounts`;
CREATE TABLE `app_accounts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `account_type` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '账户类型',
  `app_id` varchar(32) NOT NULL DEFAULT '',
  `account_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_idx` (`app_id`,`account_id`,`account_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=812 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `apps`;
CREATE TABLE `apps` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` varchar(32) NOT NULL DEFAULT '' COMMENT '第三方应用ID，例如华为APP ID : C10134672；可能存在GP的应用ID 32位',
  `app_name` varchar(128) NOT NULL COMMENT '应用名称',
  `pkg_name` varchar(128) NOT NULL DEFAULT '' COMMENT '应用包名或BundleID',
  `channel` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '系统平台(渠道)：华为 AppGallery；GooglePlay; AppStore',
  `tags` varchar(500) NOT NULL DEFAULT '' COMMENT '应用标签',
  `icon_url` varchar(1000) NOT NULL DEFAULT '' COMMENT '图标',
  `product_id` varchar(32) NOT NULL DEFAULT '' COMMENT '产品ID，创建任务时需要',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `app_id` (`app_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=111934 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `assets`;
CREATE TABLE `assets` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `account_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '账户ID;对应 accounts 表的id字段',
  `advertiser_id` varchar(30) NOT NULL DEFAULT '' COMMENT '广告主账户ID',
  `app_id` varchar(32) NOT NULL DEFAULT '' COMMENT '第三方应用ID，例如华为APP ID : C10134672；可能存在GP的应用ID 32位',
  `asset_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `asset_name` varchar(50) NOT NULL DEFAULT '',
  `asset_type` varchar(50) NOT NULL DEFAULT '',
  `file_url` varchar(1000) NOT NULL DEFAULT '',
  `width` smallint(5) unsigned NOT NULL DEFAULT '0',
  `height` smallint(5) unsigned NOT NULL DEFAULT '0',
  `video_play_duration` bigint(15) unsigned NOT NULL DEFAULT '0',
  `file_size` bigint(15) unsigned NOT NULL DEFAULT '0',
  `file_format` varchar(50) NOT NULL DEFAULT '',
  `file_hash_sha256` varchar(65) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_asset_id` (`asset_id`) USING BTREE,
  KEY `idx_app_id` (`app_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=182 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `campaigns`;
CREATE TABLE `campaigns` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `campaign_id` varchar(35) NOT NULL COMMENT '计划 ID',
  `app_id` varchar(32) NOT NULL DEFAULT '' COMMENT '第三方应用ID，例如华为APP ID : C10134672；可能存在GP的应用ID 32位',
  `campaign_name` varchar(100) NOT NULL DEFAULT '' COMMENT '计划名称',
  `account_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '账户 ID',
  `advertiser_id` varchar(30) NOT NULL DEFAULT '' COMMENT '广告主账户 ID',
  `opt_status` varchar(35) NOT NULL DEFAULT '' COMMENT '操作状态',
  `campaign_daily_budget_status` varchar(35) NOT NULL DEFAULT '' COMMENT '计划日预算状态',
  `product_type` varchar(20) NOT NULL DEFAULT '' COMMENT '推广产品类型',
  `show_status` varchar(60) NOT NULL DEFAULT '' COMMENT '计划状态',
  `user_balance_status` varchar(35) NOT NULL DEFAULT '' COMMENT '账户余额状态\n',
  `flow_resource` varchar(35) NOT NULL DEFAULT '' COMMENT '投放网络',
  `sync_flow_resource` varchar(5) NOT NULL DEFAULT '' COMMENT '同时同步投放搜索广告网络',
  `campaign_type` varchar(35) NOT NULL DEFAULT '' COMMENT '计划类型',
  `today_daily_budget` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '当日计划日限额',
  `tomorrow_daily_budget` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '次日计划日限额，不返回表示与当日计划日限额相同',
  `marketing_goal` varchar(35) NOT NULL DEFAULT '' COMMENT '营销目标',
  `is_callback` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '是否通过查询计划任务回调完整信息',
  `created_at` datetime NOT NULL COMMENT '添加时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_campaign_id` (`campaign_id`) USING BTREE,
  KEY `idx_app_id` (`app_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=560 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `continents`;
CREATE TABLE `continents` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `c_name` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `jobs`;
CREATE TABLE `jobs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `stat_day` date NOT NULL COMMENT '数据日期',
  `api_module` varchar(20) NOT NULL DEFAULT '' COMMENT '数据模块',
  `job_schedule` varchar(20) NOT NULL DEFAULT '' COMMENT 'cron 调度',
  `pause_rule` int(3) NOT NULL DEFAULT '0' COMMENT '调度截止规则：0 调度到当天；-1 停止调度此任务；> 0 为当前日期减{pause_rule}天',
  `version` int(6) unsigned NOT NULL DEFAULT '1' COMMENT '版本：每次有规则或调度修改 +1',
  `remark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注',
  `order_by` tinyint(2) unsigned NOT NULL DEFAULT '99',
  `last_schedule` datetime DEFAULT NULL COMMENT '最后调度完成时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_module` (`api_module`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `oauth_access_tokens`;
CREATE TABLE `oauth_access_tokens` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(15) NOT NULL DEFAULT '0',
  `client_id` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '普通用户的授权，默认为1',
  `token` varchar(2000) NOT NULL,
  `action_name` varchar(128) NOT NULL DEFAULT '' COMMENT 'login|refresh|reset表示token生成动作',
  `scopes` varchar(128) DEFAULT NULL,
  `revoked` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否撤销',
  `client_ip` varchar(128) NOT NULL COMMENT 'ipv6最长为128位',
  `created_at` datetime NOT NULL COMMENT '添加时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次修改时间',
  `expires_at` datetime NOT NULL COMMENT '过期时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `oauth_access_tokens_user_id_index` (`user_id`) USING BTREE,
  KEY `idx_user_id_expires_at` (`user_id`,`expires_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

DROP TABLE IF EXISTS `overseas_area_regions`;
CREATE TABLE `overseas_area_regions` (
  `area_id` int(11) unsigned NOT NULL DEFAULT '0',
  `c_code` varchar(2) NOT NULL DEFAULT '',
  UNIQUE KEY `idx_unique` (`area_id`,`c_code`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `overseas_areas`;
CREATE TABLE `overseas_areas` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL COMMENT '区域名称，例如欧洲区，非洲区，东南亚区等',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `overseas_regions`;
CREATE TABLE `overseas_regions` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `c_id` varchar(32) NOT NULL DEFAULT '',
  `pid` varchar(32) NOT NULL DEFAULT '',
  `c_code` varchar(2) NOT NULL DEFAULT '',
  `c_name` varchar(50) NOT NULL DEFAULT '',
  `level` tinyint(2) unsigned NOT NULL DEFAULT '0',
  `continent_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '大洲ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=58271 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `position_base_prices`;
CREATE TABLE `position_base_prices` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `creative_size_id` varchar(20) NOT NULL DEFAULT '' COMMENT '版位ID',
  `price_type` varchar(30) NOT NULL DEFAULT '' COMMENT '付费方式',
  `base_price` decimal(5,2) NOT NULL DEFAULT '0.00' COMMENT '底价',
  PRIMARY KEY (`id`),
  KEY `idx_creative_size_id` (`creative_size_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=228 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `position_elements`;
CREATE TABLE `position_elements` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `creative_size_id` varchar(20) NOT NULL DEFAULT '' COMMENT '版位ID',
  `sub_type` varchar(50) NOT NULL DEFAULT '' COMMENT '版位子形式',
  `group_number` smallint(5) unsigned NOT NULL DEFAULT '1' COMMENT '版位子形式分组',
  `element_id` varchar(20) NOT NULL DEFAULT '' COMMENT '版位元素id',
  `element_name` varchar(50) NOT NULL DEFAULT '' COMMENT '版位元素类型',
  `element_title` varchar(50) NOT NULL DEFAULT '' COMMENT '版位元素名称',
  `element_caption` varchar(100) NOT NULL DEFAULT '' COMMENT '版位元素描述',
  `width` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '图片宽',
  `height` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '图片高',
  `min_width` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '视频最小宽度',
  `min_height` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '视频最小高度',
  `min_length` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最小输入长度',
  `max_length` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文案、摘要、品牌名称，都是指中文长度，英文长度',
  `file_size_kb_limit` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '文件大小上限，单位KB',
  `gif_size_kb_limit` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'Gif文件大小上限，单位KB',
  `file_format` varchar(20) NOT NULL DEFAULT '' COMMENT '文件类型',
  `pattern` varchar(2000) NOT NULL DEFAULT '' COMMENT '输入校验规则',
  `duration` varchar(1000) NOT NULL DEFAULT '' COMMENT '视频时长',
  `min_occurs` varchar(20) NOT NULL DEFAULT '' COMMENT '最小出现次数，为0表示元素为可选',
  `max_occurs` varchar(20) NOT NULL DEFAULT '' COMMENT '最大出现次数',
  PRIMARY KEY (`id`),
  KEY `idx_creative_size_id` (`creative_size_id`,`sub_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=868 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `position_placements`;
CREATE TABLE `position_placements` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `creative_size_id` varchar(20) NOT NULL DEFAULT '' COMMENT '版位ID',
  `placement_size_id` varchar(50) NOT NULL DEFAULT '' COMMENT '规格ID',
  `creative_size` varchar(30) NOT NULL DEFAULT '' COMMENT '尺寸',
  `creative_size_sub_type` varchar(50) NOT NULL DEFAULT '' COMMENT '版位子形式',
  `is_support_multiple_creatives` varchar(50) NOT NULL DEFAULT '' COMMENT '是否支持多创意',
  PRIMARY KEY (`id`),
  KEY `idx_creative_size_id` (`creative_size_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1509606 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `position_samples`;
CREATE TABLE `position_samples` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `creative_size_id` varchar(20) NOT NULL DEFAULT '' COMMENT '版位ID',
  `creative_size_sample` varchar(2000) NOT NULL DEFAULT '' COMMENT '预览图地址',
  `preview_title` varchar(200) NOT NULL DEFAULT '' COMMENT '预览图标题',
  PRIMARY KEY (`id`),
  KEY `idx_creative_size_id` (`creative_size_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=175595 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `positions`;
CREATE TABLE `positions` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `account_id` bigint(15) unsigned NOT NULL DEFAULT '0',
  `advertiser_id` varchar(30) NOT NULL DEFAULT '' COMMENT '广告主账户ID',
  `creative_size_id` varchar(20) NOT NULL DEFAULT '' COMMENT '版位ID',
  `creative_size_name_dsp` varchar(255) NOT NULL DEFAULT '' COMMENT '版位名称',
  `creative_size_description` varchar(255) NOT NULL DEFAULT '' COMMENT '版位描述',
  `category` varchar(50) NOT NULL DEFAULT '' COMMENT '版位所属分类',
  `support_product_type` varchar(100) NOT NULL DEFAULT '' COMMENT '支持的推广产品',
  `support_objective_type` varchar(200) NOT NULL DEFAULT '',
  `is_support_time_period` varchar(50) NOT NULL DEFAULT '' COMMENT '是否支持选择投放时段',
  `is_support_multiple_creatives` varchar(50) NOT NULL DEFAULT '' COMMENT '是否支持多创意',
  `support_price_type` varchar(100) NOT NULL DEFAULT '' COMMENT '付费方式',
  `last_pull_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后拉取时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_creative_size_id` (`creative_size_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=56867 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `report_ads_collect_accounts`;
CREATE TABLE `report_ads_collect_accounts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `stat_day` date NOT NULL COMMENT '日: 日粒度，例如2021-09-08；',
  `country` varchar(4) NOT NULL DEFAULT '' COMMENT '国家代码，使用华为开发者文档中的广告代码库',
  `account_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '投放类型账户ID',
  `ads_account_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '变现类型账户ID',
  `app_id` varchar(32) NOT NULL DEFAULT '' COMMENT '应用ID（此处一般标识三方应用ID）',
  `earnings` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT '收入',
  `ad_requests` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '到达服务器的请求数量',
  `matched_ad_requests` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '匹配到的到达广告请求数量',
  `show_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '展示数',
  `click_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '点击数',
  PRIMARY KEY (`id`,`stat_day`),
  UNIQUE KEY `unique` (`stat_day`,`app_id`,`account_id`,`country`) USING BTREE,
  KEY `idx_stat_day` (`stat_day`,`account_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=662426 DEFAULT CHARSET=utf8mb4
/*!50100 PARTITION BY RANGE (TO_DAYS(stat_day))
(PARTITION p20231 VALUES LESS THAN (738886) ENGINE = InnoDB,
 PARTITION p20232 VALUES LESS THAN (738917) ENGINE = InnoDB,
 PARTITION p20233 VALUES LESS THAN (738945) ENGINE = InnoDB,
 PARTITION p20234 VALUES LESS THAN (738976) ENGINE = InnoDB,
 PARTITION p20235 VALUES LESS THAN (739006) ENGINE = InnoDB,
 PARTITION p20236 VALUES LESS THAN (739037) ENGINE = InnoDB,
 PARTITION p20237 VALUES LESS THAN (739067) ENGINE = InnoDB,
 PARTITION p20238 VALUES LESS THAN (739098) ENGINE = InnoDB,
 PARTITION p20239 VALUES LESS THAN (739129) ENGINE = InnoDB,
 PARTITION p202310 VALUES LESS THAN (739159) ENGINE = InnoDB,
 PARTITION p202311 VALUES LESS THAN (739190) ENGINE = InnoDB,
 PARTITION p202312 VALUES LESS THAN (739220) ENGINE = InnoDB,
 PARTITION p20241 VALUES LESS THAN (739251) ENGINE = InnoDB,
 PARTITION p20242 VALUES LESS THAN (739282) ENGINE = InnoDB,
 PARTITION p20243 VALUES LESS THAN (739311) ENGINE = InnoDB,
 PARTITION p20244 VALUES LESS THAN (739342) ENGINE = InnoDB,
 PARTITION p20245 VALUES LESS THAN (739372) ENGINE = InnoDB,
 PARTITION p20246 VALUES LESS THAN (739403) ENGINE = InnoDB,
 PARTITION p20247 VALUES LESS THAN (739433) ENGINE = InnoDB,
 PARTITION p20248 VALUES LESS THAN (739464) ENGINE = InnoDB,
 PARTITION p20249 VALUES LESS THAN (739495) ENGINE = InnoDB,
 PARTITION p202410 VALUES LESS THAN (739525) ENGINE = InnoDB,
 PARTITION p202411 VALUES LESS THAN (739556) ENGINE = InnoDB,
 PARTITION p202412 VALUES LESS THAN (739586) ENGINE = InnoDB,
 PARTITION p20251 VALUES LESS THAN (739617) ENGINE = InnoDB) */;

DROP TABLE IF EXISTS `report_ads_collects`;
CREATE TABLE `report_ads_collects` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `stat_day` date NOT NULL COMMENT '日: 日粒度，例如2021-09-08；',
  `country` varchar(4) NOT NULL DEFAULT '' COMMENT '国家代码，使用华为开发者文档中的广告代码库',
  `ads_account_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '变现类型账户ID',
  `app_id` varchar(32) NOT NULL DEFAULT '' COMMENT '应用ID（此处一般标识三方应用ID）',
  `earnings` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT '收入',
  `ad_requests` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '到达服务器的请求数量',
  `matched_ad_requests` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '匹配到的到达广告请求数量',
  `show_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '展示数',
  `click_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '点击数',
  PRIMARY KEY (`id`,`stat_day`),
  UNIQUE KEY `unique` (`stat_day`,`app_id`,`country`,`ads_account_id`) USING BTREE,
  KEY `idx_stat_day` (`stat_day`,`country`,`app_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=559700 DEFAULT CHARSET=utf8mb4
/*!50100 PARTITION BY RANGE (TO_DAYS(stat_day))
(PARTITION p20231 VALUES LESS THAN (738886) ENGINE = InnoDB,
 PARTITION p20232 VALUES LESS THAN (738917) ENGINE = InnoDB,
 PARTITION p20233 VALUES LESS THAN (738945) ENGINE = InnoDB,
 PARTITION p20234 VALUES LESS THAN (738976) ENGINE = InnoDB,
 PARTITION p20235 VALUES LESS THAN (739006) ENGINE = InnoDB,
 PARTITION p20236 VALUES LESS THAN (739037) ENGINE = InnoDB,
 PARTITION p20237 VALUES LESS THAN (739067) ENGINE = InnoDB,
 PARTITION p20238 VALUES LESS THAN (739098) ENGINE = InnoDB,
 PARTITION p20239 VALUES LESS THAN (739129) ENGINE = InnoDB,
 PARTITION p202310 VALUES LESS THAN (739159) ENGINE = InnoDB,
 PARTITION p202311 VALUES LESS THAN (739190) ENGINE = InnoDB,
 PARTITION p202312 VALUES LESS THAN (739220) ENGINE = InnoDB,
 PARTITION p20241 VALUES LESS THAN (739251) ENGINE = InnoDB,
 PARTITION p20242 VALUES LESS THAN (739282) ENGINE = InnoDB,
 PARTITION p20243 VALUES LESS THAN (739311) ENGINE = InnoDB,
 PARTITION p20244 VALUES LESS THAN (739342) ENGINE = InnoDB,
 PARTITION p20245 VALUES LESS THAN (739372) ENGINE = InnoDB,
 PARTITION p20246 VALUES LESS THAN (739403) ENGINE = InnoDB,
 PARTITION p20247 VALUES LESS THAN (739433) ENGINE = InnoDB,
 PARTITION p20248 VALUES LESS THAN (739464) ENGINE = InnoDB,
 PARTITION p20249 VALUES LESS THAN (739495) ENGINE = InnoDB,
 PARTITION p202410 VALUES LESS THAN (739525) ENGINE = InnoDB,
 PARTITION p202411 VALUES LESS THAN (739556) ENGINE = InnoDB,
 PARTITION p202412 VALUES LESS THAN (739586) ENGINE = InnoDB,
 PARTITION p20251 VALUES LESS THAN (739617) ENGINE = InnoDB) */;

DROP TABLE IF EXISTS `report_ads_sources`;
CREATE TABLE `report_ads_sources` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `stat_day` date NOT NULL COMMENT '日: 日粒度，例如2021-09-08；',
  `stat_hour` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '时间: 小时',
  `country` varchar(4) NOT NULL DEFAULT '' COMMENT '国家代码，使用华为开发者文档中的广告代码库',
  `account_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '变现类型账户ID',
  `app_id` varchar(32) NOT NULL DEFAULT '' COMMENT '应用ID（此处一般标识三方应用ID）',
  `ad_type` varchar(12) NOT NULL DEFAULT '' COMMENT ' 广告类型，字符串类型，例如banner，native',
  `placement_id` varchar(30) NOT NULL DEFAULT '' COMMENT '广告位ID',
  `ad_requests` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '到达服务器的请求数量',
  `matched_ad_requests` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '匹配到的到达广告请求数量',
  `show_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '展示数',
  `click_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '点击数',
  `ad_requests_match_rate` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT '填充率',
  `ad_requests_show_rate` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT '请求展示率',
  `click_through_rate` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT '点击率',
  `earnings` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT '收入',
  `ecpm` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT 'ECPM',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`,`stat_day`),
  UNIQUE KEY `unique` (`stat_day`,`account_id`,`app_id`,`country`,`placement_id`,`ad_type`) USING BTREE,
  KEY `idx_stat_day_IDX` (`stat_day`,`app_id`) USING BTREE,
  KEY `idx_app_id_IDX` (`app_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5755244 DEFAULT CHARSET=utf8mb4
/*!50100 PARTITION BY RANGE (TO_DAYS(stat_day))
(PARTITION p20231 VALUES LESS THAN (738886) ENGINE = InnoDB,
 PARTITION p20232 VALUES LESS THAN (738917) ENGINE = InnoDB,
 PARTITION p20233 VALUES LESS THAN (738945) ENGINE = InnoDB,
 PARTITION p20234 VALUES LESS THAN (738976) ENGINE = InnoDB,
 PARTITION p20235 VALUES LESS THAN (739006) ENGINE = InnoDB,
 PARTITION p20236 VALUES LESS THAN (739037) ENGINE = InnoDB,
 PARTITION p20237 VALUES LESS THAN (739067) ENGINE = InnoDB,
 PARTITION p20238 VALUES LESS THAN (739098) ENGINE = InnoDB,
 PARTITION p20239 VALUES LESS THAN (739129) ENGINE = InnoDB,
 PARTITION p202310 VALUES LESS THAN (739159) ENGINE = InnoDB,
 PARTITION p202311 VALUES LESS THAN (739190) ENGINE = InnoDB,
 PARTITION p202312 VALUES LESS THAN (739220) ENGINE = InnoDB,
 PARTITION p20241 VALUES LESS THAN (739251) ENGINE = InnoDB,
 PARTITION p20242 VALUES LESS THAN (739282) ENGINE = InnoDB,
 PARTITION p20243 VALUES LESS THAN (739311) ENGINE = InnoDB,
 PARTITION p20244 VALUES LESS THAN (739342) ENGINE = InnoDB,
 PARTITION p20245 VALUES LESS THAN (739372) ENGINE = InnoDB,
 PARTITION p20246 VALUES LESS THAN (739403) ENGINE = InnoDB,
 PARTITION p20247 VALUES LESS THAN (739433) ENGINE = InnoDB,
 PARTITION p20248 VALUES LESS THAN (739464) ENGINE = InnoDB,
 PARTITION p20249 VALUES LESS THAN (739495) ENGINE = InnoDB,
 PARTITION p202410 VALUES LESS THAN (739525) ENGINE = InnoDB,
 PARTITION p202411 VALUES LESS THAN (739556) ENGINE = InnoDB,
 PARTITION p202412 VALUES LESS THAN (739586) ENGINE = InnoDB,
 PARTITION p20251 VALUES LESS THAN (739617) ENGINE = InnoDB) */;

DROP TABLE IF EXISTS `report_columns`;
CREATE TABLE `report_columns` (
  `column_key` varchar(50) NOT NULL DEFAULT '',
  `columns` varchar(1000) NOT NULL DEFAULT '',
  UNIQUE KEY `idx_column_key` (`column_key`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `report_market_collects`;
CREATE TABLE `report_market_collects` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `stat_day` date NOT NULL COMMENT '日: 日粒度，例如2021-09-08；',
  `country` varchar(4) NOT NULL DEFAULT '' COMMENT '国家代码，使用华为开发者文档中的广告代码库',
  `account_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '应用所属账户ID',
  `app_id` varchar(32) NOT NULL DEFAULT '' COMMENT '应用ID（此处一般标识三方应用ID）',
  `app_name` varchar(128) NOT NULL DEFAULT '' COMMENT '应用名称',
  `cost` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT '花费',
  `show_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '展示数',
  `click_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '点击数',
  `download_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '下载数',
  `install_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '安装数',
  `activate_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '激活数',
  `retain_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '留存数',
  `three_retain_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '三日留存数',
  `seven_retain_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '七日留存数',
  PRIMARY KEY (`id`,`stat_day`),
  UNIQUE KEY `unique` (`stat_day`,`account_id`,`app_id`,`country`) USING BTREE,
  KEY `idx_day_app` (`stat_day`,`app_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=167763 DEFAULT CHARSET=utf8mb4
/*!50100 PARTITION BY RANGE (TO_DAYS(stat_day))
(PARTITION p20231 VALUES LESS THAN (738886) ENGINE = InnoDB,
 PARTITION p20232 VALUES LESS THAN (738917) ENGINE = InnoDB,
 PARTITION p20233 VALUES LESS THAN (738945) ENGINE = InnoDB,
 PARTITION p20234 VALUES LESS THAN (738976) ENGINE = InnoDB,
 PARTITION p20235 VALUES LESS THAN (739006) ENGINE = InnoDB,
 PARTITION p20236 VALUES LESS THAN (739037) ENGINE = InnoDB,
 PARTITION p20237 VALUES LESS THAN (739067) ENGINE = InnoDB,
 PARTITION p20238 VALUES LESS THAN (739098) ENGINE = InnoDB,
 PARTITION p20239 VALUES LESS THAN (739129) ENGINE = InnoDB,
 PARTITION p202310 VALUES LESS THAN (739159) ENGINE = InnoDB,
 PARTITION p202311 VALUES LESS THAN (739190) ENGINE = InnoDB,
 PARTITION p202312 VALUES LESS THAN (739220) ENGINE = InnoDB,
 PARTITION p20241 VALUES LESS THAN (739251) ENGINE = InnoDB,
 PARTITION p20242 VALUES LESS THAN (739282) ENGINE = InnoDB,
 PARTITION p20243 VALUES LESS THAN (739311) ENGINE = InnoDB,
 PARTITION p20244 VALUES LESS THAN (739342) ENGINE = InnoDB,
 PARTITION p20245 VALUES LESS THAN (739372) ENGINE = InnoDB,
 PARTITION p20246 VALUES LESS THAN (739403) ENGINE = InnoDB,
 PARTITION p20247 VALUES LESS THAN (739433) ENGINE = InnoDB,
 PARTITION p20248 VALUES LESS THAN (739464) ENGINE = InnoDB,
 PARTITION p20249 VALUES LESS THAN (739495) ENGINE = InnoDB,
 PARTITION p202410 VALUES LESS THAN (739525) ENGINE = InnoDB,
 PARTITION p202411 VALUES LESS THAN (739556) ENGINE = InnoDB,
 PARTITION p202412 VALUES LESS THAN (739586) ENGINE = InnoDB,
 PARTITION p20251 VALUES LESS THAN (739617) ENGINE = InnoDB) */;

DROP TABLE IF EXISTS `report_market_sources`;
CREATE TABLE `report_market_sources` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `stat_day` date NOT NULL COMMENT '日: 日粒度，例如2021-09-08；',
  `stat_hour` tinyint(2) unsigned NOT NULL DEFAULT '0' COMMENT '时间: 小时',
  `country` varchar(4) NOT NULL DEFAULT '' COMMENT '国家代码，使用华为开发者文档中的广告代码库',
  `account_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '应用所属账户ID',
  `app_id` varchar(32) NOT NULL DEFAULT '' COMMENT '应用ID（此处一般标识三方应用ID）',
  `app_name` varchar(128) NOT NULL DEFAULT '' COMMENT '应用名称',
  `pkg_name` varchar(128) NOT NULL DEFAULT '' COMMENT '投放的应用包名',
  `campaign_id` varchar(20) NOT NULL DEFAULT '' COMMENT '广告计划ID',
  `campaign_name` varchar(200) NOT NULL DEFAULT '' COMMENT '广告计划名称',
  `adgroup_id` varchar(20) NOT NULL DEFAULT '' COMMENT '广告任务ID',
  `adgroup_name` varchar(200) NOT NULL DEFAULT '' COMMENT '广告任务名称',
  `creative_id` varchar(20) NOT NULL DEFAULT '' COMMENT '广告创意ID',
  `creative_name` varchar(200) NOT NULL DEFAULT '' COMMENT '广告创意名称',
  `cost` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT '花费',
  `show_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '展示数',
  `click_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '点击数',
  `download_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '下载数',
  `install_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '安装数',
  `activate_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '激活数',
  `retain_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '留存数',
  `three_retain_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '三日留存数',
  `seven_retain_count` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '七日留存数',
  `click_through_rate` decimal(10,6) NOT NULL DEFAULT '0.000000' COMMENT '点击率',
  `click_download_rate` decimal(10,6) NOT NULL DEFAULT '0.000000' COMMENT '点击下载转化率',
  `download_activate_rate` decimal(10,6) NOT NULL DEFAULT '0.000000' COMMENT '下载激活转化率',
  `cpm` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT 'CPM(千人成本): 花费*1000/展示量',
  `cpc` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT 'CPC: 花费/点击量',
  `cpd` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT 'CPD: 花费/下载数',
  `cpi` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT 'CPI: 花费/安装数',
  `cpa` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT 'CPA: 花费/激活数',
  `seven_retain_cost` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT 'seven_retain_cost: 平均留存花费=花费/七日留存数',
  `retain_cost` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT 'retain_cost: 平均留存花费=花费/留存数',
  `three_retain_cost` decimal(10,5) NOT NULL DEFAULT '0.00000' COMMENT 'three_retain_cost: 平均留存花费=花费/三日留存数',
  `created_at` datetime NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`,`stat_day`),
  UNIQUE KEY `unique` (`stat_day`,`stat_hour`,`country`,`account_id`,`app_id`,`campaign_id`,`adgroup_id`,`creative_id`) USING BTREE,
  KEY `idx_day` (`stat_day`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=665122 DEFAULT CHARSET=utf8mb4
/*!50100 PARTITION BY RANGE (TO_DAYS(stat_day))
(PARTITION p20231 VALUES LESS THAN (738886) ENGINE = InnoDB,
 PARTITION p20232 VALUES LESS THAN (738917) ENGINE = InnoDB,
 PARTITION p20233 VALUES LESS THAN (738945) ENGINE = InnoDB,
 PARTITION p20234 VALUES LESS THAN (738976) ENGINE = InnoDB,
 PARTITION p20235 VALUES LESS THAN (739006) ENGINE = InnoDB,
 PARTITION p20236 VALUES LESS THAN (739037) ENGINE = InnoDB,
 PARTITION p20237 VALUES LESS THAN (739067) ENGINE = InnoDB,
 PARTITION p20238 VALUES LESS THAN (739098) ENGINE = InnoDB,
 PARTITION p20239 VALUES LESS THAN (739129) ENGINE = InnoDB,
 PARTITION p202310 VALUES LESS THAN (739159) ENGINE = InnoDB,
 PARTITION p202311 VALUES LESS THAN (739190) ENGINE = InnoDB,
 PARTITION p202312 VALUES LESS THAN (739220) ENGINE = InnoDB,
 PARTITION p20241 VALUES LESS THAN (739251) ENGINE = InnoDB,
 PARTITION p20242 VALUES LESS THAN (739282) ENGINE = InnoDB,
 PARTITION p20243 VALUES LESS THAN (739311) ENGINE = InnoDB,
 PARTITION p20244 VALUES LESS THAN (739342) ENGINE = InnoDB,
 PARTITION p20245 VALUES LESS THAN (739372) ENGINE = InnoDB,
 PARTITION p20246 VALUES LESS THAN (739403) ENGINE = InnoDB,
 PARTITION p20247 VALUES LESS THAN (739433) ENGINE = InnoDB,
 PARTITION p20248 VALUES LESS THAN (739464) ENGINE = InnoDB,
 PARTITION p20249 VALUES LESS THAN (739495) ENGINE = InnoDB,
 PARTITION p202410 VALUES LESS THAN (739525) ENGINE = InnoDB,
 PARTITION p202411 VALUES LESS THAN (739556) ENGINE = InnoDB,
 PARTITION p202412 VALUES LESS THAN (739586) ENGINE = InnoDB,
 PARTITION p20251 VALUES LESS THAN (739617) ENGINE = InnoDB) */;

DROP TABLE IF EXISTS `role_permission_rules`;
CREATE TABLE `role_permission_rules` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_role_permission_rules` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=177 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='权限';

DROP TABLE IF EXISTS `role_permissions`;
CREATE TABLE `role_permissions` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `permission` varchar(100) NOT NULL DEFAULT '' COMMENT '路由URL，示例：/api/login，/api/user/*',
  `p_name` varchar(50) NOT NULL DEFAULT '' COMMENT '路由名称',
  `method` varchar(10) NOT NULL DEFAULT '' COMMENT '运行请求方式：GET/POST/PUT/DELETE/*',
  `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '路由分组，上级ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `idx_permission` (`permission`,`method`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `role_name` varchar(50) NOT NULL DEFAULT '',
  `state` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '1正常0停用',
  `sys` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '角色所属系统',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC;

DROP TABLE IF EXISTS `sys_configs`;
CREATE TABLE `sys_configs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `_k` varchar(50) NOT NULL DEFAULT '',
  `_v` varchar(500) NOT NULL DEFAULT '',
  `_desc` varchar(50) NOT NULL DEFAULT '',
  `state` tinyint(1) unsigned NOT NULL DEFAULT '1',
  `bak1` varchar(50) NOT NULL DEFAULT '',
  `bak2` varchar(1000) NOT NULL DEFAULT '',
  `remark` varchar(100) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `sys_logs`;
CREATE TABLE `sys_logs` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `stat_day` date NOT NULL,
  `msg` varchar(100) NOT NULL DEFAULT '',
  `module` varchar(100) NOT NULL DEFAULT '',
  `info` text,
  `level` varchar(10) NOT NULL DEFAULT '',
  `log_id` varchar(30) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique` (`module`,`level`,`log_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `targeting_dictionaries`;
CREATE TABLE `targeting_dictionaries` (
  `dict_key` varchar(32) NOT NULL DEFAULT '' COMMENT '属于什么字典',
  `id` varchar(32) NOT NULL DEFAULT '' COMMENT '字典的元素 ID',
  `pid` varchar(32) NOT NULL DEFAULT '' COMMENT '父节点元素ID',
  `label` varchar(200) NOT NULL DEFAULT '' COMMENT '显示的内容',
  `value` varchar(200) NOT NULL DEFAULT '' COMMENT '元素的值',
  `code` varchar(500) NOT NULL DEFAULT '',
  `seq` varchar(32) NOT NULL DEFAULT '',
  `data_struct` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '数据结构：1 line，2 tree',
  KEY `idx_dict_key` (`dict_key`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `targetings`;
CREATE TABLE `targetings` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `account_id` bigint(15) unsigned NOT NULL DEFAULT '0',
  `advertiser_id` varchar(30) NOT NULL DEFAULT '' COMMENT '广告主账户ID',
  `targeting_id` bigint(15) unsigned NOT NULL DEFAULT '0',
  `targeting_name` varchar(100) NOT NULL DEFAULT '' COMMENT '定向名称',
  `targeting_type` varchar(32) NOT NULL DEFAULT '' COMMENT '定向类型',
  `location_type` varchar(15) NOT NULL DEFAULT '' COMMENT '地域定向类型',
  `include_location` varchar(2000) NOT NULL COMMENT '地域 - 包含',
  `exclude_location` varchar(2000) NOT NULL COMMENT '地域 - 排除',
  `carriers` varchar(2000) NOT NULL DEFAULT '' COMMENT '运营商',
  `language` varchar(1000) NOT NULL COMMENT '语言',
  `age` varchar(50) NOT NULL COMMENT '年龄',
  `gender` varchar(30) NOT NULL DEFAULT '' COMMENT '性别',
  `app_category` varchar(5) NOT NULL DEFAULT '' COMMENT 'App 行为类型',
  `app_categories` varchar(1000) NOT NULL DEFAULT '' COMMENT 'App 行为',
  `installed_apps` varchar(5) NOT NULL DEFAULT '' COMMENT 'app 安装',
  `app_interest` varchar(5) NOT NULL DEFAULT '' COMMENT 'App 兴趣类型',
  `app_interests` varchar(1000) NOT NULL DEFAULT '' COMMENT 'App 兴趣',
  `series` varchar(1000) NOT NULL DEFAULT '' COMMENT '设备',
  `network_type` varchar(100) NOT NULL DEFAULT '' COMMENT '联网方式',
  `not_audiences` varchar(2000) NOT NULL DEFAULT '' COMMENT '排除人群',
  `audiences` varchar(2000) NOT NULL DEFAULT '' COMMENT '包含人群',
  `app_category_of_media` varchar(1000) NOT NULL DEFAULT '' COMMENT '媒体类型',
  `created_at` datetime NOT NULL COMMENT '添加时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_targeting_name` (`targeting_name`) USING BTREE,
  KEY `idx_targeting_id` (`targeting_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4246 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `trackings`;
CREATE TABLE `trackings` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `account_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '账户ID;对应 accounts 表的id字段',
  `advertiser_id` varchar(30) NOT NULL DEFAULT '' COMMENT '广告主账户ID',
  `app_id` varchar(32) NOT NULL DEFAULT '' COMMENT '第三方应用ID，例如华为APP ID : C10134672；可能存在GP的应用ID 32位',
  `tracking_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '转化跟踪指标ID',
  `effect_type` varchar(30) NOT NULL DEFAULT '' COMMENT '转化目标',
  `effect_name` varchar(100) NOT NULL DEFAULT '' COMMENT '转化跟踪指标名称',
  `click_tracking_url` varchar(1000) NOT NULL DEFAULT '' COMMENT '点击监测地址',
  `imp_tracking_url` varchar(1000) NOT NULL DEFAULT '' COMMENT '曝光监测地址',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_tracking_id` (`app_id`,`tracking_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `user_accounts`;
CREATE TABLE `user_accounts` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `account_id` bigint(20) unsigned NOT NULL DEFAULT '0',
  `account_type` tinyint(2) unsigned NOT NULL DEFAULT '1' COMMENT '账户类型',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique` (`user_id`,`account_id`,`account_type`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `sso_uid` varchar(50) NOT NULL DEFAULT '',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '登陆邮箱',
  `username` varchar(30) NOT NULL DEFAULT '' COMMENT '用户名',
  `mobile` varchar(15) NOT NULL DEFAULT '' COMMENT '手机号',
  `state` tinyint(2) NOT NULL DEFAULT '1' COMMENT '账号状态',
  `secret` varchar(15) NOT NULL DEFAULT '' COMMENT '密码加密符',
  `pass` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  `created_at` datetime NOT NULL COMMENT '添加时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '最后一次修改时间',
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色ID',
  `is_internal` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否内部账号',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_email` (`email`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4;



/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;