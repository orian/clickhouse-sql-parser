-- Origin SQL:
CREATE DICTIONARY test.flat_dict (
    id UInt64,
    value String
)
PRIMARY KEY id
SOURCE(CLICKHOUSE(
    host 'localhost'
    port 9000
    table 'src'
))
LIFETIME(MIN 0 MAX 1000)
LAYOUT(FLAT);

CREATE DICTIONARY test.ip_dict (
    prefix String,
    asn UInt32
)
PRIMARY KEY prefix
SOURCE(CLICKHOUSE(
    host 'localhost'
    port 9000
    table 'src'
))
LIFETIME(MIN 0 MAX 1000)
LAYOUT(IP_TRIE);

CREATE DICTIONARY test.hashed_dict (
    id UInt64,
    value String
)
PRIMARY KEY id
SOURCE(CLICKHOUSE(
    host 'localhost'
    port 9000
    table 'src'
))
LIFETIME(MIN 0 MAX 1000)
LAYOUT(HASHED());


-- Format SQL:
CREATE DICTIONARY test.flat_dict (id UInt64, value String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'localhost' port 9000 table 'src')) LIFETIME(MIN 0 MAX 1000) LAYOUT(FLAT());
CREATE DICTIONARY test.ip_dict (prefix String, asn UInt32) PRIMARY KEY prefix SOURCE(CLICKHOUSE(host 'localhost' port 9000 table 'src')) LIFETIME(MIN 0 MAX 1000) LAYOUT(IP_TRIE());
CREATE DICTIONARY test.hashed_dict (id UInt64, value String) PRIMARY KEY id SOURCE(CLICKHOUSE(host 'localhost' port 9000 table 'src')) LIFETIME(MIN 0 MAX 1000) LAYOUT(HASHED());
