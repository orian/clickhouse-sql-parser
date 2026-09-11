SELECT
    datacenter, 
    SUM (quantity) qty
FROM
    servers
GROUP BY
    datacenter;
