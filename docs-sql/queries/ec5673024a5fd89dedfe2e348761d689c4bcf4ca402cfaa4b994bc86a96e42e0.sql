CREATE TABLE hits_all (EventDate Date, ClientIP6 FixedString(16)) ENGINE = Memory;

INSERT INTO hits_all SELECT today(), IPv6StringToNum(addr) FROM VALUES('addr String',
    ('2a02:2168:aaa:bbbb::2'), ('2a02:2168:aaa:bbbb::2'), ('2a02:2168:aaa:bbbb::2'),
    ('2a02:6b8:0:fff::ff'), ('2a02:6b8:0:fff::ff'),
    ('::ffff:94.26.111.111'), ('::ffff:94.26.111.111'), ('::ffff:94.26.111.111'), ('::ffff:94.26.111.111'),
    ('::ffff:37.143.222.4'));

-- The addresses that are not an IPv4 address mapped into the IPv6 space.
SELECT
    IPv6NumToString(ClientIP6 AS k),
    count() AS c
FROM hits_all
WHERE EventDate = today() AND substring(ClientIP6, 1, 12) != unhex('00000000000000000000FFFF')
GROUP BY k
ORDER BY c DESC, k
LIMIT 10
