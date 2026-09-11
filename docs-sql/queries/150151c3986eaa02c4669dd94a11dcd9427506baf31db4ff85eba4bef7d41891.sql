SELECT count(*) FROM file('big_dir/**/file002', 'CSV', 'name String, value UInt32');
