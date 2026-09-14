-- Удалить осиротевшие файлы старше указанной временной метки
ALTER TABLE iceberg_table EXECUTE remove_orphan_files('2026-03-01 00:00:00');

-- Пробный запуск: просмотр файлов, которые будут удалены
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(dry_run = 1);

-- Сканировать только каталог данных
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    older_than = '2026-03-01 00:00:00',
    location = 'data/'
);

-- Совместное использование позиционного аргумента older_than с именованными аргументами
ALTER TABLE iceberg_table EXECUTE remove_orphan_files(
    '2026-03-01 00:00:00',
    dry_run = 1
);
