-- Encerra forçosamente todas as consultas com o query_id especificado:
KILL QUERY WHERE query_id='2-857d-4a57-9ee0-327da5d60a90'

-- Encerra de forma síncrona todas as consultas executadas pelo 'username':
KILL QUERY WHERE user='username' SYNC
