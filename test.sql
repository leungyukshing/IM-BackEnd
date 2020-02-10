
CREATE TABLE `user` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(32) NOT NULL,
    `password` varchar(32) NOT NULL,
    `email` varchar(32) NOT NULL,
    `time` datetime DEFAULT NULL,
    `encryptkey` varchar(32) ,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
 
INSERT INTO `admin_user` VALUES ('1', 'admin', 'admin', 'admin@qq.com',NULL,NULL);


CREATE TABLE `friend` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `userid1` int(11) NOT NULL,
    `userid2` int(11) NOT NULL,
    `time` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
 

CREATE TABLE `chat` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `userid` int(11) NOT NULL,
    `chatname` varchar(32) NOT NULL,
    `create_time` datetime DEFAULT NULL,
    `last_update_time` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;


CREATE TABLE `chat_user_ref` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `userid` int(11) NOT NULL,
    `chatid` varchar(32) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;