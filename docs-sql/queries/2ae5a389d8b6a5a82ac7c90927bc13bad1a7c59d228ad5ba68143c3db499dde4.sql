CREATE TABLE [IF NOT EXISTS] [db2.]table_clone AS [db.]table [ENGINE = engine];
ALTER TABLE [db2.]table_clone ATTACH PARTITION ALL FROM [db.]table;
