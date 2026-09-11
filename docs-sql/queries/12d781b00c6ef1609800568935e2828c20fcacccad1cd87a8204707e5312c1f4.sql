CREATE TABLE user_activity
(
    UserID UInt64,
    PageViews UInt8,
    Duration UInt8,
    Sign Int8
)
ENGINE = CollapsingMergeTree(Sign)
ORDER BY UserID

-- 초기 상태
INSERT INTO user_activity VALUES (4324182021466249494, 5, 146, 1)

-- 이전 행 취소 및 새 상태 삽입
INSERT INTO user_activity VALUES (4324182021466249494, 5, 146, -1)
INSERT INTO user_activity VALUES (4324182021466249494, 6, 185, 1)

-- 올바른 집계를 사용한 쿼리
SELECT
    UserID,
    sum(PageViews * Sign) AS PageViews,
    sum(Duration * Sign) AS Duration
FROM user_activity
GROUP BY UserID
HAVING sum(Sign) > 0
