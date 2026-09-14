CREATE TABLE test (d Dynamic) ENGINE=Memory;
INSERT INTO test VALUES (42), (43), ('abc'), ('abd'), ([1, 2, 3]), ([]), (NULL);
