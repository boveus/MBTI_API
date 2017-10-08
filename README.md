# MBTI_API

This is my attempt to create an API using the Go programming language.  I am planning to use a sqlite3 database that will be populated from a CSV file located [here](https://www.kaggle.com/datasnaek/mbti-type/data).  The database seeding file will be a seperate binary from the API server.

Once the database is seeded, there will be a seperate program that renders data from the database as JSON.  There isn't much complexity to the dataset, so I am only going to make an API rather than do any analytics on it at this time.  I may consume this API in another application to analyze the data, though.  For now this is purely an exercise to familiarize myself with APIs and the Go programming language.
