taxi=# \d taxi.trips
                                   Foreign table "taxi.trips"
        Column         |           Type           | Collation | Nullable | Default | FDW options
-----------------------+--------------------------+-----------+----------+---------+-------------
 trip_id               | bigint                   |           | not null |         |
 vendor_id             | text                     |           | not null |         |
 pickup_date           | date                     |           | not null |         |
 pickup_datetime       | timestamp with time zone |           | not null |         |
 dropoff_date          | date                     |           | not null |         |
 dropoff_datetime      | timestamp with time zone |           | not null |         |
 store_and_fwd_flag    | smallint                 |           | not null |         |
 rate_code_id          | smallint                 |           | not null |         |
 pickup_longitude      | double precision         |           | not null |         |
 pickup_latitude       | double precision         |           | not null |         |
 dropoff_longitude     | double precision         |           | not null |         |
 dropoff_latitude      | double precision         |           | not null |         |
 passenger_count       | smallint                 |           | not null |         |
 trip_distance         | double precision         |           | not null |         |
 fare_amount           | numeric(10,2)            |           | not null |         |
 extra                 | numeric(10,2)            |           | not null |         |
 mta_tax               | numeric(10,2)            |           | not null |         |
 tip_amount            | numeric(10,2)            |           | not null |         |
 tolls_amount          | numeric(10,2)            |           | not null |         |
 ehail_fee             | numeric(10,2)            |           | not null |         |
 improvement_surcharge | numeric(10,2)            |           | not null |         |
 total_amount          | numeric(10,2)            |           | not null |         |
 payment_type          | text                     |           | not null |         |
 trip_type             | smallint                 |           | not null |         |
 pickup                | character varying(25)    |           | not null |         |
 dropoff               | character varying(25)    |           | not null |         |
 cab_type              | text                     |           | not null |         |
 pickup_nyct2010_gid   | smallint                 |           | not null |         |
 pickup_ctlabel        | real                     |           | not null |         |
 pickup_borocode       | smallint                 |           | not null |         |
 pickup_ct2010         | text                     |           | not null |         |
 pickup_boroct2010     | text                     |           | not null |         |
 pickup_cdeligibil     | text                     |           | not null |         |
 pickup_ntacode        | character varying(4)     |           | not null |         |
 pickup_ntaname        | text                     |           | not null |         |
 pickup_puma           | integer                  |           | not null |         |
 dropoff_nyct2010_gid  | smallint                 |           | not null |         |
 dropoff_ctlabel       | real                     |           | not null |         |
 dropoff_borocode      | smallint                 |           | not null |         |
 dropoff_ct2010        | text                     |           | not null |         |
 dropoff_boroct2010    | text                     |           | not null |         |
 dropoff_cdeligibil    | text                     |           | not null |         |
 dropoff_ntacode       | character varying(4)     |           | not null |         |
 dropoff_ntaname       | text                     |           | not null |         |
 dropoff_puma          | integer                  |           | not null |         |
Server: taxi_srv
FDW options: (database 'taxi', table_name 'trips', engine 'MergeTree')
