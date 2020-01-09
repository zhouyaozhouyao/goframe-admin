/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50727
 Source Host           : localhost
 Source Database       : goframe

 Target Server Type    : MySQL
 Target Server Version : 50727
 File Encoding         : utf-8

 Date: 01/09/2020 17:15:23 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `users`
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `uuid` varchar(32) DEFAULT NULL COMMENT 'UUID',
  `username` varchar(32) NOT NULL COMMENT '登录名/11111',
  `password` varchar(32) NOT NULL COMMENT '密码',
  `salt` varchar(16) NOT NULL DEFAULT '1111' COMMENT '密码盐',
  `real_name` varchar(32) DEFAULT NULL COMMENT '真实姓名',
  `depart_id` int(11) DEFAULT '0' COMMENT '部门/11111/dict',
  `user_type` int(11) DEFAULT '2' COMMENT '类型//select/1,管理员,2,普通用户,3,前台用户,4,第三方用户,5,API用户',
  `status` int(11) DEFAULT '10' COMMENT '状态',
  `thirdid` varchar(200) DEFAULT NULL COMMENT '第三方ID',
  `endtime` varchar(32) DEFAULT NULL COMMENT '结束时间',
  `email` varchar(64) DEFAULT NULL COMMENT 'email',
  `tel` varchar(32) DEFAULT NULL COMMENT '手机号',
  `address` varchar(32) DEFAULT NULL COMMENT '地址',
  `title_url` varchar(200) DEFAULT NULL COMMENT '头像地址',
  `remark` varchar(1000) DEFAULT NULL COMMENT '说明',
  `theme` varchar(64) DEFAULT 'default' COMMENT '主题',
  `back_site_id` int(11) DEFAULT '0' COMMENT '后台选择站点ID',
  `create_site_id` int(11) DEFAULT '1' COMMENT '创建站点ID',
  `project_id` bigint(20) DEFAULT '0' COMMENT '项目ID',
  `project_name` varchar(100) DEFAULT NULL COMMENT '项目名称',
  `enable` tinyint(1) DEFAULT '1' COMMENT '是否启用//radio/1,启用,2,禁用',
  `update_time` varchar(24) DEFAULT NULL COMMENT '更新时间',
  `update_id` bigint(20) DEFAULT '0' COMMENT '更新人',
  `create_time` varchar(24) DEFAULT NULL COMMENT '创建时间',
  `create_id` bigint(20) DEFAULT '0' COMMENT '创建者',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `uni_user_username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT COMMENT='用户';

-- ----------------------------
--  Records of `users`
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES ('1', '94091b1fa6ac4a27a06c0b92155aea6a', 'admin', '9ddf3f8018feb07382504742b503781a', '1111', '系统管理员', '10001', '1', '10', '', '', 'zcool321@sina.com', '123', '', '', '时间是最好的老师，但遗憾的是&mdash;&mdash;最后他把所有的学生都弄死了', 'flat', '5', '1', '1', 'test', '1', '2019-07-08 18:12:28', '1', '2017-03-19 20:41:25', '1'), ('9', 'xa5450ztN08S37tKj93ujhJ66069q92R', 'test', 'ea8207ee50ccf367e99c8444fda7da32', 'GM26Mq', 'test', '10002', '2', '10', null, null, null, null, null, null, null, 'default', '0', '1', '0', null, '1', '2019-11-12 15:31:31', '1', '2019-07-11 15:49:24', '1'), ('12', '8609WdcTI1337Y7e5kQ94v872Z02mh24', 'testLogin', '7f4d0d8db5546f395e87dfd294608b9b', '3n7Ci8', 'testLogin', '10002', '2', '10', null, null, null, null, null, null, null, 'default', '0', '1', '0', null, '1', '2019-11-12 15:31:08', '1', '2019-11-12 15:31:08', '1'), ('13', 'PTMB2mcqk87n1x15K84E56T75SY11Q5w', 'testLogout', '961c0645f7ae271d6e1fc1ff01e786d7', '0X6509', 'testLogout', '10002', '2', '10', null, null, null, null, null, null, null, 'default', '0', '1', '0', null, '1', '2019-11-12 15:31:19', '1', '2019-11-12 15:31:19', '1');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
