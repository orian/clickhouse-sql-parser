CREATE TABLE task_priority_source (task_id UInt64, priority_level String) ENGINE = Memory;
INSERT INTO task_priority_source VALUES (1, 'low'), (2, 'high'), (3, 'medium'), (4, 'high');
CREATE DICTIONARY task_id_to_priority_dictionary (task_id UInt64, priority_level String) PRIMARY KEY task_id SOURCE(CLICKHOUSE(TABLE 'task_priority_source' DB currentDatabase())) LAYOUT(HASHED()) LIFETIME(MIN 0 MAX 0);
SELECT arraySort(dictGetKeys('task_id_to_priority_dictionary', 'priority_level', 'high')) AS ids;
