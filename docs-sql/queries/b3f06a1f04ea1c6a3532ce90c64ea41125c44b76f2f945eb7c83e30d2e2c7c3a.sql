system.parts 테이블을 사용하여 crypto_prices 테이블의 파트가 어느 디스크에 있는지 확인합니다:

SELECT
    name,
    disk_name
FROM system.parts
WHERE (table = 'my_table') AND (active = 1)
