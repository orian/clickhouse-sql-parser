DELETE FROM system.webassembly_modules WHERE name = 'collatz';

-- 이름이 `tmp_`로 시작하는 모든 모듈을 일괄 삭제합니다 (리터럴 밑줄은 `\_`로 이스케이프 처리됩니다):
DELETE FROM system.webassembly_modules WHERE name LIKE 'tmp\_%';
