SELECT count(*) FROM file('big_dir/**', 'CSV', 'name String, value UInt32');
