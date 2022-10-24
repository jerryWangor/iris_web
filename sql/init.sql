use easygoadmin;

CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `realname` varchar(20) not null DEFAULT '' comment '真实姓名',
  `nickname` varchar(20) not null DEFAULT '' comment '昵称',
  `gender` tinyint(1) not null DEFAULT '1' comment '性别 1男 2女 3保密',
  `avatar` varchar(200) DEFAULT '' comment '头像地址',
  `mobile` varchar(20) DEFAULT '' comment '手机号',
  `email` varchar(20) DEFAULT '' comment '邮箱',
  `birthday` date comment '生日',
  `deptid` int(11) DEFAULT '0' comment '部门ID',
  `levelid` int(11) DEFAULT '0' comment '职级ID',
  `positionid` int(11) DEFAULT '0' comment '岗位ID',
  `provincecode` varchar(20) DEFAULT '' comment '国家',
  `citycode` varchar(20) DEFAULT '' comment '城市',
  `districtcode` varchar(20) DEFAULT '' comment '地区',
  `address` varchar(100) DEFAULT '' comment '地址',
  `username` varchar(100) not null DEFAULT '' comment '用户名',
  `password` varchar(100) not null DEFAULT '' comment '密码',
  `intro` varchar(100) DEFAULT '' comment '个人简介',
  `status` tinyint(1) not null DEFAULT '1' comment '状态 1正常 2禁用',
  `note` varchar(100) DEFAULT '' comment '备注',
  `sort` int(11) DEFAULT '0' comment '排序',
  `roleids` text not null comment '角色IDs',
  `logintime` int(11) DEFAULT '0' comment '登录时间',
  `loginip` varchar(20) DEFAULT '' comment 'IP',
  `createuser` int(11) not null DEFAULT '0' comment '创建人',
  `createtime` int(11) not null DEFAULT '0' comment '创建时间',
  `updateuser` int(11) not null DEFAULT '0' comment '更新人',
  `updatetime` int(11) not null DEFAULT '0' comment '更新时间',
  `mark` int(11) DEFAULT '0' comment '标记',
  PRIMARY KEY (`id`),
  unique key (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

CREATE TABLE `menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) not null DEFAULT '' comment '菜单名',
  `icon` varchar(100) DEFAULT '' comment '菜单图标',
  `url` varchar(100) DEFAULT '' comment '菜单地址',
  `param` varchar(100) DEFAULT '' comment '额外参数',
  `pid` int(11) not null default '0' comment '上级ID',
  `type` tinyint(1) not null default '1' comment '类型：1模块 2导航 3菜单 4节点',
  `permission` varchar(100) DEFAULT '' comment '权限标识',
  `status` tinyint(1) not null DEFAULT '1' comment '状态 1正常 2禁用',
  `target` tinyint(1) not null default '1' comment '打开方式：1内部打开 2外部打开',
  `note` varchar(100) DEFAULT '' comment '备注',
  `sort` int(11) DEFAULT '0' comment '排序',
  `func` text comment '权限节点',
  `createuser` int(11) not null DEFAULT '0' comment '创建人',
  `createtime` int(11) not null DEFAULT '0' comment '创建时间',
  `updateuser` int(11) not null DEFAULT '0' comment '更新人',
  `updatetime` int(11) not null DEFAULT '0' comment '更新时间',
  `mark` int(11) DEFAULT '0' comment '标记',
  PRIMARY KEY (`id`),
  unique key (`url`, `name`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) not null DEFAULT '' comment '角色名',
  `code` varchar(100) DEFAULT '' comment '角色编码',
  `status` tinyint(1) not null DEFAULT '1' comment '状态 1正常 2禁用',
  `note` varchar(100) DEFAULT '' comment '备注',
  `sort` int(11) DEFAULT '0' comment '排序',
  `createuser` int(11) not null DEFAULT '0' comment '创建人',
  `createtime` int(11) not null DEFAULT '0' comment '创建时间',
  `updateuser` int(11) not null DEFAULT '0' comment '更新人',
  `updatetime` int(11) not null DEFAULT '0' comment '更新时间',
  `mark` int(11) DEFAULT '0' comment '标记',
  PRIMARY KEY (`id`),
  unique key (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

CREATE TABLE `rolemenu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `roleid` int(11) not null DEFAULT '0' comment '角色ID',
  `menuids` text comment '菜单IDs',
  `createuser` int(11) not null DEFAULT '0' comment '创建人',
  `createtime` int(11) not null DEFAULT '0' comment '创建时间',
  `updateuser` int(11) not null DEFAULT '0' comment '更新人',
  `updatetime` int(11) not null DEFAULT '0' comment '更新时间',
  `mark` int(11) DEFAULT '0' comment '标记',
  PRIMARY KEY (`id`),
  unique key (`roleid`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;


-----
系统管理
权限管理
用户管理	角色管理	菜单管理

用户列表    
添加用户    添加角色	添加菜单
编辑用户    编辑角色	编辑菜单
删除用户    删除角色	删除菜单
重置密码    角色权限
