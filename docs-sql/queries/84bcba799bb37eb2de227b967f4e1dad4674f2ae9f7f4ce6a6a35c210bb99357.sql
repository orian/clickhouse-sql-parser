-- `hits` 테이블에서 `Title` 컬럼에 `hello` 텍스트가 포함된 모든 행을 삭제합니다
DELETE FROM hits WHERE Title LIKE '%hello%';
