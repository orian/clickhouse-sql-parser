SELECT
    datacenter,
    distro, 
    SUM (quantity) qty
FROM
    servers
GROUP BY
    datacenter,
    distro;
