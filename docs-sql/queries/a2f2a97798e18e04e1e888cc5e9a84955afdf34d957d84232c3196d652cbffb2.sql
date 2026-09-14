    SELECT *
    FROM jdbc('redshift', 'select username, firstname, lastname from users limit 5')
