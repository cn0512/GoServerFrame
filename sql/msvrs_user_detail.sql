CREATE TABLE msvrs.user_detail
(
    id bigint(20) PRIMARY KEY NOT NULL,
    age int(11),
    phone varchar(33),
    email varchar(33),
    CONSTRAINT user_detail_users_id_fk FOREIGN KEY (id) REFERENCES users (id)
);
CREATE UNIQUE INDEX user_detail_id_uindex ON msvrs.user_detail (id);
INSERT INTO msvrs.user_detail (id, age, phone, email) VALUES (10001, 12, '135', 'a@a.com');