-- 위치 기반 형식: 이름 없는 단일 older_than 인수
ALTER TABLE iceberg_table EXECUTE remove_orphan_files('timestamp')

-- 명명된 형식
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    older_than = 'timestamp',
    location = 'path',
    dry_run = 0|1
)

-- 인수 없음: 모든 기본값 사용 (older_than = 3일 전)
ALTER TABLE iceberg_table EXECUTE remove_orphan_files()
