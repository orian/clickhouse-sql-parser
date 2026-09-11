-- (1) 삽입 일시 중지
-- (2) 대상 테이블의 복제본 생성

CREATE TABLE pypi_v2 AS pypi

SELECT count() FROM pypi_v2
