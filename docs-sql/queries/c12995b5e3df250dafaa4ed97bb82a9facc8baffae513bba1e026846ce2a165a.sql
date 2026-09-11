CREATE TABLE tab (a Map(String, String))
ENGINE = MergeTree
ORDER BY tuple();

INSERT INTO tab VALUES ({'abc':'abc','def':'def'}), ({'hij':'hij','klm':'klm'});

SELECT mapContainsValueLike(a, 'a%') FROM tab;
