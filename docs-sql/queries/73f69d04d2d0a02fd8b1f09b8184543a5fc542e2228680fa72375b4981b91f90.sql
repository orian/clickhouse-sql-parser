      SET enable_json_type = 1;

      CREATE TABLE logs_system_syslog
      (
          `@timestamp` DateTime,
          `agent` Tuple(
              ephemeral_id String,
              id String,
              name String,
              type String,
              version String),
          `cloud` Tuple(
              account Tuple(
                  id String),
              availability_zone String,
              image Tuple(
                  id String),
              instance Tuple(
                  id String),
              machine Tuple(
                  type String),
              provider String,
              region String,
              service Tuple(
                  name String)),
          `data_stream` Tuple(
              dataset String,
              namespace String,
              type String),
          `ecs` Tuple(
              version String),
          `elastic_agent` Tuple(
              id String,
              snapshot UInt8,
              version String),
          `event` Tuple(
              agent_id_status String,
              dataset String,
              ingested DateTime,
              module String,
              timezone String),
          `host` Tuple(
              architecture String,
              containerized UInt8,
              hostname String,
              id String,
              ip Array(Variant(IPv4, IPv6)),
              mac Array(String),
              name String,
              os Tuple(
                  build String,
                  codename String,
                  family String,
                  kernel String,
                  name String,
                  platform String,
                  type String,
                  version String)),
          `input` Tuple(
              type String),
          `log` Tuple(
              file Tuple(
                  path String),
              offset Int64),
          `message` String,
          `process` Tuple(
              name String,
              pid Int64),
          `system` Tuple(
              syslog JSON)
      )
      ENGINE = MergeTree
      ORDER BY (`host.name`, `@timestamp`)
