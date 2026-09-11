SELECT *
FROM jdbc('redshift', 'select count(*) from sales')
