    -- Conceda permissões para ler informações de schema
    GRANT SELECT ON information_schema.* TO your_databrain_user;

    -- Conceda acesso de leitura ao seu banco de dados e às suas tabelas
    GRANT SELECT ON your_database.* TO your_databrain_user;
