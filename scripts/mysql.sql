create TABLE `blog_tag`(
	`id` int(10) UNSIGNED not null AUTO_INCREMENT,
	`name` VARCHAR(100) DEFAULT '' COMMENT '标签名称',
	`created_on` int(10) UNSIGNED DEFAULT 0 COMMENT '创建时间',
	`created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
	`modified_on` int(10) UNSIGNED DEFAULT 0 COMMENT '修改时间',
	`modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
	`deleted_on` int(10) UNSIGNED DEFAULT 0 COMMENT '删除时间',
	`is_del` tinyint(10) UNSIGNED DEFAULT 0 COMMENT '是否删除 0 未删除  1 已删除',
	`state` tinyint(3) UNSIGNED DEFAULT 1 COMMENT '状态 0 禁用  1 启用',
	PRIMARY KEY (`id`)
)ENGINE=Innodb DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';


create TABLE `blog_article`(
	`id` int(10) UNSIGNED not null AUTO_INCREMENT,
	`title` VARCHAR(100) DEFAULT '' COMMENT '文章标题',
	`desc` VARCHAR(255) DEFAULT '' COMMENT '文章简述',
	`cover_image_url` VARCHAR(255) DEFAULT '' COMMENT '封面图片地址',
	`content` longtext	COMMENT '文章内容',
	`created_on` int(10) UNSIGNED DEFAULT 0 COMMENT '创建时间',
	`created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
	`modified_on` int(10) UNSIGNED DEFAULT 0 COMMENT '修改时间',
	`modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
	`deleted_on` int(10) UNSIGNED DEFAULT 0 COMMENT '删除时间',
	`is_del` tinyint(10) UNSIGNED DEFAULT 0 COMMENT '是否删除 0 未删除  1 已删除',
	`state` tinyint(3) UNSIGNED DEFAULT 1 COMMENT '状态 0 禁用  1 启用',
	PRIMARY KEY (`id`)
)ENGINE=Innodb DEFAULT CHARSET=utf8mb4 COMMENT='文章管理';


create TABLE `blog_article_tag`(
	`id` int(10) UNSIGNED not null AUTO_INCREMENT,
	`article_id` int(11) not null COMMENT '文章ID',
	`tag_id` int(10) UNSIGNED not null DEFAULT 0 COMMENT '标签ID',
	`created_on` int(10) UNSIGNED DEFAULT 0 COMMENT '创建时间',
	`created_by` VARCHAR(100) DEFAULT '' COMMENT '创建人',
	`modified_on` int(10) UNSIGNED DEFAULT 0 COMMENT '修改时间',
	`modified_by` VARCHAR(100) DEFAULT '' COMMENT '修改人',
	`deleted_on` int(10) UNSIGNED DEFAULT 0 COMMENT '删除时间',
	`is_del` tinyint(10) UNSIGNED DEFAULT 0 COMMENT '是否删除 0 未删除  1 已删除',
	`state` tinyint(3) UNSIGNED DEFAULT 1 COMMENT '状态 0 禁用  1 启用',
	PRIMARY KEY (`id`)
)ENGINE=Innodb DEFAULT CHARSET=utf8mb4 COMMENT='文章标签关联';
