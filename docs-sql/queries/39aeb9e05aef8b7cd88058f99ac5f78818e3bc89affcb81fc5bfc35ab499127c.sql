-- In the example below the `countMatches` function expects a constant second argument.
-- This behaviour can be debugged by using the `materialize` function to turn a constant into a full column,
-- verifying that the function throws an error for a non-constant argument.

SELECT countMatches('foobarfoo', 'foo');
