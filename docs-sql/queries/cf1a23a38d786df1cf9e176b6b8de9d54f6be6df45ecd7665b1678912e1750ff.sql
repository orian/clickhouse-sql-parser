-- Primeiro, identifique o banco de dados usado para armazenar os dados.
SELECT * FROM [database].connect_state

-- Identifique a chave que corresponde ao tópico e à partição.
ALTER TABLE [database].connect_state DELETE WHERE key = [keyname]
