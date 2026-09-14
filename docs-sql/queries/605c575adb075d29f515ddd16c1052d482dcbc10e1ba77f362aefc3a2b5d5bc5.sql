-- A bitmap is a binary value, so it is shown with `hex`.
SELECT hex(bitmapBuild([1, 2, 3, 4, 5])) AS res, toTypeName(bitmapBuild([1, 2, 3, 4, 5])) AS type;
