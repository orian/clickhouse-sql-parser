INSERT INTO events (data)
FROM INFILE 'custom.json'
FORMAT JSONAsString
