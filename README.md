# MBTI_API

This is a simple API made with the Go programming language. The data is broken down into two fields.  One column contains the Meyers-Briggs personality type of the poster, and the other field contains posts authored the user in question seperated by two pipes. There is a script for seeding and creating the SQLite3 database and a second script to start a local server to serve the JSON.

**I did not implement any security mechanisms at this time.  Do not use this application in a production environment or any other place you will receive unknown input from users.  This application is vulnerable to SQL injection due to the method in which the server handles incoming HTTP requests.**

# Instructions for setup:

Clone this repository

Download the CSV dataset [here](https://www.kaggle.com/datasnaek/mbti-type/data) and save it in the 'data' directory.  Leave the file name as 'mbti_1.csv'

If you want to compile the binaries yourself or run the scripts directly, you can will need to [install Go](https://golang.org/doc/install).  To run the go scripts or to compile the binaries after the language is installed on your computer, use the following instructions:

You can simply run the scripts using the commands below from the root directory of the cloned project.
```bash
> go run src/setup_db/populate_database.go
#this step may take a few minutes
> go run src/API/MBTI_API.go
#this will launch your server
```

if you'd like to compile the binaries and run them, you can do so with the following commands:

```bash
> go build src/setup_db/populate_database.go
> ./src/setup_db/populate_database
#this step may take a few minutes
> go build src/API/MBTI_API.go
> ./bin/API/MBTI_API
#this will launch your server
```

Then you can navigate to localhost:3000/MBTI?type=<MBTI Type> to see json by MBTI type.
This will need to be in all caps with a lowercase type in order to view the proper page.

The API has the following endpoints available:
```
localhost:3000/MBTI?type=ISTJ
localhost:3000/MBTI?type=ISFJ
localhost:3000/MBTI?type=INFJ
localhost:3000/MBTI?type=INTJ
localhost:3000/MBTI?type=ISTP
localhost:3000/MBTI?type=ISFP
localhost:3000/MBTI?type=INFP
localhost:3000/MBTI?type=INTP
localhost:3000/MBTI?type=ESTP
localhost:3000/MBTI?type=ESFP
localhost:3000/MBTI?type=ENFP
localhost:3000/MBTI?type=ENTP
localhost:3000/MBTI?type=ESTJ
localhost:3000/MBTI?type=ESFJ
localhost:3000/MBTI?type=ENFJ
localhost:3000/MBTI?type=ENTJ
```

Sample output:

```json
{"MBTI_type":"ENTJ","Posts":"'Post one data||Post two data||Post three data'"}
```
