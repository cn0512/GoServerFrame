CREATE TABLE msvrs.users
(
    id bigint(20) PRIMARY KEY NOT NULL AUTO_INCREMENT,
    nickname varchar(33),
    account varchar(33) NOT NULL,
    pwd varchar(33) NOT NULL,
    regTime timestamp DEFAULT '0000-00-00 00:00:00' NOT NULL
);
CREATE UNIQUE INDEX users_id_uindex ON msvrs.users (id);
INSERT INTO msvrs.users (nickname, account, pwd, regTime) VALUES ('t1', 'test001', '123456', '2018-03-15 16:12:52');
INSERT INTO msvrs.users (nickname, account, pwd, regTime) VALUES ('t2', 'test002', '123456', '2018-03-15 16:12:52');