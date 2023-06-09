CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `type` tinyint DEFAULT '0' COMMENT '用户类型：0前台；1后台',
  `username` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `nickname` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '昵称',
  `phone` varchar(11) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '手机号',
  `email` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '邮箱',
  `gender` tinyint DEFAULT '0' COMMENT '性别：1男；2女',
  `age` tinyint DEFAULT '0' COMMENT '年龄',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '头像',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '备注',
  `lastIp` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '最后登录IP',
  `status` tinyint DEFAULT '1' COMMENT '状态：0禁用；1启用',
  `createdAt` datetime DEFAULT NULL COMMENT '创建时间',
  `updatedAt` datetime DEFAULT NULL COMMENT '更新时间',
  `deletedAt` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='用户表';

CREATE TABLE `role` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '角色名称',
  `mark` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '角色标志',
  `sort` int DEFAULT '0' COMMENT '排序',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '备注',
  `status` tinyint DEFAULT '1' COMMENT '状态：0禁用；1启用',
  `createdAt` datetime DEFAULT NULL COMMENT '创建时间',
  `updatedAt` datetime DEFAULT NULL COMMENT '更新时间',
  `deletedAt` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='角色表';

CREATE TABLE `menu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `parentId` int DEFAULT '0' COMMENT '父级id',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '菜单名称',
  `type` tinyint DEFAULT '0' COMMENT '类型：1目录 2菜单 3按钮',
  `key`varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '菜单标识',
  `icon` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '图标',
  `path` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '地址',
  `redirect` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '重定向',
  `component` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '组件名称',
  `remark` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '备注',
  `status` tinyint DEFAULT '1' COMMENT '状态：0禁用；1启用',
  `createdAt` datetime DEFAULT NULL COMMENT '创建时间',
  `updatedAt` datetime DEFAULT NULL COMMENT '更新时间',
  `deletedAt` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='菜单表';

CREATE TABLE `permission` (
  `id` int NOT NULL AUTO_INCREMENT,
  `path` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '路径',
  `method` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '请求方式',
  `group` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '分组',
  `description` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci COMMENT '描述',
  `createdAt` datetime DEFAULT NULL COMMENT '创建时间',
  `updatedAt` datetime DEFAULT NULL COMMENT '更新时间',
  `deletedAt` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='权限表';

CREATE TABLE `userHasRole` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userId` int DEFAULT '0' COMMENT '用户id',
  `roleId` int DEFAULT '0' COMMENT '角色id',
  `createdAt` datetime DEFAULT NULL COMMENT '创建时间',
  `updatedAt` datetime DEFAULT NULL COMMENT '更新时间',
  `deletedAt` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='用户&角色表';

CREATE TABLE `roleHasMenu` (
  `id` int NOT NULL AUTO_INCREMENT,
  `roleId` int DEFAULT '0' COMMENT '角色id',
  `menuId` int DEFAULT '0' COMMENT '菜单id',
  `createdAt` datetime DEFAULT NULL COMMENT '创建时间',
  `updatedAt` datetime DEFAULT NULL COMMENT '更新时间',
  `deletedAt` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='角色&菜单表';

CREATE TABLE `roleHasPermission` (
  `id` int NOT NULL AUTO_INCREMENT,
  `roleId` int DEFAULT '0' COMMENT '角色id',
  `permissionId` int DEFAULT '0' COMMENT '权限id',
  `createdAt` datetime DEFAULT NULL COMMENT '创建时间',
  `updatedAt` datetime DEFAULT NULL COMMENT '更新时间',
  `deletedAt` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='角色&权限表';

CREATE TABLE IF NOT EXISTS `optLog` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userId` int NOT NULL DEFAULT '0' COMMENT '操作人',
  `ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT 'ip地址',
  `method` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '请求方式',
  `path` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '请求地址',
  `status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '状态码',
  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '请求代理',
  `req` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '请求数据',
  `resp` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '响应数据',
  `createdAt` datetime DEFAULT NULL COMMENT '请求时间',
  `updatedAt` datetime DEFAULT NULL,
  `deletedAt` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='操作日志';