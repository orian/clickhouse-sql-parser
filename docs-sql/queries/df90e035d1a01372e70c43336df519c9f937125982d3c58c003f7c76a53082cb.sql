-- 범위가 크게 다른 인수(또는 카디널리티가 다른 인수)에 대해 유사한 분포가 필요할 때
-- 범위 확장이 유용할 수 있습니다.
-- 예: 'IP Address' (0...FFFFFFFF)와 'Country code' (0...FF).
-- 참고: Tuple의 크기는 나머지 인수의 수와 같아야 합니다.
SELECT mortonEncode((1,2), 1024, 16)
