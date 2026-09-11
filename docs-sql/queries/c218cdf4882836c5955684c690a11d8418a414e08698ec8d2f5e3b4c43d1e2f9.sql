    DECLARE export_path STRING;
    DECLARE n INT64;
    DECLARE i INT64;
    SET i = 0;

    -- We recommend setting n to correspond to x billion rows. So 5 billion rows, n = 5
    SET n = 100;

    WHILE i < n DO
      SET export_path = CONCAT('gs://mybucket/mytable/', i,'-*.parquet');
      EXPORT DATA
        OPTIONS (
          uri = export_path,
          format = 'PARQUET',
          overwrite = true
        )
      AS (
        SELECT * FROM mytable WHERE export_id = i
      );
      SET i = i + 1;
    END WHILE;
