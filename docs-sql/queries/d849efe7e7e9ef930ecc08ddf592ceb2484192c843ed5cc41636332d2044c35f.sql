-- 먼저, 데이터를 저장하는 데 사용된 데이터베이스를 확인합니다.
SELECT * FROM [database].connect_state

-- 해당 토픽과 파티션에 해당하는 키를 확인합니다.
ALTER TABLE [database].connect_state DELETE WHERE key = [keyname]
