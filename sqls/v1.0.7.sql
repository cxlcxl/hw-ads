alter table users add column is_internal tinyint(1) unsigned default 1 comment '是否内部账号';