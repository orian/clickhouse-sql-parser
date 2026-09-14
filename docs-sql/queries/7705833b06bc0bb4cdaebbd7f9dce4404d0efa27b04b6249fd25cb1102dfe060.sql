-- 먼저 테이블을 생성하고 데이터를 삽입합니다
CREATE TABLE hilbert_numbers(
    n1 UInt32,
    n2 UInt32
)
ENGINE=MergeTree()
ORDER BY n1;
insert into hilbert_numbers (*) values(1, 2);

-- 함수 인수로 상수 대신 컬럼 이름을 사용합니다
SELECT hilbertEncode(n1, n2) FROM hilbert_numbers;
