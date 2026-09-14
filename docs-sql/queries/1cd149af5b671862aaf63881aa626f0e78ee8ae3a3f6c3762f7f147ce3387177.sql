{{ config(
       materialized='dictionary',
       fields=[
           ('LocationID', 'UInt16 DEFAULT 0'),
           ('Borough', 'String'),
           ('Zone', 'String'),
       ],
       primary_key='LocationID',
       layout='HASHED()',
       lifetime='MIN 0 MAX 0',
       source_type='http',
       url='https://datasets-documentation.s3.eu-west-3.amazonaws.com/nyc-taxi/taxi_zone_lookup.csv',
       format='CSVWithNames'
) }}

select 1
