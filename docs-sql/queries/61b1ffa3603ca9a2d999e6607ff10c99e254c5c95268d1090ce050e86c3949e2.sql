SELECT
    IPv4NumToStringClassC(ClientIP) AS k,
    count() AS c
FROM VALUES('ClientIP UInt32',
    (toUInt32(toIPv4('83.149.9.10'))), (toUInt32(toIPv4('83.149.9.11'))), (toUInt32(toIPv4('83.149.9.12'))),
    (toUInt32(toIPv4('217.118.81.5'))), (toUInt32(toIPv4('217.118.81.6'))),
    (toUInt32(toIPv4('213.87.129.1'))))
GROUP BY k
ORDER BY c DESC, k
