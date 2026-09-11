  -- Crie uma tabela com partes compact
  CREATE TABLE compact (
    number UInt32
  )
  ENGINE = MergeTree()
  ORDER BY number 
  AS SELECT * FROM numbers(100000); -- Não é grande o suficiente para exceder o valor padrão de min_bytes_for_wide_part = 10485760

  -- Verifique o tipo das partes
  SELECT table, name, part_type from system.parts where table = 'compact';

  -- Obtenha os tamanhos comprimido e não comprimido das colunas da tabela compact
  SELECT name,
     formatReadableSize(sum(data_compressed_bytes)) AS compressed_size,
     formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed_size,
     round(sum(data_uncompressed_bytes) / sum(data_compressed_bytes), 2) AS ratio
  FROM system.columns
  WHERE table = 'compact'
  GROUP BY name;

  -- Crie uma tabela com partes wide 
  CREATE TABLE wide (
    number UInt32
  )
  ENGINE = MergeTree()
  ORDER BY number
  SETTINGS min_bytes_for_wide_part=0
  AS SELECT * FROM numbers(100000);

  -- Verifique o tipo das partes
  SELECT table, name, part_type from system.parts where table = 'wide';

  -- Obtenha os tamanhos comprimido e não comprimido da tabela wide
  SELECT name,
     formatReadableSize(sum(data_compressed_bytes)) AS compressed_size,
     formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed_size,
     round(sum(data_uncompressed_bytes) / sum(data_compressed_bytes), 2) AS ratio
  FROM system.columns
  WHERE table = 'wide'
  GROUP BY name;
