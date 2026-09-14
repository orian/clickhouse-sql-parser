SELECT
   datacenter,
   distro,
   version,
   SUM(quantity)
FROM
   servers
GROUP BY
   CUBE(datacenter,distro,version)
ORDER BY
   datacenter,
   distro;
