-- usando el nombre de la partición
ALTER TABLE mt DELETE IN PARTITION 2 WHERE p = 2;

-- usando el id de la partición
ALTER TABLE mt DELETE IN PARTITION ID '2' WHERE p = 2;
