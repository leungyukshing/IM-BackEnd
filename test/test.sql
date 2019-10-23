CREATE TABLE `admin_user` (
  `id` int(3) unsigned NOT NULL AUTO_INCREMENT,
  `user_name` varchar(32) NOT NULL,
  `password` varchar(32) NOT NULL,
  `email` varchar(32) DEFAULT NULL,
  `author` int(1) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
 
INSERT INTO `admin_user` VALUES ('1', 'admin', 'admin', 'admin@qq.com', '1');
