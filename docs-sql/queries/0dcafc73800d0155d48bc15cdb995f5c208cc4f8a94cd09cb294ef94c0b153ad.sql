-- 원본 테이블 생성
CREATE TABLE source_data (
    id UInt32,
    name String,
    value Float64
) ENGINE = MergeTree
ORDER BY id;

-- 데이터 삽입
INSERT INTO source_data VALUES (1, 'one', 10.1), (2, 'two', 20.2);

-- 별칭 생성
CREATE TABLE data_alias ENGINE = Alias('source_data');

-- 별칭을 통한 쿼리
SELECT * FROM data_alias;
