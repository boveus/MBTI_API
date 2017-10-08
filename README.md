# MBTI_API

This is a simple API made with the Go programming language. The data is broken down into two columns.  One column contains the Meyers-Briggs personality type, and the other column contains posts by the user in question seperated by two pipes. There is a script for seeding the database and a second script to start a local server and serve the JSON.

I did not implement any security mechanisms at this time.  Do not use this application in a production environment or any other place you will receive unknown input from users.  This application is vulnerable to SQL injection due to the method in which the server handles incoming HTTP requests.  

Instructions for setup:

Clone this repository

Download the CSV dataset [here](https://www.kaggle.com/datasnaek/mbti-type/data) and save it in the 'data' directory.  Leave the file name as 'mbti_1.csv'

If you want to compile the binaries yourself, you can will need to [install Go](https://golang.org/doc/install).  To run the binaries after the language is installed on your computer, use the following commands:

You can simply run the scripts using the commands below from the root directory of the cloned project.
```bash
> go run setup_db/populate_database.go
#this step may take a few minutes
> go run API/MBTI_API.go
#this will launch your server
```

if you'd like to compile the binaries and run them, you can do so with the following commands:
```bash
> go build setup_db/populate_database.go
> ./setup_db/populate_database
#this step may take a few minutes
> go build API/MBTI_API.go
> ./API/MBTI_API
#this will launch your server
```

Then you can navigate to localhost:3000/MBTI?type=<MBTI Type> to see json by MBTI type.
This will need to be in all caps with a lowercase type in order to view the proper page
