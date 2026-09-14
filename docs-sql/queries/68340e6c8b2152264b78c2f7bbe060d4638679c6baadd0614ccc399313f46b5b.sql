INSERT INTO trips
   SELECT *
   FROM s3('https://datasets-documentation.s3.eu-west-3.amazonaws.com/nyc-taxi/trips_*.gz', NOSIGN, 'TabSeparatedWithNames')
   LIMIT 1000000;
