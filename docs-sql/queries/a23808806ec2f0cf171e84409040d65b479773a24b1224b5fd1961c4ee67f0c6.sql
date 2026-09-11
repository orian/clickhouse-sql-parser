  -- Crear una tabla con partes compact
  CREATE TABLE compact (
    number UInt32
  )
  ENGINE = MergeTree()
  ORDER BY number 
  AS SELECT * FROM numbers(100000); -- No es lo bastante grande como para superar el valor predeterminado de min_bytes_for_wide_part = 10485760

  -- Comprobar el tipo de las partes
  SELECT table, name, part_type from system.parts where table = 'compact';

  -- Obtener los tamaños comprimido y sin comprimir de las columnas para la tabla compact
  SELECT name,
     formatReadableSize(sum(data_compressed_bytes)) AS compressed_size,
     formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed_size,
     round(sum(data_uncompressed_bytes) / sum(data_compressed_bytes), 2) AS ratio
  FROM system.columns
  WHERE table = 'compact'
  GROUP BY name;

  -- Crear una tabla con partes wide 
  CREATE TABLE wide (
    number UInt32
  )
  ENGINE = MergeTree()
  ORDER BY number
  SETTINGS min_bytes_for_wide_part=0
  AS SELECT * FROM numbers(100000);

  -- Comprobar el tipo de las partes
  SELECT table, name, part_type from system.parts where table = 'wide';

  -- Obtener los tamaños comprimido y sin comprimir para la tabla wide
  SELECT name,
     formatReadableSize(sum(data_compressed_bytes)) AS compressed_size,
     formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed_size,
     round(sum(data_uncompressed_bytes) / sum(data_compressed_bytes), 2) AS ratio
  FROM system.columns
  WHERE table = 'wide'
  GROUP BY name;
