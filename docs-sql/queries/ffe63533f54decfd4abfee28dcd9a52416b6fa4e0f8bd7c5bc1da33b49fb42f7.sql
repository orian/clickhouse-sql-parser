CREATE TABLE user_activity
(
    UserID UInt64,
    PageViews UInt8,
    Duration UInt8,
    Sign Int8
)
ENGINE = CollapsingMergeTree(Sign)
ORDER BY UserID

-- Начальное состояние
INSERT INTO user_activity VALUES (4324182021466249494, 5, 146, 1)

-- Отменить старую строку и вставить новое состояние
INSERT INTO user_activity VALUES (4324182021466249494, 5, 146, -1)
INSERT INTO user_activity VALUES (4324182021466249494, 6, 185, 1)

-- Запрос с корректной агрегацией
SELECT
    UserID,
    sum(PageViews * Sign) AS PageViews,
    sum(Duration * Sign) AS Duration
FROM user_activity
GROUP BY UserID
HAVING sum(Sign) > 0
