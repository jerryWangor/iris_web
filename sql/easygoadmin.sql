/*
SQLyog Enterprise - MySQL GUI v7.12 
MySQL - 5.7.34-log : Database - easygoadmin
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

CREATE DATABASE /*!32312 IF NOT EXISTS*/`easygoadmin` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `easygoadmin`;

/*Table structure for table `sys_menu` */

DROP TABLE IF EXISTS `sys_menu`;

CREATE TABLE `sys_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '上级ID',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '菜单名',
  `icon` varchar(100) DEFAULT '' COMMENT '菜单图标',
  `url` varchar(100) DEFAULT '' COMMENT '菜单地址',
  `component` varchar(150) DEFAULT NULL COMMENT '菜单组件',
  `target` tinyint(1) NOT NULL DEFAULT '1' COMMENT '打开方式：0组件 1内链 2外链',
  `permission` varchar(100) DEFAULT '' COMMENT '权限标识',
  `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '类型：0菜单 1节点',
  `method` varchar(100) DEFAULT '' COMMENT '请求方式',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1正常 2禁用',
  `hide` tinyint(1) DEFAULT '1' COMMENT '是否可见',
  `note` varchar(100) DEFAULT '' COMMENT '备注',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `create_user` int(11) NOT NULL DEFAULT '0' COMMENT '创建人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_user` int(11) NOT NULL COMMENT '更新人',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `mark` int(11) unsigned DEFAULT '1' COMMENT '标记',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `url` (`url`,`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

/*Data for the table `sys_menu` */

insert  into `sys_menu`(`id`,`pid`,`name`,`icon`,`url`,`component`,`target`,`permission`,`type`,`method`,`status`,`hide`,`note`,`sort`,`create_user`,`create_time`,`update_user`,`update_time`,`mark`) values (1,0,'系统管理','','',NULL,1,'',0,'',1,1,'',1,1,'2022-10-27 11:54:45',1,'2022-10-27 11:54:45',1),(2,1,'权限管理','','',NULL,1,'',0,'',1,1,'',1,1,'2022-10-27 11:54:45',1,'2022-10-27 11:54:45',1),(3,2,'用户管理','','/user/index',NULL,1,'',0,'',1,1,'',1,1,'2022-10-27 11:54:45',1,'2022-10-27 14:53:59',1),(4,2,'角色管理','','/role/index',NULL,1,'',0,'',1,1,'',3,1,'2022-10-27 11:54:45',1,'2022-10-27 11:54:45',1),(5,2,'菜单管理','','/menu/index',NULL,1,'',0,'',1,1,'',2,1,'2022-10-27 11:54:45',1,'2022-10-27 11:54:45',1),(6,3,'用户列表','','/user/list',NULL,1,'sys:user:list',1,'',1,1,'',2,1,'2022-10-27 11:54:45',1,'2022-10-27 14:56:06',1),(7,3,'用户添加','','/user/add',NULL,1,'sys:user:add',1,'',1,1,'',3,1,'2022-10-27 11:54:45',1,'2022-10-27 14:56:15',1),(8,3,'用户修改操作','','/user/update',NULL,1,'sys:user:update',1,'',1,1,'',5,1,'2022-10-27 11:54:45',1,'2022-10-27 15:48:01',1),(9,3,'用户删除','','/user/delete',NULL,1,'sys:user:delete',1,'',1,1,'',6,1,'2022-10-27 11:54:45',1,'2022-10-27 14:58:33',1),(10,3,'设置状态','','/user/setStatus','',1,'sys:user:setStatus',1,'',1,0,'',7,1,'2022-10-27 15:01:03',1,'2022-10-27 15:01:03',1),(11,3,'重置密码','','/user/resetPwd','',1,'sys:user/resetPwd',1,'',1,0,'',8,1,'2022-10-27 15:02:05',1,'2022-10-27 15:02:05',1),(12,3,'用户修改页面','','/user/edit','',1,'sys:user:edit',1,'',1,0,'',4,1,'2022-10-27 15:49:48',1,'2022-10-27 15:49:48',1),(15,5,'菜单列表','','/menu/list','',1,'sys:menu:list',1,'',1,0,'',2,1,'2022-10-27 15:54:13',1,'2022-10-27 15:54:13',1),(16,5,'菜单修改页面','','/menu/edit','',1,'sys:menu:edit',1,'',1,0,'',4,1,'2022-10-27 15:55:02',1,'2022-10-27 15:55:02',1),(17,5,'菜单添加','','/menu/add','',1,'sys:menu:add',1,'',1,0,'',3,1,'2022-10-27 15:55:53',1,'2022-10-27 15:55:53',1),(18,5,'菜单修改操作','','menu/update','',1,'sys:menu:update',1,'',1,0,'',5,1,'2022-10-27 15:56:24',1,'2022-10-27 15:56:41',1),(19,5,'菜单删除','','/menu/delete','',1,'sys:menu:delete',1,'',1,0,'',6,1,'2022-10-27 15:57:08',1,'2022-10-27 15:57:08',1),(21,4,'角色列表','','/role/list','',1,'sys:role:list',1,'',1,0,'',2,1,'2022-10-27 15:54:13',1,'2022-10-27 15:54:13',1),(22,4,'角色修改页面','','/role/edit','',1,'sys:role:edit',1,'',1,0,'',4,1,'2022-10-27 15:55:02',1,'2022-10-27 15:55:02',1),(23,4,'角色添加','','/role/add','',1,'sys:role:add',1,'',1,0,'',3,1,'2022-10-27 15:55:53',1,'2022-10-27 15:55:53',1),(24,4,'角色修改操作','','role/update','',1,'sys:role:update',1,'',1,0,'',5,1,'2022-10-27 15:56:24',1,'2022-10-27 15:56:41',1),(25,4,'角色删除','','/role/delete','',1,'sys:role:delete',1,'',1,0,'',6,1,'2022-10-27 15:57:08',1,'2022-10-27 15:57:08',1),(26,4,'设置状态','','/role/setStatus','',1,'sys:role:setStatus',1,'',1,0,'',7,1,'2022-10-27 16:02:03',1,'2022-10-27 16:02:03',1),(27,4,'角色权限列表','','/rolemenu/index','',1,'sys:rolemenu:index',1,'',1,0,'',8,1,'2022-10-27 16:06:11',1,'2022-10-27 16:06:11',1),(28,4,'角色权限设置','','/rolemenu/save','',1,'sys:rolemenu:save',1,'',1,0,'',9,1,'2022-10-27 16:06:36',1,'2022-10-27 16:06:36',1);

