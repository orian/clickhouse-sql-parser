
        SELECT
            sum({func}(vec, arrayMap(i -> to{type}(i + 1), range(64)), {p}))
          + sum({func}(vec, arrayMap(i -> to{type}(i + 2), range(64)), {p}))
          + sum({func}(vec, arrayMap(i -> to{type}(i + 3), range(64)), {p}))
          + sum({func}(vec, arrayMap(i -> to{type}(i + 4), range(64)), {p}))
          + sum({func}(vec, arrayMap(i -> to{type}(i + 5), range(64)), {p}))
          + sum({func}(vec, arrayMap(i -> to{type}(i + 6), range(64)), {p}))
          + sum({func}(vec, arrayMap(i -> to{type}(i + 7), range(64)), {p}))
          + sum({func}(vec, arrayMap(i -> to{type}(i + 8), range(64)), {p}))
        FROM tab_{type}_wide
    