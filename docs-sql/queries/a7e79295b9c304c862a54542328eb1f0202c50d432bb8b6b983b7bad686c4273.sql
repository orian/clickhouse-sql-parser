-- テーブルを作成してデータを挿入する
CREATE TABLE test_lazy (json JSON) ENGINE = MergeTree ORDER BY tuple();
INSERT INTO test_lazy VALUES ('{"user_id": "123", "score": "95.5"}');

-- 実験的な設定を有効にする
SET allow_experimental_json_lazy_type_hints = 1;

-- 型ヒントを追加する - ミューテーションなしで即座に完了する
ALTER TABLE test_lazy MODIFY COLUMN json JSON(user_id UInt64, score Float64);

-- データをクエリする - 型ヒントは読み取り時に適用される
SELECT json.user_id, toTypeName(json.user_id), json.score, toTypeName(json.score) FROM test_lazy;
