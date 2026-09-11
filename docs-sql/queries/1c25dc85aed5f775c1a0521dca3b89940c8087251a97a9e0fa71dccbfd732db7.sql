-- 테이블 생성 및 데이터 삽입
CREATE TABLE test_lazy (json JSON) ENGINE = MergeTree ORDER BY tuple();
INSERT INTO test_lazy VALUES ('{"user_id": "123", "score": "95.5"}');

-- 실험적 설정 활성화
SET allow_experimental_json_lazy_type_hints = 1;

-- 타입 힌트 추가 - 뮤테이션 없이 즉시 완료됨
ALTER TABLE test_lazy MODIFY COLUMN json JSON(user_id UInt64, score Float64);

-- 데이터 쿼리 - 타입 힌트는 읽기 시점에 적용됨
SELECT json.user_id, toTypeName(json.user_id), json.score, toTypeName(json.score) FROM test_lazy;
