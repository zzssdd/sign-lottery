DROP TABLE IF EXISTS user;
CREATE TABLE user(
                     id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                     created_at DATETIME(3) DEFAULT NULL,
                     deleted_at DATETIME(3) DEFAULT NULL,
                     updated_at DATETIME(3) DEFAULT NULL,
                     email VARCHAR(50) NOT NULL,
                     name VARCHAR(30) NOT NULL,
                     password VARCHAR(100) NOT NULL,
                     avater VARCHAR(50) DEFAULT NULL,
                     address VARCHAR(50) DEFAULT NULL,
                     UNIQUE(email,deleted_at),
                     KEY key_deleted(deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS sign_group;
CREATE TABLE sign_group(
                           id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                           name VARCHAR(50) NOT NULL,
                           created_at DATETIME(3) DEFAULT NULL,
                           deleted_at DATETIME(3) DEFAULT NULL,
                           updated_at DATETIME(3) DEFAULT NULL,
                           start DATETIME NOT NULL,
                           end DATETIME NOT NULL,
                           count BIGINT UNSIGNED DEFAULT 1,
                           avater VARCHAR(50) DEFAULT NULL,
                           owner BIGINT UNSIGNED NOT NULL,
                           KEY key_deleted (deleted_at),
                           KEY key_created (created_at)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS user_group;
CREATE TABLE user_group(
                           id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                           created_at DATETIME(3) DEFAULT NULL,
                           deleted_at DATETIME(3) DEFAULT NULL,
                           updated_at DATETIME(3) DEFAULT NULL,
                           uid BIGINT UNSIGNED NOT NULL,
                           gid INT UNSIGNED NOT NULL,
                           score BIGINT UNSIGNED DEFAULT 0,
                           count INT UNSIGNED DEFAULT 0,
                           KEY key_deleted(deleted_at),
                           KEY key_gid(gid),
                           UNIQUE (uid,gid,deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS sign_month;
CREATE TABLE sign_month(
                           id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                           created_at DATETIME(3) DEFAULT NULL,
                           deleted_at DATETIME(3) DEFAULT NULL,
                           updated_at DATETIME(3) DEFAULT NULL,
                           month char(7) NOT NULL,
                           gid INT UNSIGNED NOT NULL ,
                           bit_sign INT DEFAULT 0,
                           uid BIGINT UNSIGNED NOT NULL,
                           KEY key_deleted(deleted_at),
                           KEY key_gid(gid),
                           unique(uid,month,gid) ,
                           KEY key_month(month)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS sign_record;
CREATE TABLE sign_record(
                            id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                            created_at DATETIME(3) DEFAULT NULL,
                            deleted_at DATETIME(3) DEFAULT NULL,
                            updated_at DATETIME(3) DEFAULT NULL,
                            gid INT UNSIGNED NOT NULL ,
                            uid BIGINT UNSIGNED NOT NULL,
                            KEY key_deleted(deleted_at),
                            KEY key_gid(gid),
                            UNIQUE (uid,gid,deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE  IF EXISTS sign_group_pos;
CREATE TABLE sign_group_pos(
                     id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                     name VARCHAR(50) NOT NULL,
                     gid INT UNSIGNED NOT NULL ,
                     created_at DATETIME(3) DEFAULT NULL,
                     deleted_at DATETIME(3) DEFAULT NULL,
                     updated_at DATETIME(3) DEFAULT NULL,
                     longtitude DOUBLE(9,6) NOT NULL,
                     latitude DOUBLE(9,6) NOT NULL,
                     KEY key_deleted(deleted_at),
                     UNIQUE (gid,name,longtitude,latitude,deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS ask_leave;
CREATE TABLE  ask_leave(
                           id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                           uid BIGINT UNSIGNED NOT NULL,
                           created_at DATETIME(3) DEFAULT NULL,
                           deleted_at DATETIME(3) DEFAULT NULL,
                           updated_at DATETIME(3) DEFAULT NULL,
                           issue TEXT NOT NULL,
                           time DATETIME DEFAULT NULL,
                           gid INT UNSIGNED NOT NULL,
                           KEY key_deleted(deleted_at),
                           UNIQUE (gid,uid,time,deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS admin;
CREATE TABLE admin(
                      id INT UNSIGNED PRIMARY KEY  AUTO_INCREMENT,
                      created_at DATETIME(3) DEFAULT NULL,
                      deleted_at DATETIME(3) DEFAULT NULL,
                      updated_at DATETIME(3) DEFAULT NULL,
                      name VARCHAR(50) NOT NULL,
                      password VARCHAR(100) NOT NULL,
                      KEY key_deleted(deleted_at),
                      UNIQUE (name,password,deleted_at)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS activity;
CREATE TABLE activity(
                         id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                         created_at DATETIME(3) DEFAULT NULL,
                         deleted_at DATETIME(3) DEFAULT NULL,
                         updated_at DATETIME(3) DEFAULT NULL,
                         name VARCHAR(50) NOT NULL,
                         des TINYTEXT DEFAULT NULL,
                         picture VARCHAR(30) DEFAULT NULL,
                         cost INT UNSIGNED DEFAULT 0,
                         uid BIGINT UNSIGNED DEFAULT 0,
                         gid INT UNSIGNED NOT NULL,
                         start DATETIME  DEFAULT NULL,
                         end DATETIME DEFAULT NULL,
                         num BIGINT UNSIGNED DEFAULT 0,
                         count INT UNSIGNED DEFAULT 0,
                         KEY key_deleted(deleted_at),
                         KEY key_uid(uid),
                         KEY key_gid(gid)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS prize;
CREATE TABLE prize(
                      id INT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                      created_at DATETIME(3) DEFAULT NULL,
                      deleted_at DATETIME(3) DEFAULT NULL,
                      updated_at DATETIME(3) DEFAULT NULL,
                      name VARCHAR(50) NOT NULL,
                      num BIGINT UNSIGNED DEFAULT 0,
                      picture VARCHAR(30) DEFAULT NULL,
                      aid INT UNSIGNED NOT NULL,
                      KEY key_deleted(deleted_at),
                      KEY key_aid(aid)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DROP TABLE IF EXISTS user_order;
CREATE TABLE user_order(
                           id BIGINT UNSIGNED PRIMARY KEY AUTO_INCREMENT,
                           created_at DATETIME(3) DEFAULT NULL,
                           deleted_at DATETIME(3) DEFAULT NULL,
                           updated_at DATETIME(3) DEFAULT NULL,
                           uid BIGINT UNSIGNED NOT NULL,
                           pid INT UNSIGNED NOT NULL,
                           KEY key_deleted(deleted_at),
                           KEY key_uid(uid,pid)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;