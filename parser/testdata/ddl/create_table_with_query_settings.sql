CREATE TABLE d.kafka_t (id UInt64, n Nested(a UInt8, b String)) ENGINE = Kafka SETTINGS kafka_broker_list = 'b', kafka_topic_list = 't', kafka_group_name = 'g', kafka_format = 'JSONEachRow' SETTINGS flatten_nested = 0;
CREATE TABLE d.merge_t (id UInt64) ENGINE = MergeTree ORDER BY id SETTINGS index_granularity = 8192 SETTINGS flatten_nested = 0;
CREATE TABLE d.memory_t (id UInt64) ENGINE = Memory COMMENT 'comment' SETTINGS max_threads = 4;
