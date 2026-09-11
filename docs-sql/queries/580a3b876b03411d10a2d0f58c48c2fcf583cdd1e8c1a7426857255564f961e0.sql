-- Primero, identifica la base de datos utilizada para almacenar los datos.
SELECT * FROM [database].connect_state

-- Identifica la clave que coincide con el topic y la partición.
ALTER TABLE [database].connect_state DELETE WHERE key = [keyname]
