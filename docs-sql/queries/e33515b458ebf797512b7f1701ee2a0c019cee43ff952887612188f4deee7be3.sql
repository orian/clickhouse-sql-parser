DETACH TABLE table_name;  -- DROP REPLICA 실행 전 필요

SYSTEM DROP REPLICA 'replica_name' FROM ZKPATH '/table_path_in_zk/'; -- /table_path_in_zk의 모든 내용을 제거합니다

ATTACH TABLE table_name;  -- ZK에 메타데이터가 없으므로 테이블이 읽기 전용 모드로 전환됩니다
