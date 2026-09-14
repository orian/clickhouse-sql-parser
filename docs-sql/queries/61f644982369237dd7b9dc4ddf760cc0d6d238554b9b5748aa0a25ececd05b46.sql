    -- Accorder les autorisations pour lire les informations de schéma
    GRANT SELECT ON information_schema.* TO your_databrain_user;

    -- Accorder un accès en lecture à votre base de données et à vos tables
    GRANT SELECT ON your_database.* TO your_databrain_user;
