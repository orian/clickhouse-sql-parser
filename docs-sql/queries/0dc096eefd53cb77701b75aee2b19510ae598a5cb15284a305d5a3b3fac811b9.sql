-- 특정 타임스탬프보다 오래된 고아 파일 제거
ALTER TABLE iceberg_table EXECUTE remove_orphan_files('2026-03-01 00:00:00');

-- 드라이 런: 삭제될 파일 미리 보기
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(dry_run = 1);

-- 데이터 디렉터리만 스캔
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    older_than = '2026-03-01 00:00:00',
    location = 'data/'
);

-- 위치 인수 older_than과 명명된 인수 조합
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    '2026-03-01 00:00:00',
    dry_run = 1
);
