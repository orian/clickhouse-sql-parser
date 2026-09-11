SHOW COLUMNS IN system.settings

SELECT * EXCEPT (default, alias_for, readonly, description)
FROM system.settings
LIMIT 5
