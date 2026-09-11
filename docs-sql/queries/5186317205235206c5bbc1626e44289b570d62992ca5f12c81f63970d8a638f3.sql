-- 範囲の拡張は、範囲（またはカーディナリティ）が大きく異なる引数に対して
-- 同様の分布が必要な場合に有効です
-- 例: 'IP Address' (0...FFFFFFFF) と 'Country code' (0...FF)
-- 注意: Tuple のサイズは他の引数の数と等しくなければなりません
SELECT mortonEncode((1,2), 1024, 16)
