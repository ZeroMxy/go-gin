CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userType` tinyint NOT NULL DEFAULT '0' COMMENT '用户类型：0前台用户；1后台用户；',
  `roleId` int NOT NULL DEFAULT '9999' COMMENT '角色id',
  `username` varchar(20) CHARACTER SET utf8 COLLATE utf8_General_Ci NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_General_Ci NOT NULL COMMENT '密码',
  `nickname` varchar(20) CHARACTER SET utf8 COLLATE utf8_General_Ci DEFAULT NULL COMMENT '昵称',
  `phone` varchar(11) CHARACTER SET utf8 COLLATE utf8_General_Ci DEFAULT NULL COMMENT '手机号',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_General_Ci DEFAULT NULL COMMENT '邮箱',
  `sex` tinyint NOT NULL DEFAULT '0' COMMENT '性别：1男；2女',
  `age` tinyint NOT NULL DEFAULT '0' COMMENT '年龄',
  `avatar` varchar(50) CHARACTER SET utf8 COLLATE utf8_General_Ci DEFAULT NULL COMMENT '头像',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：0禁用；1启用',
  `createdAt` datetime DEFAULT NULL COMMENT '创建时间',
  `updatedAt` datetime DEFAULT NULL COMMENT '更新时间',
  `deletedAt` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='用户表';

CREATE TABLE `role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `menuId` LONGTEXT NULL DEFAULT NULL COMMENT '菜单id' COLLATE 'utf8_General_Ci',
  `roleName` varchar(20) CHARACTER SET utf8 COLLATE utf8_General_Ci DEFAULT NULL COMMENT '角色名称',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：0禁用；1启用',
  `createdAt` datetime DEFAULT NULL COMMENT '创建时间',
  `updatedAt` datetime DEFAULT NULL COMMENT '更新时间',
  `deletedAt` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='角色表';

CREATE TABLE `menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `parentId` int NOT NULL DEFAULT '0' COMMENT '父级id',
  `menuName` varchar(50) CHARACTER SET utf8 COLLATE utf8_General_Ci DEFAULT NULL COMMENT '菜单名称',
  `key`varchar(50) CHARACTER SET utf8 COLLATE utf8_General_Ci DEFAULT NULL COMMENT '菜单标识',
  `icon` varchar(50) CHARACTER SET utf8 COLLATE utf8_General_Ci DEFAULT NULL COMMENT '图标',
  `path` varchar(50) CHARACTER SET utf8 COLLATE utf8_General_Ci NOT NULL COMMENT '地址',
  `redirect` varchar(50) CHARACTER SET utf8 COLLATE utf8_General_Ci DEFAULT NULL COMMENT '重定向',
  `component` varchar(50) CHARACTER SET utf8 COLLATE utf8_General_Ci DEFAULT NULL COMMENT '组件名称',
  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态：0禁用；1启用',
  `createdAt` datetime DEFAULT NULL COMMENT '创建时间',
  `updatedAt` datetime DEFAULT NULL COMMENT '更新时间',
  `deletedAt` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='菜单表';

CREATE TABLE `permission` (
  `id` int NOT NULL AUTO_INCREMENT,
  `menuId` int NOT NULL DEFAULT '0' COMMENT '菜单id',
  `permissionName` varchar(20) CHARACTER SET utf8 COLLATE utf8_General_Ci DEFAULT NULL COMMENT '路由名称',
  `path` varchar(50) CHARACTER SET utf8 COLLATE utf8_General_Ci DEFAULT NULL COMMENT '路由地址',
  `method` varchar(20) CHARACTER SET utf8 COLLATE utf8_General_Ci DEFAULT NULL COMMENT '请求方式',
  `createdAt` datetime DEFAULT NULL COMMENT '创建时间',
  `updatedAt` datetime DEFAULT NULL COMMENT '更新时间',
  `deletedAt` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='权限表';