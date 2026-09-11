CREATE TABLE user_activity
(
    UserID UInt64,
    PageViews UInt8,
    Duration UInt8,
    Sign Int8
)
ENGINE = CollapsingMergeTree(Sign)
ORDER BY UserID

-- 初期状態
INSERT INTO user_activity VALUES (4324182021466249494, 5, 146, 1)

-- 古い行をキャンセルして新しい状態を挿入
INSERT INTO user_activity VALUES (4324182021466249494, 5, 146, -1)
INSERT INTO user_activity VALUES (4324182021466249494, 6, 185, 1)

-- 適切な集計を用いたクエリ
SELECT
    UserID,
    sum(PageViews * Sign) AS PageViews,
    sum(Duration * Sign) AS Duration
FROM user_activity
GROUP BY UserID
HAVING sum(Sign) > 0
