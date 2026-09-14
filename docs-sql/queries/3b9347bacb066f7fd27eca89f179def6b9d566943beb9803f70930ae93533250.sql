-- ERROR: Goals.ID has 2 elements, but Goals.Serial has 1
INSERT INTO test.visits
    (CounterID, StartDate, Sign, IsNew, VisitID, UserID,
     Goals.ID, Goals.Serial, Goals.EventTime, Goals.Price, Goals.OrderID, Goals.CurrencyID)
VALUES
    (101500, '2014-03-18', 1, 1, 2003, 200003,
     [1073752, 591325], [1],
     ['2014-03-18 12:00:00'], [0], [''], [0]);
