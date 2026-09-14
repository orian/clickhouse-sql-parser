    -- Предоставьте разрешения на чтение информации о схеме
    GRANT SELECT ON information_schema.* TO your_databrain_user;

    -- Предоставьте доступ на чтение к вашей базе данных и таблицам
    GRANT SELECT ON your_database.* TO your_databrain_user;
