SELECT conname, conrelid::regclass, contype
FROM pg_constraint
WHERE conname = 'your_constraint_name';
