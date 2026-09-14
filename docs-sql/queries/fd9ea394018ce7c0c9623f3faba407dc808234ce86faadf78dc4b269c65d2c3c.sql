-- Termina forzosamente todas las consultas con el query_id especificado:
KILL QUERY WHERE query_id='2-857d-4a57-9ee0-327da5d60a90'

-- Termina de forma síncrona todas las consultas ejecutadas por 'username':
KILL QUERY WHERE user='username' SYNC
