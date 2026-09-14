-- rate limiting 데이터가 있는 모든 로그를 찾습니다:
SELECT * FROM logs WHERE has(mapKeys(attributes), 'rate_limit'); -- 느린 전체 테이블 스캔

-- 특정 IP의 모든 로그를 찾습니다:
SELECT * FROM logs WHERE has(mapValues(attributes), '192.168.1.1'); -- 느린 전체 테이블 스캔
