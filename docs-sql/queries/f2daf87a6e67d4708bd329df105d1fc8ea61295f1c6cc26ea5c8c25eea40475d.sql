SELECT
    sum(LO_REVENUE),
    D_YEAR,
    P_BRAND
FROM
    lineorder,
    date,
    part,
    supplier
WHERE
    LO_ORDERDATE = D_DATEKEY
    AND LO_PARTKEY = P_PARTKEY
    AND LO_SUPPKEY = S_SUPPKEY
    AND P_BRAND = 'MFGR#2221'
    AND S_REGION = 'EUROPE'
GROUP BY
    D_YEAR,
    P_BRAND
ORDER BY
    D_YEAR,
    P_BRAND;
