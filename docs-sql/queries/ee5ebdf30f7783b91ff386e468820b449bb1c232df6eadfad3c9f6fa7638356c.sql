DELETE FROM system.webassembly_modules WHERE name = 'collatz';

-- 名前が `tmp_` で始まるすべてのモジュールを一括削除する（リテラルのアンダースコアは `\_` としてエスケープされる）:
DELETE FROM system.webassembly_modules WHERE name LIKE 'tmp\_%';
