       -- Acesso ao banco de dados de destino: CREATE TABLE, DROP TABLE, SELECT, INSERT, TRUNCATE etc.
       GRANT ALL ON <database>.* TO <user>;

       -- Acesso às tabelas do sistema para descoberta de metadados e gerenciamento de partições.
       -- Essas permissões NÃO são abrangidas pela concessão ao banco de dados acima.
       GRANT SELECT ON system.columns TO <user>;
       GRANT SELECT ON system.parts TO <user>;
       GRANT SELECT ON system.tables TO <user>;
