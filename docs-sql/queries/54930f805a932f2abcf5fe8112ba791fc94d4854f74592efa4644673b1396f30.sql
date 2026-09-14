CREATE TABLE topic1_sink
(
    field1 String,
    field2 String
)
ENGINE = Kafka()
SETTINGS
kafka_broker_list = 'kafka-broker',
kafka_topic_list = 'topic1',
kafka_format = 'AvroConfluent',
format_avro_schema_registry_url = 'http://schema-registry-url',
output_format_avro_confluent_subject = 'topic1-value';

INSERT INTO topic1_sink VALUES ('hello', 'world');