/*Table structure for table `sys_role` */

DROP TABLE IF EXISTS `sys_role`;

CREATE TABLE `sys_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '角色名',
  `code` varchar(100) DEFAULT '' COMMENT '角色编码',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1正常 2禁用',
  `note` varchar(100) DEFAULT '' COMMENT '备注',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `create_user` int(11) NOT NULL DEFAULT '0' COMMENT '创建人',
  `create_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `update_user` int(11) NOT NULL DEFAULT '0' COMMENT '更新人',
  `update_time` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新时间',
  `mark` int(11) DEFAULT '1' COMMENT '标记',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

/*Data for the table `sys_role` */

insert  into `sys_role`(`id`,`name`,`code`,`status`,`note`,`sort`,`create_user`,`create_time`,`update_user`,`update_time`,`mark`) values (1,'超级管理员','AAA',1,'',1,1,'2022-10-26 16:01:44',1,'2022-10-26 17:49:53',1),(2,'部门负责人','BBB',1,'',2,1,'2022-10-26 16:01:44',1,'2022-10-26 17:49:44',1),(5,'普通员工','CCC',1,'',3,1,'2022-10-26 16:01:44',1,'2022-10-26 17:49:41',1);

/*Table structure for table `sys_role_menu` */

DROP TABLE IF EXISTS `sys_role_menu`;

