
        CREATE TABLE nested_events
        (
            event Tuple(
                bidtime Nullable(Int64),
                sourceid Nullable(Int64),
                liveEvent Bool,
                targeting Map(String, Array(Nullable(String))),
                slots Array(Tuple(
                    bidId Nullable(String),
                    directDeals Array(Tuple(
                        activeDeal Nullable(Bool),
                        bidderIds Array(Nullable(Int64)),
                        currency Nullable(String),
                        dealCPM Nullable(Int64),
                        dealId Nullable(String),
                        wseats Array(Nullable(String)),
                        dealTerms Nullable(String))),
                    dtbMaxBid Nullable(Int64))))
        )
        ENGINE = MergeTree ORDER BY tuple()
        SETTINGS index_granularity = 28, max_bytes_to_merge_at_max_space_in_pool = 1
    