-- 指定されたquery_idを持つすべてのクエリを強制終了します:
KILL QUERY WHERE query_id='2-857d-4a57-9ee0-327da5d60a90'

-- 'username'が実行しているすべてのクエリを同期的に終了します:
KILL QUERY WHERE user='username' SYNC
