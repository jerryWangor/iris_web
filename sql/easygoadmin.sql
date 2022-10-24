/*
 Navicat MySQL Data Transfer

 Source Server         : 本机
 Source Server Type    : MySQL
 Source Server Version : 50738
 Source Host           : localhost:3306
 Source Schema         : easygoadmin

 Target Server Type    : MySQL
 Target Server Version : 50738
 File Encoding         : 65001

 Date: 24/10/2022 19:52:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

CREATE DATABASE `easygoadmin` /*!40100 DEFAULT CHARACTER SET utf8mb4 */

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL DEFAULT 0 COMMENT '上级ID',
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '菜单名',
  `icon` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '菜单图标',
  `url` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '菜单地址',
  `component` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单组件',
  `target` tinyint(1) NOT NULL DEFAULT 1 COMMENT '打开方式：0组件 1内链 2外链',
  `permission` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '权限标识',
  `type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '类型：0菜单 1节点',
  `method` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT 'Method\r\n请求方式',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态 1正常 2禁用',
  `hide` tinyint(1) NULL DEFAULT 1 COMMENT '是否可见',
  `note` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '备注',
  `sort` int(11) NULL DEFAULT 0 COMMENT '排序',
  `create_user` int(11) NOT NULL DEFAULT 0 COMMENT '创建人',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` int(11) NOT NULL COMMENT '更新人',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `mark` int(11) UNSIGNED NULL DEFAULT 1 COMMENT '标记',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `url`(`url`, `name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, 0, '系统管理', '', '', NULL, 1, '', 0, '', 1, 1, '', 0, 1, NULL, 1, NULL, 1);
INSERT INTO `sys_menu` VALUES (2, 1, '权限管理', '', '', NULL, 1, '', 0, '', 1, 1, '', 0, 1, NULL, 1, NULL, 1);
INSERT INTO `sys_menu` VALUES (3, 2, '用户管理', '', 'user/index', NULL, 1, '', 0, '', 1, 1, '', 0, 1, NULL, 1, NULL, 1);
INSERT INTO `sys_menu` VALUES (4, 2, '角色管理', '', '/role/index', NULL, 1, '', 0, '', 1, 1, '', 0, 1, NULL, 1, NULL, 1);
INSERT INTO `sys_menu` VALUES (5, 2, '菜单管理', '', '/menu/index', NULL, 1, '', 0, '', 1, 1, '', 0, 1, NULL, 1, NULL, 1);
INSERT INTO `sys_menu` VALUES (6, 3, '用户列表', '', 'user/list', NULL, 1, '', 1, '', 1, 1, '', 0, 1, NULL, 1, NULL, 1);
INSERT INTO `sys_menu` VALUES (7, 3, '添加用户', '', '/user/add', NULL, 1, '', 1, '', 1, 1, '', 0, 1, NULL, 1, NULL, 1);
INSERT INTO `sys_menu` VALUES (8, 3, '编辑用户', '', '/user/edit', NULL, 1, '', 1, '', 1, 1, '', 0, 1, NULL, 1, NULL, 1);
INSERT INTO `sys_menu` VALUES (9, 3, '删除用户', '', '/user/delete', NULL, 1, '', 1, '', 1, 1, '', 0, 1, NULL, 1, NULL, 1);

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '角色名',
  `code` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '角色编码',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态 1正常 2禁用',
  `note` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '备注',
  `sort` int(11) NULL DEFAULT 0 COMMENT '排序',
  `createuser` int(11) NOT NULL DEFAULT 0 COMMENT '创建人',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updateuser` int(11) NOT NULL DEFAULT 0 COMMENT '更新人',
  `updatetime` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
  `mark` int(11) NULL DEFAULT 0 COMMENT '标记',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '超级管理员', '', 1, '', 0, 0, 0, 0, 0, 1);

-- ----------------------------
-- Table structure for sys_rolemenu
-- ----------------------------
DROP TABLE IF EXISTS `sys_rolemenu`;
CREATE TABLE `sys_rolemenu`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `roleid` int(11) NOT NULL DEFAULT 0 COMMENT '角色ID',
  `menuids` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '菜单IDs',
  `createuser` int(11) NOT NULL DEFAULT 0 COMMENT '创建人',
  `createtime` int(11) NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updateuser` int(11) NOT NULL DEFAULT 0 COMMENT '更新人',
  `updatetime` int(11) NOT NULL DEFAULT 0 COMMENT '更新时间',
  `mark` int(11) NULL DEFAULT 0 COMMENT '标记',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `roleid`(`roleid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_rolemenu
-- ----------------------------

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `realname` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '真实姓名',
  `nickname` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '昵称',
  `gender` tinyint(1) NULL DEFAULT 3 COMMENT '性别:1男 2女 3保密',
  `avatar` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '头像',
  `mobile` char(11) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '手机号码',
  `email` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '邮箱地址',
  `birthday` date NULL DEFAULT NULL COMMENT '出生日期',
  `dept_id` int(11) NULL DEFAULT 0 COMMENT '部门ID',
  `level_id` int(11) NULL DEFAULT 0 COMMENT '职级ID',
  `position_id` smallint(3) NULL DEFAULT 0 COMMENT '岗位ID',
  `province_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '省份编号',
  `city_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '市区编号',
  `district_code` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '区县编号',
  `address` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '详细地址',
  `city_name` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '所属城市',
  `username` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '登录用户名',
  `password` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '登录密码',
  `salt` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '盐加密',
  `intro` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '个人简介',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '状态：1正常 2禁用',
  `note` varchar(500) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '备注',
  `sort` int(11) NULL DEFAULT 125 COMMENT '排序号',
  `login_num` int(11) NULL DEFAULT 0 COMMENT '登录次数',
  `login_ip` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'NULL' COMMENT '最近登录IP',
  `login_time` datetime NULL DEFAULT NULL COMMENT '最近登录时间',
  `create_user` int(10) NULL DEFAULT 0 COMMENT '添加人',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_user` int(10) NULL DEFAULT 0 COMMENT '更新人',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  `mark` tinyint(1) NOT NULL DEFAULT 1 COMMENT '有效标识(1正常 0删除)',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', 'admin', 3, 'NULL', 'NULL', 'NULL', NULL, 0, 0, 0, 'NULL', 'NULL', 'NULL', 'NULL', 'NULL', 'admin', '43286a86708820e38c333cdd4c496355', 'NULL', 'NULL', 1, 'NULL', 125, 0, 'NULL', NULL, 0, NULL, 0, NULL, 1);
INSERT INTO `sys_user` VALUES (3, '汪进', 'Jerry', 3, 'NULL', 'NULL', 'NULL', NULL, 0, 0, 0, 'NULL', 'NULL', 'NULL', 'NULL', 'NULL', 'wangjin', 'bf09946f22243c26890bbf0f60f7771e', 'NULL', 'NULL', 1, 'NULL', 125, 0, 'NULL', NULL, 0, NULL, 0, NULL, 1);

SET FOREIGN_KEY_CHECKS = 1;
