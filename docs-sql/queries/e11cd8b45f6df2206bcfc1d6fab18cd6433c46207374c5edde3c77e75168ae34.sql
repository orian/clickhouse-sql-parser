SELECT *
FROM file('data.binary', RowBinary, 'path String, month Date, hits UInt32')
LIMIT 5
