/*
 Navicat Premium Data Transfer

 Target Server Type    : MySQL
 Target Server Version : 50734
 File Encoding         : 65001

 Date: 29/11/2022 18:23:15
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for rustdesk_peers
-- ----------------------------
DROP TABLE IF EXISTS `rustdesk_peers`;
CREATE TABLE `rustdesk_peers` (
  `deviceid` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `uid` int(10) unsigned NOT NULL COMMENT '用户ID',
  `id` char(16) NOT NULL DEFAULT '' COMMENT '设备ID',
  `username` varchar(128) DEFAULT NULL COMMENT '操作系统用户名',
  `hostname` varchar(128) DEFAULT NULL COMMENT '操作系统名',
  `alias` char(20) DEFAULT NULL COMMENT '别名',
  `platform` char(20) DEFAULT NULL COMMENT '平台',
  `tags` varchar(256) DEFAULT NULL COMMENT '标签',
  PRIMARY KEY (`deviceid`),
  UNIQUE KEY `uuid` (`id`,`uid`)
) ENGINE=MyISAM AUTO_INCREMENT=61 DEFAULT CHARSET=utf8 COMMENT='远程设备表';

-- ----------------------------
-- Records of rustdesk_peers
-- ----------------------------

-- ----------------------------
-- Table structure for rustdesk_tags
-- ----------------------------
DROP TABLE IF EXISTS `rustdesk_tags`;
CREATE TABLE `rustdesk_tags` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT COMMENT 'tagID',
  `uid` int(10) unsigned NOT NULL COMMENT '用户ID',
  `tag` char(20) NOT NULL DEFAULT '' COMMENT 'tag名称',
  `color` char(10) NULL COMMENT 'Tag Color',
  PRIMARY KEY (`id`),
  UNIQUE KEY `tag` (`tag`,`uid`)
) ENGINE=MyISAM AUTO_INCREMENT=136 DEFAULT CHARSET=utf8 COMMENT='tags表';

-- ----------------------------
-- Records of rustdesk_tags
-- ----------------------------

-- ----------------------------
-- Table structure for rustdesk_token
-- ----------------------------
DROP TABLE IF EXISTS `rustdesk_token`;
CREATE TABLE `rustdesk_token` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` char(16) NOT NULL COMMENT '用户名',
  `uid` int(10) unsigned NOT NULL COMMENT '用户ID',
  `client_id` char(16) NOT NULL COMMENT '设备码',
  `uuid` char(64) NOT NULL COMMENT '设备ID',
  `access_token` varchar(128) NOT NULL DEFAULT '' COMMENT '登录token',
  `login_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '登录时间',
  `expire_time` int(11) DEFAULT NULL COMMENT '过期时间',
  `active_time` int(10) DEFAULT NULL COMMENT '最后一次活动时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `login_token` (`uid`,`client_id`,`uuid`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=6 DEFAULT CHARSET=utf8 COMMENT='登录Token表';

-- ----------------------------
-- Records of rustdesk_token
-- ----------------------------

-- ----------------------------
-- Table structure for rustdesk_users
-- ----------------------------
DROP TABLE IF EXISTS `rustdesk_users`;
CREATE TABLE `rustdesk_users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(16) NOT NULL COMMENT '用户名',
  `password` char(32) NOT NULL COMMENT '密码',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '用户状态',
  `last_login_ip` varchar(255) NOT NULL COMMENT '最后一次登录IP',
  `last_login_time` int(10) NOT NULL DEFAULT '0' COMMENT '最后一次登录时间',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `update_time` int(10) NOT NULL DEFAULT '0' COMMENT '修改时间',
  `delete_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- Records of rustdesk_users
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
