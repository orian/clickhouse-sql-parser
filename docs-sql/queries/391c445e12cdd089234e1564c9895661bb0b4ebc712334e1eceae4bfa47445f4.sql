-- OK: stays within the Goals structure
ALTER TABLE test.visits RENAME COLUMN Goals.Price TO Goals.Amount;
