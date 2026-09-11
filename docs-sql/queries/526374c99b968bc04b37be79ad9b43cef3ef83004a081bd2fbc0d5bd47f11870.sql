INSERT INTO test.visits
    (CounterID, StartDate, Sign, IsNew, VisitID, UserID,
     Goals.ID, Goals.Serial, Goals.EventTime, Goals.Price, Goals.OrderID, Goals.CurrencyID)
VALUES
    -- A visit with two goals: each nested sub-column gets an array of length 2
    (101500, '2014-03-18', 1, 1, 2001, 200001,
     [1073752, 591325], [1, 2],
     ['2014-03-18 10:00:00', '2014-03-18 10:05:00'],
     [100, 200], ['order_a', 'order_b'], [1, 2]),
    -- A visit with no goals: all nested sub-columns get empty arrays
    (101500, '2014-03-18', 1, 0, 2002, 200002,
     [], [], [], [], [], []);
