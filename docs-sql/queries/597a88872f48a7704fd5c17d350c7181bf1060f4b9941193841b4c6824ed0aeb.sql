SELECT
    datacenter,
    distro,
    version,
    SUM(quantity)
FROM servers
GROUP BY
    GROUPING SETS (
        (datacenter, distro, version),
        (datacenter, distro))
