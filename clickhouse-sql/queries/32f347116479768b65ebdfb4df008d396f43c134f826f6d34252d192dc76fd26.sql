
        SELECT count() FROM nested_events
        PREWHERE has(arrayFlatten(arrayMap(s -> s.dealId, event.slots.directDeals)), 'target_deal')
        WHERE (has(event.targeting['geo'], 'US') OR has(event.targeting['ipGeo'], 'US')) AND event.liveEvent = true
    