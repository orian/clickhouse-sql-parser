
        SELECT
            sum({func}(vec, arrayMap(i -> to{type}(i + 1), range({vector_size})), {p}))
          + sum({func}(vec, arrayMap(i -> to{type}(i + 2), range({vector_size})), {p}))
          + sum({func}(vec, arrayMap(i -> to{type}(i + 3), range({vector_size})), {p}))
          + sum({func}(vec, arrayMap(i -> to{type}(i + 4), range({vector_size})), {p}))
          + sum({func}(vec, arrayMap(i -> to{type}(i + 5), range({vector_size})), {p}))
          + sum({func}(vec, arrayMap(i -> to{type}(i + 6), range({vector_size})), {p}))
          + sum({func}(vec, arrayMap(i -> to{type}(i + 7), range({vector_size})), {p}))
          + sum({func}(vec, arrayMap(i -> to{type}(i + 8), range({vector_size})), {p}))
        FROM tab_{type}
    