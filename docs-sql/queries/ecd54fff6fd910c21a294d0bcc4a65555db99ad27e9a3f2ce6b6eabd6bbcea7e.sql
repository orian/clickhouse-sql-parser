CREATE TABLE prod
(
  uid Int16,
  name String,
  age Int16
)
ENGINE=MergeTree
ORDER BY ();

CREATE TABLE staging
(
  uid Int16,
  name String,
  age Int16
)
ENGINE=MergeTree
ORDER BY ();

-- 초기 데이터
INSERT INTO prod VALUES (123, 'John', 33);
INSERT INTO prod VALUES (456, 'Ksenia', 48);
-- 데이터 로드
INSERT INTO staging VALUES (8811, 'Alice', 50);
INSERT INTO staging VALUES (8812, 'Bob', 23);

-- 가져오기 검증
SELECT 'Staging count:', COUNT() FROM staging;
-- 파티션 이동
ALTER TABLE staging MOVE PARTITION tuple() TO TABLE prod; -- 원자적 연산

-- 데이터 확인
SELECT 'Prod count:', COUNT() FROM prod;
SELECT * FROM prod;
