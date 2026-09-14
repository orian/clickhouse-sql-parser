    -- Otorga permisos para leer información del esquema
    GRANT SELECT ON information_schema.* TO your_databrain_user;

    -- Otorga acceso de lectura a tu base de datos y tablas
    GRANT SELECT ON your_database.* TO your_databrain_user;
