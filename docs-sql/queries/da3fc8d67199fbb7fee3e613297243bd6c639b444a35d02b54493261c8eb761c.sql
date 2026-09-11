-- usando nome da partição
ALTER TABLE mt DELETE IN PARTITION 2 WHERE p = 2;

-- usando id da partição
ALTER TABLE mt DELETE IN PARTITION ID '2' WHERE p = 2;
