SELECT
  groupArray(1)(date)[1] AS date,
  arrayFold(
    (acc, x) -> arrayMap(i -> plus(acc[i], x[i]), range(1, plus(length(date), 1))),
    groupArray(ifNull(total, 0)),
    arrayWithConstant(length(date), reinterpretAsFloat64(0))
  ) AS total,
  arrayMap(
    i -> if(
      ifNull(ifNull(greaterOrEquals(row_number, 25), 0), 0),
      '$$_posthog_breakdown_other_$$',
      i
    ),
    breakdown_value
  ) AS breakdown_value
FROM
  (
    SELECT
      arrayMap(
        number -> plus(
          toStartOfInterval(
            assumeNotNull(toDateTime('2025-09-09 00:00:00', 'UTC')),
            toIntervalHour(1)
          ),
          toIntervalHour(number)
        ),
        range(
          0,
          plus(
            coalesce(
              dateDiff(
                'hour',
                toStartOfInterval(
                  assumeNotNull(toDateTime('2025-09-09 00:00:00', 'UTC')),
                  toIntervalHour(1)
                ),
                toStartOfInterval(
                  assumeNotNull(toDateTime('2025-12-08 08:59:59', 'UTC')),
                  toIntervalHour(1)
                )
              )
            ),
            1
          )
        )
      ) AS date,
      arrayMap(
        _match_date -> arraySum(
          arraySlice(
            groupArray(ifNull(count, 0)),
            indexOf(groupArray(day_start) AS _days_for_count, _match_date) AS _index,
            plus(
              minus(
                arrayLastIndex(
                  x -> ifNull(equals(x, _match_date), isNull(x) AND isNull(_match_date)),
                  _days_for_count
                ),
                _index
              ),
              1
            )
          )
        ),
        date
      ) AS total,
      breakdown_value AS breakdown_value,
      rowNumberInAllBlocks() AS row_number
    FROM
      (
        SELECT
          sum(total) AS count,
          day_start AS day_start,
          [ifNull(toString(breakdown_value_1), '$$_posthog_breakdown_null_$$'), ifNull(toString(breakdown_value_2), '$$_posthog_breakdown_null_$$')] AS breakdown_value
        FROM
          (
            SELECT
              ifNull(
                sum(
                  accurateCastOrNull(
                    accurateCastOrNull(
                      has(e.properties_group_ai, '$ai_cache_read_input_tokens') ? e.properties_group_ai['$ai_cache_read_input_tokens'] : null,
                      'Float64'
                    ),
                    'Float64'
                  )
                ),
                0
              ) AS total,
              toStartOfHour(toTimeZone(e.timestamp, 'UTC')) AS day_start,
              ifNull(nullIf(toString(e.`mat_$ai_model`), ''), '$$_posthog_breakdown_null_$$') AS breakdown_value_1,
              ifNull(nullIf(toString(e.`mat_$ai_provider`), ''), '$$_posthog_breakdown_null_$$') AS breakdown_value_2
            FROM events AS e SAMPLE 1
            WHERE
              and(equals(e.team_id, 1), greaterOrEquals(toTimeZone(e.timestamp, 'UTC'), assumeNotNull(toDateTime('2025-09-09 00:00:00', 'UTC'))), lessOrEquals(toTimeZone(e.timestamp, 'UTC'), assumeNotNull(toDateTime('2025-12-08 08:59:59', 'UTC'))), equals(e.event, '$ai_generation'))
            GROUP BY
              day_start, breakdown_value_1, breakdown_value_2
          )
        GROUP BY
          day_start, breakdown_value_1, breakdown_value_2
        ORDER BY day_start ASC, breakdown_value ASC
      )
    GROUP BY
      breakdown_value
    ORDER BY
      if(has(breakdown_value, '$$_posthog_breakdown_other_$$'), 2, if(has(breakdown_value, '$$_posthog_breakdown_null_$$'), 1, 0)) ASC,
      arraySum(total) DESC,
      breakdown_value ASC
  )
WHERE arrayExists(x -> isNotNull(x), breakdown_value)
GROUP BY
  breakdown_value
ORDER BY
  if(has(breakdown_value, '$$_posthog_breakdown_other_$$'), 2, if(has(breakdown_value, '$$_posthog_breakdown_null_$$'), 1, 0)) ASC,
  arraySum(total) DESC,
  breakdown_value ASC
LIMIT 50000
SETTINGS
  readonly=2,
  max_execution_time=600,
  allow_experimental_object_type=1,
  format_csv_allow_double_quotes=0,
  max_ast_elements=4000000,
  max_expanded_ast_elements=4000000,
  max_bytes_before_external_group_by=0,
  transform_null_in=1,
  optimize_min_equality_disjunction_chain_length=4294967295,
  allow_experimental_join_condition=1,
  use_hive_partitioning=0
