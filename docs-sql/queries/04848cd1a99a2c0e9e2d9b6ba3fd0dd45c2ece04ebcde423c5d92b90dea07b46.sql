CREATE TABLE queue (
    key UInt64,
    value UInt64,
    date DateTime
  ) ENGINE = RabbitMQ SETTINGS rabbitmq_host_port = 'localhost:5672',
                            rabbitmq_exchange_name = 'exchange1',
                            rabbitmq_format = 'JSONEachRow',
                            rabbitmq_num_consumers = 5,
                            date_time_input_format = 'best_effort';
