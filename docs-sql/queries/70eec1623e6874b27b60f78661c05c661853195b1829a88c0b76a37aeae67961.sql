SELECT
    date,
    toDate(date)
FROM file('data.avro', Avro)
LIMIT 3;
