drop table if exists `user_tbl`;
create table if not exists `user_tbl` (
    `id` BIGINT,
    `fullname` VARCHAR(30),
    `email` VARCHAR(50) NOT NULL,
    `hash_pw` VARCHAR(100),
    PRIMARY KEY (id)
);