       -- Accès à la base de données cible : CREATE TABLE, DROP TABLE, SELECT, INSERT, TRUNCATE, etc.
       GRANT ALL ON <database>.* TO <user>;

       -- Accès aux tables système pour la découverte des métadonnées et la gestion des partitions.
       -- Ces autorisations ne sont PAS couvertes par l’autorisation accordée sur la base de données ci-dessus.
       GRANT SELECT ON system.columns TO <user>;
       GRANT SELECT ON system.parts TO <user>;
       GRANT SELECT ON system.tables TO <user>;
