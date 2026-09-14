SELECT  _path, _file, trip_id
FROM s3('https://datasets-documentation.s3.eu-west-3.amazonaws.com/nyc-taxi/trips_0.gz', NOSIGN, 'TabSeparatedWithNames')
LIMIT 5;
