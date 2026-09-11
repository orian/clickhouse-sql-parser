-- используя имя партиции
ALTER TABLE mt DELETE IN PARTITION 2 WHERE p = 2;

-- используя идентификатор партиции
ALTER TABLE mt DELETE IN PARTITION ID '2' WHERE p = 2;
