SELECT
    distro, 
    SUM (quantity) qty
FROM
    servers
GROUP BY
    distro;
