-- Shard 1: consumes partitions 0, 3, 6, 9
CREATE TABLE kafka_shard1 (key UInt64, value String)
ENGINE = Kafka('localhost:9092', 'my-topic', 'my-group', 'JSONEachRow')
SETTINGS
    kafka_keeper_path = '/clickhouse/kafka/{database}',
    kafka_replica_name = '{replica}',
    kafka_partition_shard_num = '1',
    kafka_shard_count = 3
SETTINGS allow_experimental_kafka_offsets_storage_in_keeper = 1;

-- Shard 2: consumes partitions 1, 4, 7, 10
CREATE TABLE kafka_shard2 (key UInt64, value String)
ENGINE = Kafka('localhost:9092', 'my-topic', 'my-group', 'JSONEachRow')
SETTINGS
    kafka_keeper_path = '/clickhouse/kafka/{database}',
    kafka_replica_name = '{replica}',
    kafka_partition_shard_num = '2',
    kafka_shard_count = 3
SETTINGS allow_experimental_kafka_offsets_storage_in_keeper = 1;
