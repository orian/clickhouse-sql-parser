       -- Доступ к целевой базе данных: CREATE TABLE, DROP TABLE, SELECT, INSERT, TRUNCATE и т. д.
       GRANT ALL ON <database>.* TO <user>;

       -- Доступ к системным таблицам для обнаружения метаданных и управления партициями.
       -- Эти разрешения НЕ предоставляются указанным выше разрешением для базы данных.
       GRANT SELECT ON system.columns TO <user>;
       GRANT SELECT ON system.parts TO <user>;
       GRANT SELECT ON system.tables TO <user>;
