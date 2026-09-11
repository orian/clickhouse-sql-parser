CREATE TABLE enum
(
    x Enum('One' = 1, 'Two', 'Three')
) ENGINE = Memory;
ALTER TABLE enum MODIFY COLUMN x ADD ENUM VALUES ('Zero' = 0, 'Four' = 4);
SHOW CREATE TABLE enum;
