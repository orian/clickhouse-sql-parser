-- 튜플 없이 단일 인수만 제공된 경우, 차원 매핑이 필요하지 않으므로
-- 함수는 인수 자체를 Hilbert 인덱스로 반환합니다.
SELECT hilbertEncode(1)
