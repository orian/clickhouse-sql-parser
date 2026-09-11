    # CRIAÇÃO DA TABELA com 3 colunas
    CREATE TABLE users_imported
    (
       `username` String,
       `firstname` String,
       `lastname` String
    )
    ENGINE = MergeTree
    ORDER BY firstname
