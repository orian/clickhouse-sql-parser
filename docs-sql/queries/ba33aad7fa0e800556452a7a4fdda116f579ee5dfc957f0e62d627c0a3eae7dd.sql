-- 强制终止所有具有指定 query_id 的查询：
KILL QUERY WHERE query_id='2-857d-4a57-9ee0-327da5d60a90'

-- 同步终止 'username' 运行的所有查询：
KILL QUERY WHERE user='username' SYNC
