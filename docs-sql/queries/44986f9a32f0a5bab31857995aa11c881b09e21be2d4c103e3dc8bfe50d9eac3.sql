CREATE TABLE my_test_table
ENGINE = MergeTree
ORDER BY tuple()
AS SELECT * 
FROM generateRandom(
    'col1 UInt32, col2 String, col3 Float64, col4 DateTime',
    1,  -- 데이터 생성용 시드
    10  -- 서로 다른 랜덤 값의 개수
)
LIMIT 100;  -- 100행

-- 2단계: 새로 만든 테이블 조회
SELECT * FROM my_test_table LIMIT 5;
