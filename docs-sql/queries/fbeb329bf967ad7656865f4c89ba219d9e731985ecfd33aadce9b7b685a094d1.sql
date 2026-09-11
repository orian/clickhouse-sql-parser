    -- Использование HTTPS URL (рекомендуется)
    CREATE TABLE table_name
    ENGINE = DeltaLake('https://storage.googleapis.com/<bucket>/<path>/', '<access_key_id>', '<secret_access_key>')
