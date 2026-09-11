-- Сначала определите базу данных, используемую для хранения данных.
SELECT * FROM [database].connect_state

-- Определите ключ, соответствующий топику и партиции.
ALTER TABLE [database].connect_state DELETE WHERE key = [keyname]
