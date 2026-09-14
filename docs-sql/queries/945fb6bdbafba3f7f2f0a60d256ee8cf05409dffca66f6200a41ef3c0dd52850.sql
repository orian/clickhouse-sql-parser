  -- أنشئ جدولًا بأجزاء compact
  CREATE TABLE compact (
    number UInt32
  )
  ENGINE = MergeTree()
  ORDER BY number 
  AS SELECT * FROM numbers(100000); -- ليست كبيرة بما يكفي لتجاوز القيمة الافتراضية لـ min_bytes_for_wide_part = 10485760

  -- تحقّق من نوع الأجزاء
  SELECT table, name, part_type from system.parts where table = 'compact';

  -- احصل على الأحجام المضغوطة وغير المضغوطة للأعمدة في الجدول compact
  SELECT name,
     formatReadableSize(sum(data_compressed_bytes)) AS compressed_size,
     formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed_size,
     round(sum(data_uncompressed_bytes) / sum(data_compressed_bytes), 2) AS ratio
  FROM system.columns
  WHERE table = 'compact'
  GROUP BY name;

  -- أنشئ جدولًا بأجزاء wide 
  CREATE TABLE wide (
    number UInt32
  )
  ENGINE = MergeTree()
  ORDER BY number
  SETTINGS min_bytes_for_wide_part=0
  AS SELECT * FROM numbers(100000);

  -- تحقّق من نوع الأجزاء
  SELECT table, name, part_type from system.parts where table = 'wide';

  -- احصل على الأحجام المضغوطة وغير المضغوطة للجدول wide
  SELECT name,
     formatReadableSize(sum(data_compressed_bytes)) AS compressed_size,
     formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed_size,
     round(sum(data_uncompressed_bytes) / sum(data_compressed_bytes), 2) AS ratio
  FROM system.columns
  WHERE table = 'wide'
  GROUP BY name;
