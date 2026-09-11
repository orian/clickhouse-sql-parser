-- Rejected over the native protocol: the client sends the data separately
INSERT INTO target_table FORMAT TSV
INSERT INTO target_table SELECT * FROM input('n UInt64') FORMAT TSV

-- Accepted: the server produces the data itself
INSERT INTO target_table SELECT number FROM numbers(1000000)
