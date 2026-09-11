-- (1) インサートを一時停止する
-- (2) ターゲットテーブルの複製を作成する

CREATE TABLE pypi_v2 AS pypi

SELECT count() FROM pypi_v2
