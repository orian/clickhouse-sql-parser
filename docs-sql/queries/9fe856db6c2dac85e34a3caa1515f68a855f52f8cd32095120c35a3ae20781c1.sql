-- 속도 제한된 모든 요청 조회:
SELECT * FROM logs WHERE mapContainsKey(attributes, 'rate_limit'); -- 빠름

-- 특정 IP의 모든 로그 조회:
SELECT * FROM logs WHERE has(mapValues(attributes), '192.168.1.1'); -- 빠름

-- 속성에 오류가 포함된 모든 로그 조회:
SELECT * FROM logs WHERE mapContainsValueLike(attributes, '% error %'); -- 빠름
