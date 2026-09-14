  -- Создайте таблицу с компактными частями
  CREATE TABLE compact (
    number UInt32
  )
  ENGINE = MergeTree()
  ORDER BY number 
  AS SELECT * FROM numbers(100000); -- Недостаточно большой объём, чтобы превысить значение по умолчанию min_bytes_for_wide_part = 10485760

  -- Проверьте тип частей
  SELECT table, name, part_type from system.parts where table = 'compact';

  -- Получите сжатый и несжатый размеры столбцов для компактной таблицы
  SELECT name,
     formatReadableSize(sum(data_compressed_bytes)) AS compressed_size,
     formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed_size,
     round(sum(data_uncompressed_bytes) / sum(data_compressed_bytes), 2) AS ratio
  FROM system.columns
  WHERE table = 'compact'
  GROUP BY name;

  -- Создайте таблицу с широкими частями 
  CREATE TABLE wide (
    number UInt32
  )
  ENGINE = MergeTree()
  ORDER BY number
  SETTINGS min_bytes_for_wide_part=0
  AS SELECT * FROM numbers(100000);

  -- Проверьте тип частей
  SELECT table, name, part_type from system.parts where table = 'wide';

  -- Получите сжатый и несжатый размеры для широкой таблицы
  SELECT name,
     formatReadableSize(sum(data_compressed_bytes)) AS compressed_size,
     formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed_size,
     round(sum(data_uncompressed_bytes) / sum(data_compressed_bytes), 2) AS ratio
  FROM system.columns
  WHERE table = 'wide'
  GROUP BY name;
