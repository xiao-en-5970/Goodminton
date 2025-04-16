-- 用户表
CREATE TABLE `user` (
  `user_id` BIGINT NOT NULL AUTO_INCREMENT ,-- '用户编号'
  `nickname` VARCHAR(50) DEFAULT NULL ,-- '可选，仅供查看'
  `avatar_path` VARCHAR(255) NOT NULL DEFAULT 'default-avatar.png' ,-- '头像路径(相对路径)'
  `self_evaluated_level` VARCHAR(20) DEFAULT NULL ,-- '自评技术水平'
  `system_score` INT DEFAULT NULL ,-- '系统评估得分(0~100)'
  `personality_tags` JSON DEFAULT NULL ,-- '性格标签'
  `play_style_tags` JSON DEFAULT NULL ,-- '打球风格'
  `preferred_skill_level` VARCHAR(20) DEFAULT NULL ,-- '希望对手技术水平'
  `preferred_time_slots` JSON DEFAULT NULL ,-- '时间偏好'
  `preferred_regions` JSON DEFAULT NULL ,-- '常活动区域'
  `max_cost` INT DEFAULT NULL ,-- '可接受的花销（单位：元）'
  `historical_partners` JSON DEFAULT NULL ,-- '历史搭档ID列表'
  `ratings_given` JSON DEFAULT NULL ,-- '对别人的评价(用户ID:评分)'
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ,-- '创建时间'
  `update_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,-- '更新时间'
  PRIMARY KEY (`user_id`),
  CHECK (`system_score` BETWEEN 0 AND 100),
  CHECK (`max_cost` >= 0)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ;--'用户表'