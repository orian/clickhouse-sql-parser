-- Insert via text columns:
TRUNCATE texts;
INSERT INTO texts
SELECT n, sha224(bytea('val'||n)), decode(md5('int'||n), 'hex')
  FROM generate_series(1, 4) n;

-- View the data.
SELECT c1, encode(c2::bytea, 'hex'), encode(c3::bytea, 'hex') FROM texts ORDER BY c1;
