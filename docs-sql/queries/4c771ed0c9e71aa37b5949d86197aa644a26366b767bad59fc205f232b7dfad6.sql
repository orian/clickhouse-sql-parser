    CREATE DICTIONARY taxi_zone_dictionary
    (
      `LocationID` UInt16 DEFAULT 0,
      `Borough` String,
      `Zone` String,
      `service_zone` String
    )
    PRIMARY KEY LocationID
    SOURCE(HTTP(URL 'https://datasets-documentation.s3.eu-west-3.amazonaws.com/nyc-taxi/taxi_zone_lookup.csv' FORMAT 'CSVWithNames'))
    LIFETIME(MIN 0 MAX 0)
    LAYOUT(HASHED_ARRAY())
