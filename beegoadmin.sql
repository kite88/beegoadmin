/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50553
Source Host           : 127.0.0.1:3306
Source Database       : beegoadmin

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2020-03-06 18:38:24
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `go_admin`
-- ----------------------------
DROP TABLE IF EXISTS `go_admin`;
CREATE TABLE `go_admin` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(50) DEFAULT NULL COMMENT '用户名',
  `password` varchar(200) DEFAULT NULL COMMENT '密码',
  `role_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of go_admin
-- ----------------------------
INSERT INTO `go_admin` VALUES ('1', 'hejunsheng', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('2', 'cms0', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('3', 'cms1', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('4', 'cms2', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('5', 'cms3', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('6', 'cms4', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('7', 'cms5', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('8', 'cms6', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('9', 'cms7', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('10', 'cms8', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('11', 'cms9', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('12', 'cms0', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('13', 'cms1', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('14', 'cms2', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('15', 'cms3', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('16', 'cms4', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('17', 'cms5', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('18', 'cms6', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('19', 'cms7', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('20', 'cms8', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('21', 'cms9', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('22', 'cms0', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('23', 'cms1', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('24', 'cms2', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('25', 'cms3', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('26', 'cms4', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('27', 'cms5', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('28', 'cms6', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('29', 'cms7', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('30', 'cms8', 'e10adc3949ba59abbe56e057f20f883e', null);
INSERT INTO `go_admin` VALUES ('33', 'hjs002002', '202cb962ac59075b964b07152d234b70', '12');
INSERT INTO `go_admin` VALUES ('32', 'hjs001', '93279e3308bdbbeed946fc965017f67a', null);
INSERT INTO `go_admin` VALUES ('34', '2', 'c81e728d9d4c2f636f067f89cc14862c', '9');
INSERT INTO `go_admin` VALUES ('35', '555', '202cb962ac59075b964b07152d234b70', '7');
INSERT INTO `go_admin` VALUES ('36', '何俊盛', '4297f44b13955235245b2497399d7a93', '1');
INSERT INTO `go_admin` VALUES ('37', 'hejunsheng00252', 'e10adc3949ba59abbe56e057f20f883e', '8');
INSERT INTO `go_admin` VALUES ('38', '小卡', '4297f44b13955235245b2497399d7a93', '11');
INSERT INTO `go_admin` VALUES ('39', '何俊胜', '4297f44b13955235245b2497399d7a93', '15');

-- ----------------------------
-- Table structure for `go_nav`
-- ----------------------------
DROP TABLE IF EXISTS `go_nav`;
CREATE TABLE `go_nav` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(50) DEFAULT NULL COMMENT '名称',
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父ID',
  `router` varchar(200) DEFAULT NULL COMMENT '路由',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4 COMMENT='导航表';

-- ----------------------------
-- Records of go_nav
-- ----------------------------
INSERT INTO `go_nav` VALUES ('1', '平台管理', '0', '');
INSERT INTO `go_nav` VALUES ('2', '系统管理', '0', null);
INSERT INTO `go_nav` VALUES ('3', '安全管理', '0', null);
INSERT INTO `go_nav` VALUES ('4', '数据中心', '0', null);
INSERT INTO `go_nav` VALUES ('5', '表单管理', '1', '');
INSERT INTO `go_nav` VALUES ('6', '流程管理', '1', '');
INSERT INTO `go_nav` VALUES ('7', '交流中心', '6', '');
INSERT INTO `go_nav` VALUES ('8', '管理员', '5', '/cms/admin/index');
INSERT INTO `go_nav` VALUES ('9', '角色', '5', '/cms/role/index');
INSERT INTO `go_nav` VALUES ('10', '开发配置', '0', null);
INSERT INTO `go_nav` VALUES ('11', '配置导航', '10', null);
INSERT INTO `go_nav` VALUES ('12', '配置菜单', '11', '/cms/menu/list');
INSERT INTO `go_nav` VALUES ('19', '服务器信息', '18', '/cms/home/sys');
INSERT INTO `go_nav` VALUES ('18', '系统信息', '2', '');

-- ----------------------------
-- Table structure for `go_role`
-- ----------------------------
DROP TABLE IF EXISTS `go_role`;
CREATE TABLE `go_role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(50) DEFAULT NULL COMMENT '角色名称',
  `is_root` tinyint(1) DEFAULT '0' COMMENT '是否超级开发者 1是 0否',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- ----------------------------
-- Records of go_role
-- ----------------------------
INSERT INTO `go_role` VALUES ('1', '总经理', '1');
INSERT INTO `go_role` VALUES ('2', '总经理', '1');
INSERT INTO `go_role` VALUES ('3', '行政主管', '0');
INSERT INTO `go_role` VALUES ('5', '行政主管', '0');
INSERT INTO `go_role` VALUES ('6', '总裁', '0');
INSERT INTO `go_role` VALUES ('7', '总统', '0');
INSERT INTO `go_role` VALUES ('8', '主席', '0');
INSERT INTO `go_role` VALUES ('9', '技术总监', '1');
INSERT INTO `go_role` VALUES ('10', '总理', '0');
INSERT INTO `go_role` VALUES ('11', '外交部发言人', '0');
INSERT INTO `go_role` VALUES ('12', '行政总裁', '0');
INSERT INTO `go_role` VALUES ('13', '技术总裁', '0');
INSERT INTO `go_role` VALUES ('14', '财务司', '0');
INSERT INTO `go_role` VALUES ('15', '西亚卡姆', '0');
