SYSTEM RESTORE REPLICA table_name;  -- 모든 파티션을 분리하고, ZK에서 메타데이터를 재생성한 후(새로운 빈 테이블처럼), 모든 파티션을 다시 연결합니다

SYSTEM SYNC REPLICA table_name; -- 레플리카가 파트를 동기화할 때까지 기다립니다. 복원 완료 후 모든 레플리카에서 `system.detached_parts`를 확인하는 것을 권장합니다.
