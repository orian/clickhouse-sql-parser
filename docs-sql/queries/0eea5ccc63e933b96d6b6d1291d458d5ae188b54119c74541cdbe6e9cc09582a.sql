-- 지정된 query_id를 가진 모든 쿼리를 강제로 종료합니다:
KILL QUERY WHERE query_id='2-857d-4a57-9ee0-327da5d60a90'

-- 'username'이 실행한 모든 쿼리를 동기적으로 종료합니다:
KILL QUERY WHERE user='username' SYNC
