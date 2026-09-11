SET allow_insert_into_iceberg = 1;

INSERT INTO uk.uk_iceberg SELECT *
FROM uk.uk_price_paid
WHERE town = 'LONDON'
