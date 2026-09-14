-- 비트 시프트를 지정하는 tuple과 함께 단일 인수가 제공되면, 함수는
-- 해당 인수를 지정된 비트 수만큼 왼쪽으로 시프트합니다.
SELECT hilbertEncode(tuple(2), 128)
