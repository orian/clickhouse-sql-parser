CREATE TABLE user_activity
(
    UserID UInt64,
    PageViews UInt8,
    Duration UInt8,
    Sign Int8
)
ENGINE = CollapsingMergeTree(Sign)
ORDER BY UserID

-- 初始状态
INSERT INTO user_activity VALUES (4324182021466249494, 5, 146, 1)

-- 取消旧行并插入新状态
INSERT INTO user_activity VALUES (4324182021466249494, 5, 146, -1)
INSERT INTO user_activity VALUES (4324182021466249494, 6, 185, 1)

-- 使用正确的聚合方式进行查询
SELECT
    UserID,
    sum(PageViews * Sign) AS PageViews,
    sum(Duration * Sign) AS Duration
FROM user_activity
GROUP BY UserID
HAVING sum(Sign) > 0
