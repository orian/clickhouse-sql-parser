SELECT *
FROM system.scheduler
WHERE resource = 'network_read' AND path = '/prio/fair/prod'
FORMAT Vertical
