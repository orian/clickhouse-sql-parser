CREATE TABLE Employees (Name String) ENGINE = Memory;
INSERT INTO Employees VALUES ('John'), ('Jane'), ('Bob');

SELECT groupConcat(Name) FROM Employees;
