SELECT *
FROM s3(
   'https://s3-us-west-1.amazonaws.com/umbrella-static/top-1m-2018-01-1{0..2}.csv.zip :: *.csv',
   NOSIGN
);
