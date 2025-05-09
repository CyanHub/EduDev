CREATE DATABASE `uuid_test`;

CREATE TABLE uuid_test.sequence_id(
    `id` bigint(20) unsigned NOT NULL auto_increment, 
    `value` char(10) NOT NULL default '',
    `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
);
insert into sequence_id(value) VALUES ('values');

CREATE TABLE id_generator (
    `id`            BIGINT NOT NULL AUTO_INCREMENT,
    `max_id`        BIGINT NOT NULL COMMENT '当前最大id',
    `step`          BIGINT NOT NULL COMMENT '号段的步长',
    `biz_type`      int NOT NULL COMMENT '业务类型',
    `version`       BIGINT NOT NULL COMMENT '版本号',
    `created_at`    DATETIME,
    `updated_at`    DATETIME,
    `deleted_at`    DATETIME,
    PRIMARY KEY (`id`)
);