-- Принудительно завершает все запросы с указанным query_id:
KILL QUERY WHERE query_id='2-857d-4a57-9ee0-327da5d60a90'

-- Синхронно завершает все запросы, запущенные пользователем 'username':
KILL QUERY WHERE user='username' SYNC
