-- First, identify the database used to store the data.
SELECT * FROM [database].connect_state

-- Identify the key that matches the topic and partition.
ALTER TABLE [database].connect_state DELETE WHERE key = [keyname]