CREATE TABLE `sys_role_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色ID',
  `menu_ids` text COMMENT '菜单IDs',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `role_id` (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

/*Data for the table `sys_role_menu` */

insert  into `sys_role_menu`(`id`,`role_id`,`menu_ids`) values (10,2,'1,2,3,6,5,15,4,21');

/*Table structure for table `sys_user` */

DROP TABLE IF EXISTS `sys_user`;

CREATE TABLE `sys_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `realname` varchar(150) DEFAULT 'NULL' COMMENT '真实姓名',
  `nickname` varchar(150) DEFAULT 'NULL' COMMENT '昵称',
  `gender` tinyint(1) DEFAULT '3' COMMENT '性别:1男 2女 3保密',
  `avatar` varchar(150) DEFAULT 'NULL' COMMENT '头像',
  `mobile` char(11) DEFAULT 'NULL' COMMENT '手机号码',
  `email` varchar(30) DEFAULT 'NULL' COMMENT '邮箱地址',
  `birthday` date DEFAULT NULL COMMENT '出生日期',
  `dept_id` int(11) DEFAULT '0' COMMENT '部门ID',
  `level_id` int(11) DEFAULT '0' COMMENT '职级ID',
  `position_id` smallint(3) DEFAULT '0' COMMENT '岗位ID',
  `province_code` varchar(50) DEFAULT 'NULL' COMMENT '省份编号',
  `city_code` varchar(50) DEFAULT 'NULL' COMMENT '市区编号',
  `district_code` varchar(50) DEFAULT 'NULL' COMMENT '区县编号',
  `address` varchar(255) DEFAULT 'NULL' COMMENT '详细地址',
  `city_name` varchar(150) DEFAULT 'NULL' COMMENT '所属城市',
  `username` varchar(50) DEFAULT 'NULL' COMMENT '登录用户名',
  `password` varchar(150) DEFAULT 'NULL' COMMENT '登录密码',
  `salt` varchar(30) DEFAULT 'NULL' COMMENT '盐加密',
  `intro` varchar(500) DEFAULT 'NULL' COMMENT '个人简介',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态：1正常 2禁用',
  `note` varchar(500) DEFAULT 'NULL' COMMENT '备注',
  `sort` int(11) DEFAULT '125' COMMENT '排序号',
  `login_num` int(11) DEFAULT '0' COMMENT '登录次数',
  `login_ip` varchar(20) DEFAULT 'NULL' COMMENT '最近登录IP',
  `login_time` datetime DEFAULT NULL COMMENT '最近登录时间',
  `create_user` int(10) DEFAULT '0' COMMENT '添加人',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_user` int(10) DEFAULT '0' COMMENT '更新人',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `mark` tinyint(1) NOT NULL DEFAULT '1' COMMENT '有效标识(1正常 0删除)',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

/*Data for the table `sys_user` */

insert  into `sys_user`(`id`,`realname`,`nickname`,`gender`,`avatar`,`mobile`,`email`,`birthday`,`dept_id`,`level_id`,`position_id`,`province_code`,`city_code`,`district_code`,`address`,`city_name`,`username`,`password`,`salt`,`intro`,`status`,`note`,`sort`,`login_num`,`login_ip`,`login_time`,`create_user`,`create_time`,`update_user`,`update_time`,`mark`) values (1,'admin','admin',3,NULL,NULL,NULL,NULL,0,0,0,NULL,NULL,NULL,NULL,NULL,'admin','43286a86708820e38c333cdd4c496355',NULL,NULL,1,NULL,125,0,NULL,'2022-10-27 16:51:34',0,'2022-10-26 17:46:54',0,'2022-10-27 16:51:34',1),(3,'汪进','Jerry',3,NULL,NULL,NULL,'2000-11-11',0,0,0,NULL,NULL,NULL,NULL,NULL,'wangjin','5cda335f4257eef20aa1b6a109655762',NULL,NULL,1,NULL,125,0,NULL,NULL,6,'2022-10-27 15:11:59',1,'2022-10-26 18:08:26',1),(6,'部门负责人','',0,'','','','2000-11-11',0,0,0,'','','','','','bumen','aecff0150102198b5b76f392ac233be0','','',1,'部门负责人',0,0,'','2022-10-29 15:33:51',1,'2022-10-26 17:44:33',1,'2022-10-29 15:33:51',1);

/*Table structure for table `sys_user_role` */

DROP TABLE IF EXISTS `sys_user_role`;

CREATE TABLE `sys_user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_id_role_id` (`user_id`,`role_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

/*Data for the table `sys_user_role` */

insert  into `sys_user_role`(`id`,`user_id`,`role_id`) values (48,3,1),(44,6,2),(45,6,5);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
