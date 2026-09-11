-- A bitmap is a binary value, so it is shown with `hex`.
SELECT hex(bitmapBuild([-128, -1]::Array(Int8))) AS res, toTypeName(bitmapBuild([-128, -1]::Array(Int8))) AS type;
