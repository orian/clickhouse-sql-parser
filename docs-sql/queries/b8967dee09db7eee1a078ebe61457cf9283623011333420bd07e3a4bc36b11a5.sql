SELECT zookeeper_path, processed_nodes_count, processing_nodes_count, failed_nodes_count
FROM system.s3_queue_metadata

Row 1:
──────
zookeeper_path:         /clickhouse/s3queue/25ea5621-ae8c-40c7-96d0-cec959c5ab88/3b3f66a1-9866-4c2e-ba78-b6bfa154207e
processed_nodes_count:  128
processing_nodes_count: 2
failed_nodes_count:     0
