    INSERT INTO users_imported (*) SELECT *
    FROM jdbc('redshift', 'select username, firstname, lastname from users')
