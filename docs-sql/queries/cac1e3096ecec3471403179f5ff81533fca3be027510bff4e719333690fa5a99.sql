    INSERT INTO sensors
    SELECT sensor_id, lat, lon, timestamp, temperature
    FROM azureBlobStorage(
        '<YOUR CONNECTION STRING>', 
        'sensors',
        '2019-06_bmp180.csv.zst', 
        'CSVWithNames')
    SETTINGS format_csv_delimiter = ';'
