    # 创建一个包含 3 列的表
    CREATE TABLE users_imported
    (
       `username` String,
       `firstname` String,
       `lastname` String
    )
    ENGINE = MergeTree
    ORDER BY firstname
