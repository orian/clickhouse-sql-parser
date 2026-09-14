SELECT *
FROM s3(
   'https://datasets-documentation.s3.eu-west-3.amazonaws.com/aapl_stock.csv',
   NOSIGN,
   'CSVWithNames'
)
LIMIT 5;
