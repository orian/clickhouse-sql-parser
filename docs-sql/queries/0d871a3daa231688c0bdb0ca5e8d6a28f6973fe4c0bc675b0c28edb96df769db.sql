CREATE HYPOTHETICAL PROJECTION p_by_b ON t (SELECT a, b ORDER BY b);
SELECT name, type, sorting_key FROM system.hypothetical_projections;
