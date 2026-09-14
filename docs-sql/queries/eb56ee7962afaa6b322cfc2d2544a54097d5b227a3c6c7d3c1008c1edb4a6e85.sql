-- 경량한 삭제를 사용하여 2018년의 모든 데이터를 삭제합니다. 권장되지 않습니다.
DELETE FROM posts WHERE toYear(CreationDate) = 2018
