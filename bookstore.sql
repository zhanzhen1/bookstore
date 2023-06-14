/*
 Navicat Premium Data Transfer

 Source Server         : demo
 Source Server Type    : MySQL
 Source Server Version : 50737
 Source Host           : localhost:3306
 Source Schema         : bookstore

 Target Server Type    : MySQL
 Target Server Version : 50737
 File Encoding         : 65001

 Date: 14/06/2023 19:20:28
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for book
-- ----------------------------
DROP TABLE IF EXISTS `book`;
CREATE TABLE `book`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `author` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `price` double(11, 2) NOT NULL,
  `sales` int(11) NOT NULL,
  `stock` int(11) NOT NULL,
  `img_path` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of book
-- ----------------------------
INSERT INTO `book` VALUES (1, '解忧杂货店', '东野圭吾', 27.20, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (2, '边城', '沈从文', 23.00, 101, 99, 'static/img/default.jpg');
INSERT INTO `book` VALUES (3, '中国哲学史', '冯友兰', 44.50, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (4, '忽然七日', ' 劳伦', 19.33, 101, 99, 'static/img/default.jpg');
INSERT INTO `book` VALUES (5, '苏东坡传', '林语堂', 19.30, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (6, '百年孤独', '马尔克斯', 29.50, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (7, '扶桑', '严歌苓', 19.80, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (8, '给孩子的诗', '北岛', 22.20, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (9, '为奴十二年', '所罗门', 16.50, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (10, '平凡的世界', '路遥', 55.00, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (11, '悟空传', '今何在', 14.00, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (12, '硬派健身', '斌卡', 31.20, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (13, '从晚清到民国', '唐德刚', 39.90, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (14, '三体', '刘慈欣', 56.50, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (15, '看见', '柴静', 19.50, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (16, '活着', '余华', 11.00, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (17, '小王子', '安托万', 19.20, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (18, '我们仨', '杨绛', 11.30, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (19, '生命不息,折腾不止', '罗永浩', 25.20, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (20, '皮囊', '蔡崇达', 23.90, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (21, '恰到好处的幸福', '毕淑敏', 16.40, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (22, '大数据预测', '埃里克', 37.20, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (23, '人月神话', '布鲁克斯', 55.90, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (24, 'C语言入门经典', '霍尔顿', 45.00, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (25, '数学之美', '吴军', 29.90, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (26, 'Java编程思想', '埃史尔', 70.50, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (27, '设计模式之禅', '秦小波', 20.20, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (28, '图解机器学习', '杉山将', 33.80, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (29, '艾伦图灵传', '安德鲁', 47.20, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (30, '教父', '马里奥普佐', 29.00, 100, 100, 'static/img/default.jpg');
INSERT INTO `book` VALUES (34, '从入门到入土', 'zz', 123.00, 100, 100, '/bookstore/view/static/img/logo.gif');

-- ----------------------------
-- Table structure for cart
-- ----------------------------
DROP TABLE IF EXISTS `cart`;
CREATE TABLE `cart`  (
  `cart_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `total_count` int(11) NOT NULL,
  `total_amount` double(11, 2) NOT NULL,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`cart_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `cart_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cart
-- ----------------------------
INSERT INTO `cart` VALUES ('ba25b221-12b2-7674-45a4%!(EXTRA []uint8=[232 96 50 61 56 101])', 6, 146.40, 2);

-- ----------------------------
-- Table structure for cart_item
-- ----------------------------
DROP TABLE IF EXISTS `cart_item`;
CREATE TABLE `cart_item`  (
  `cart_item_id` int(11) NOT NULL AUTO_INCREMENT,
  `count` int(11) NOT NULL,
  `amount` double(11, 2) NOT NULL,
  `book_id` int(11) NOT NULL,
  `cart_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`cart_item_id`) USING BTREE,
  INDEX `book_id`(`book_id`) USING BTREE,
  INDEX `cart_id`(`cart_id`) USING BTREE,
  CONSTRAINT `cart_item_ibfk_1` FOREIGN KEY (`book_id`) REFERENCES `book` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `cart_item_ibfk_2` FOREIGN KEY (`cart_id`) REFERENCES `cart` (`cart_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cart_item
-- ----------------------------
INSERT INTO `cart_item` VALUES (17, 4, 92.00, 2, 'ba25b221-12b2-7674-45a4%!(EXTRA []uint8=[232 96 50 61 56 101])');
INSERT INTO `cart_item` VALUES (18, 2, 54.40, 1, 'ba25b221-12b2-7674-45a4%!(EXTRA []uint8=[232 96 50 61 56 101])');

-- ----------------------------
-- Table structure for order_item
-- ----------------------------
DROP TABLE IF EXISTS `order_item`;
CREATE TABLE `order_item`  (
  `order_item_id` int(11) NOT NULL AUTO_INCREMENT,
  `count` int(11) NOT NULL,
  `amount` double(11, 2) NOT NULL,
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `author` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `price` double(11, 2) NOT NULL,
  `img_path` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `order_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  PRIMARY KEY (`order_item_id`) USING BTREE,
  INDEX `order_id`(`order_id`) USING BTREE,
  CONSTRAINT `order_item_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`order_id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order_item
-- ----------------------------
INSERT INTO `order_item` VALUES (1, 1, 100.00, '入门到入土', 'zz', 100.00, 'bookstore/view/static/img/default.jpg', '9361f701-46cb-6220-5bc0%!(EXTRA []uint8=[91 186 150 181 240 245])');
INSERT INTO `order_item` VALUES (2, 1, 100.00, '入门到入土11', 'zz', 100.00, 'bookstore/view/static/img/default.jpg', '9361f701-46cb-6220-5bc0%!(EXTRA []uint8=[91 186 150 181 240 245])');
INSERT INTO `order_item` VALUES (3, 1, 23.00, '边城', '沈从文', 23.00, 'static/img/default.jpg', 'a790d4a1-1141-6d98-7ff3%!(EXTRA []uint8=[215 253 228 37 249 122])');
INSERT INTO `order_item` VALUES (4, 1, 19.33, '忽然七日', ' 劳伦', 19.33, 'static/img/default.jpg', 'a790d4a1-1141-6d98-7ff3%!(EXTRA []uint8=[215 253 228 37 249 122])');

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders`  (
  `order_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `create_time` datetime(0) NOT NULL,
  `total_count` int(11) NOT NULL,
  `total_amount` double(11, 2) NOT NULL,
  `state` int(11) NOT NULL,
  `user_id` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`order_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of orders
-- ----------------------------
INSERT INTO `orders` VALUES ('9361f701-46cb-6220-5bc0%!(EXTRA []uint8=[91 186 150 181 240 245])', '2023-06-14 07:53:46', 2, 200.00, 0, 2);
INSERT INTO `orders` VALUES ('a790d4a1-1141-6d98-7ff3%!(EXTRA []uint8=[215 253 228 37 249 122])', '2023-06-14 08:33:21', 2, 42.33, 0, 1);

-- ----------------------------
-- Table structure for session
-- ----------------------------
DROP TABLE IF EXISTS `session`;
CREATE TABLE `session`  (
  `session_id` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `username` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `user_id` int(11) NOT NULL,
  PRIMARY KEY (`session_id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE,
  CONSTRAINT `session_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of session
-- ----------------------------
INSERT INTO `session` VALUES ('2c83b585-ba6b-7f4e-498c%!(EXTRA []uint8=[249 93 160 100 37 84])', 'admin', 1);
INSERT INTO `session` VALUES ('5b0154cc-a09a-40dc-76fb%!(EXTRA []uint8=[255 235 102 229 161 195])', 'admin', 1);

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `password` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `eamil` varchar(199) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, 'admin', '123456', '12345@qq.com');
INSERT INTO `user` VALUES (2, 'zz', '1234', '213234.@qq.com');
INSERT INTO `user` VALUES (9, 'admin1234', '12345', 'admin1234@qq.com');

SET FOREIGN_KEY_CHECKS = 1;
