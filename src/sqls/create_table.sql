create table `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username`  varchar(64) COLLATE  uft8mb4_general_ci NOT NULL,
    `password` varchar(64) COLLATE  uft8mb4_general_ci NOT NULL,
    `email` varchar(64) COLLATE  uft8mb4_general_ci,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP on UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_username` (`username`) USING BTREE,
    UNIQUE KEY `id_user_id` (`user_id`) USING BTREE)
    ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

drop table exists `community`;
create table `community` (
    `id` int(11) not null auto_increment,
    `community_id` int(10) unsigned not null,
    `community_name` varchar(128) collate utf8mb4_general_ci not null,
    `introduction` varchar(256) collate utf8mb4_general_ci not null,
    `create_time` timestamp not null default CURRENT_TIMESTAMP,
    `update_time` timestamp null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    primary key (`id`),
    unique key `idx_community_id` (`community_id`),
    unique key `idx_community_name` (`community_name`)
) engine=InnoDB default charset=utf8mb4 collate=utf8mb4_general_ci;
insert into `community` values ('1', '1', 'lol', '!', '2020-01-01 08:01:01', '2020-01-01 08:01:01')

drop table if exists `post`;
create table `post` (
                        `id` bigint(20) NOT NULL AUTO_INCREMENT,
                        `post_id` bigint(20) NOT NULL COMMENT '帖子id',
                        `title` varchar(128) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
                        `content` varchar(8192) COLLATE utf8mb4_general_ci NOT NULL COMMENT '内容',
                        `author_id` bigint(20) NOT NULL COMMENT '作者id',
                        `community_id` bigint(20) NOT NULL COMMENT '社区Tagid',
                        `status` tinyint(4) NOT NULL COMMENT '状态',
                        `create_time` timestamp not null default CURRENT_TIMESTAMP,
                        `update_time` timestamp null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_post_id` (`post_id`),
                        KEY `idx_author_id` (`author_id`),
                        KEY `idx_community_id` (`community_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;