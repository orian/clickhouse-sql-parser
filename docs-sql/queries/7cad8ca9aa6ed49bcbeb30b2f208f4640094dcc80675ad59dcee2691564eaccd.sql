  SELECT groupArray((class_id, frac)) AS priors
  FROM (SELECT class_id, count() / sum(count()) OVER () AS frac FROM docs GROUP BY class_id);
