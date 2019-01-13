CREATE TABLE `jk_user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(16) NOT NULL DEFAULT '' COMMENT '用户名',
  `user_pass` varchar(16) NOT NULL DEFAULT '' COMMENT '密码',
  `icon_url` varchar(255) NOT NULL DEFAULT '' COMMENT '头像',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `idx_uid` (`user_name`),
  KEY `idx_pass` (`user_pass`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户表';