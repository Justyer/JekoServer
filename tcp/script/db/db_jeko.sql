-- create database db_jeko CHARACTER SET utf8;

create table jk_user (
    `user_id` int auto_increment,
    `user_name` varchar(16) not null default '' comment '用户名',
    `user_pass` varchar(16) not null default '' comment '密码',
    primary key(`user_id`),
    unique key `idx_uid` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=1000 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户表';

