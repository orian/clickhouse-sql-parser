CREATE OR REPLACE TABLE atp_matches_1960s ORDER BY tourney_id AS
SELECT tourney_id, surface, winner_name, loser_name, winner_seed, loser_seed, score
FROM url('https://raw.githubusercontent.com/JeffSackmann/tennis_atp/refs/heads/master/atp_matches_{1968..1969}.csv')
SETTINGS schema_inference_make_columns_nullable=0, 
         schema_inference_hints='winner_seed Nullable(String), loser_seed Nullable(UInt8)';

CREATE OR REPLACE TABLE atp_matches_1970s ORDER BY tourney_id AS 
SELECT tourney_id, surface, winner_name, loser_name, winner_seed, loser_seed, splitByWhitespace(score) AS score
FROM url('https://raw.githubusercontent.com/JeffSackmann/tennis_atp/refs/heads/master/atp_matches_{1970..1979}.csv')
SETTINGS schema_inference_make_columns_nullable=0, 
         schema_inference_hints='winner_seed Nullable(UInt8), loser_seed Nullable(UInt8)';

CREATE OR REPLACE TABLE atp_matches_1980s ORDER BY tourney_id AS
SELECT tourney_id, surface, winner_name, loser_name, winner_seed, loser_seed, splitByWhitespace(score) AS score
FROM url('https://raw.githubusercontent.com/JeffSackmann/tennis_atp/refs/heads/master/atp_matches_{1980..1989}.csv')
SETTINGS schema_inference_make_columns_nullable=0,
         schema_inference_hints='winner_seed Nullable(UInt16), loser_seed Nullable(UInt16)';

CREATE OR REPLACE TABLE atp_matches_1990s ORDER BY tourney_id AS
SELECT tourney_id, surface, winner_name, loser_name, winner_seed, loser_seed, splitByWhitespace(score) AS score,
       toBool(arrayExists(x -> position(x, 'W/O') > 0, score))::Nullable(bool) AS walkover,
       toBool(arrayExists(x -> position(x, 'RET') > 0, score))::Nullable(bool) AS retirement
FROM url('https://raw.githubusercontent.com/JeffSackmann/tennis_atp/refs/heads/master/atp_matches_{1990..1999}.csv')
SETTINGS schema_inference_make_columns_nullable=0,
         schema_inference_hints='winner_seed Nullable(UInt16), loser_seed Nullable(UInt16), surface Enum(\'Hard\', \'Grass\', \'Clay\', \'Carpet\')';
