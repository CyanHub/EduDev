-- CREATE DATABASE go_shop
-- CHARACTER SET utf8mb4
-- COLLATE utf8mb4_unicode_ci;

USE go_shop;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '用户登录名',
  `password` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `role_id` bigint UNSIGNED NULL DEFAULT 888 COMMENT '用户角色Id',
  `phone` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '用户手机号',
  `email` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '用户邮箱',
  `enable` tinyint NULL DEFAULT 1 COMMENT '用户是否被冻结 1正常 2冻结',
  `balance` double NULL DEFAULT NULL COMMENT '用户余额',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_deleted_at`(`deleted_at` ASC) USING BTREE,
  UNIQUE INDEX `idx_user_username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '2024-12-13 15:22:48.194', '2024-12-13 15:22:48.194', NULL, 'kutqrx47', '5H84Ol2OqsG2S30', '明辉84', 'https://qmplusimg.henrongyi.top/gva_header.jpg', 888, '', '', 1, 1000);
INSERT INTO `user` VALUES (2, '2024-12-18 10:57:44.387', '2024-12-18 10:57:44.387', NULL, 'kutqrx49', '5H84Ol2OqsG2S30', '明辉84', 'https://qmplusimg.henrongyi.top/gva_header.jpg', 888, '', '', 1, 1000);
INSERT INTO `user` VALUES (3, '2024-12-18 10:57:44.387', '2024-12-18 10:57:44.387', NULL, 'kutqrx99', '5H84Ol2OqsG2S30', '明辉84', 'https://qmplusimg.henrongyi.top/gva_header.jpg', 888, '', '', 1, 1000);

SET FOREIGN_KEY_CHECKS = 1;