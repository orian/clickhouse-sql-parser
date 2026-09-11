  -- compact 파트를 사용하는 테이블 생성
  CREATE TABLE compact (
    number UInt32
  )
  ENGINE = MergeTree()
  ORDER BY number 
  AS SELECT * FROM numbers(100000); -- 기본값 min_bytes_for_wide_part = 10485760을 초과할 만큼 크지 않음

  -- 파트 유형 확인
  SELECT table, name, part_type from system.parts where table = 'compact';

  -- compact 테이블의 압축된 크기 및 압축되지 않은 크기 가져오기
  SELECT name,
     formatReadableSize(sum(data_compressed_bytes)) AS compressed_size,
     formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed_size,
     round(sum(data_uncompressed_bytes) / sum(data_compressed_bytes), 2) AS ratio
  FROM system.columns
  WHERE table = 'compact'
  GROUP BY name;

  -- wide 파트를 사용하는 테이블 생성 
  CREATE TABLE wide (
    number UInt32
  )
  ENGINE = MergeTree()
  ORDER BY number
  SETTINGS min_bytes_for_wide_part=0
  AS SELECT * FROM numbers(100000);

  -- 파트 유형 확인
  SELECT table, name, part_type from system.parts where table = 'wide';

  -- wide 테이블의 압축된 크기 및 압축되지 않은 크기 가져오기
  SELECT name,
     formatReadableSize(sum(data_compressed_bytes)) AS compressed_size,
     formatReadableSize(sum(data_uncompressed_bytes)) AS uncompressed_size,
     round(sum(data_uncompressed_bytes) / sum(data_compressed_bytes), 2) AS ratio
  FROM system.columns
  WHERE table = 'wide'
  GROUP BY name;
