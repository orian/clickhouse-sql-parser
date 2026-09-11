SELECT loser_name, score
FROM merge('atp_matches*')
WHERE winner_name = 'John McEnroe'
AND loser_seed = 1;
