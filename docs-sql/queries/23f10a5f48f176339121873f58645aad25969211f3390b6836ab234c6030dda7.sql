  -- Créer une table avec des parts compactes
  CREATE TABLE compact (
    number UInt32
  )
  ENGINE = MergeTree()
  ORDER BY number 
  AS SELECT * FROM numbers(100000); -- Pas assez volumineux pour dépasser la valeur par défaut de min_bytes_for_wide_part = 10485760

  -- Vérifier le type des parts
  SELECT table, name, part_type from system.parts where table = 'compact';

  -- Obtenir les tailles compressée et non compressée des colonnes pour la table compacte
  SELECT name,
     formatReadableSize(sum(data_compressed_bytes)) AS compressed_size,
     formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed_size,
     round(sum(data_uncompressed_bytes) / sum(data_compressed_bytes), 2) AS ratio
  FROM system.columns
  WHERE table = 'compact'
  GROUP BY name;

  -- Créer une table avec des parts wide 
  CREATE TABLE wide (
    number UInt32
  )
  ENGINE = MergeTree()
  ORDER BY number
  SETTINGS min_bytes_for_wide_part=0
  AS SELECT * FROM numbers(100000);

  -- Vérifier le type des parts
  SELECT table, name, part_type from system.parts where table = 'wide';

  -- Obtenir les tailles compressée et non compressée pour la table wide
  SELECT name,
     formatReadableSize(sum(data_compressed_bytes)) AS compressed_size,
     formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed_size,
     round(sum(data_uncompressed_bytes) / sum(data_compressed_bytes), 2) AS ratio
  FROM system.columns
  WHERE table = 'wide'
  GROUP BY name;
