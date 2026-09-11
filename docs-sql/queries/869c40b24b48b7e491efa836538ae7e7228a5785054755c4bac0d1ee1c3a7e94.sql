    SELECT *
    FROM file('hacknernews.csv.gz', CSVWithNames)
    LIMIT 2
    SETTINGS input_format_try_infer_datetimes = 0
    FORMAT Vertical
