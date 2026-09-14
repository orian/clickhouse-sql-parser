SELECT
    naiveBayesClassifier('review_sentiment', 'amazing food and friendly staff') AS positive_review,
    naiveBayesClassifier('review_sentiment', 'awful service and a terrible meal') AS negative_review;
