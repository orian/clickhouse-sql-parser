SELECT _table, loser_name, score, winner_seed
FROM merge('atp_matches*')
WHERE winner_name = 'John McEnroe'
AND loser_seed = 1
AND multiIf(
  variantType(winner_seed) = 'UInt8', variantElement(winner_seed, 'UInt8') >= 3,
  variantType(winner_seed) = 'UInt16', variantElement(winner_seed, 'UInt16') >= 3,
  variantElement(winner_seed, 'String')::UInt16 >= 3
);
