       -- Acceso a la base de datos de destino: CREATE TABLE, DROP TABLE, SELECT, INSERT, TRUNCATE, etc.
       GRANT ALL ON <database>.* TO <user>;

       -- Acceso a tablas del sistema para detectar metadatos y gestionar particiones.
       -- Estos NO están incluidos en la concesión para la base de datos anterior.
       GRANT SELECT ON system.columns TO <user>;
       GRANT SELECT ON system.parts TO <user>;
       GRANT SELECT ON system.tables TO <user>;
