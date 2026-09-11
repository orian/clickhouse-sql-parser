CREATE TABLE id_val(`id` UInt32, `val` UInt32) ENGINE = Join(ANY, LEFT, id);
INSERT INTO id_val VALUES (1,11)(2,12)(4,13);

SELECT joinGet(id_val, 'val', toUInt32(1));
