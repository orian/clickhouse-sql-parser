SET allow_insert_into_iceberg = 1;

-- 데이터를 삽입하여 스냅샷 생성
INSERT INTO iceberg_table VALUES (1);
INSERT INTO iceberg_table VALUES (2);
INSERT INTO iceberg_table VALUES (3);

-- 보존 정책만 사용하여 만료
ALTER TABLE iceberg_table EXECUTE expire_snapshots();

-- 안전 퓨즈 사용: 타임스탬프보다 최신인 스냅샷 보호 (위치 인수 구문)
ALTER TABLE iceberg_table EXECUTE expire_snapshots('2025-01-01 00:00:00');

-- 명명된 인수 형식으로 동일하게 실행
ALTER TABLE iceberg_table EXECUTE expire_snapshots(expire_before = '2025-01-01 00:00:00');

-- 단일 실행에 대해 보존 매개변수 재정의
ALTER TABLE iceberg_table EXECUTE expire_snapshots(retention_period = '3d', retain_last = 10);

-- 명시적 스냅샷 만료
ALTER TABLE iceberg_table EXECUTE expire_snapshots(snapshot_ids = [101, 102, 103]);

-- 드라이런 미리보기 (메타데이터 업데이트 없음, 파일 삭제 없음)
ALTER TABLE iceberg_table EXECUTE expire_snapshots(retention_period = '1d', dry_run = 1);
