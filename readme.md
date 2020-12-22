Created basic routing in main.go
Used gorilla/mux package as I prefer to use it to the standard net/http package.

Created the JSON data and added some extra entries. Included an ID key to query later on for the getGame function.

Created the structs to organise the json data.



Decided to use gojsonq package to query the data. Makes querying json simple with SQL-like queries.
https://github.com/thedevsaddam/gojsonq

I started off using Go standard library for the getAll function, you can see the function commented out at the bottom of main.go
Changed it to gojsonq to keep the code consistent.

Used dep as the dependency management as I knew i wanted to dockerise the application later on.
I wouldnt usually use a dep manager for projects of this size usually but it was important to get this to be able to run
on your machines as well as mine quickly.

I apologise, I couldnt for the life of me complete task two fully. I started it but the package had trouble reading the nested data,
and outputting the data as I needed. By the time I realised this it was too late to go back to Go native.

Couldnt order outputted JSON because gojsonq uses maps and that orders it lexographically.
Would need to use structs not maps to preserve the order. 

I decided to use hardcoded json for this task to reduce the scope. I have experience with postgres databases and couldve used a DB for this task,
however I wasn't sure if I could complete that within the timeframe. In hindsight it would've made task two a lot easier.

Created a dockerfile and dockerized the application.

To run the application, please follow the instruction in the email.
Github Repo: 
https://github.com/JamieW87/CompHouse

Here is the command to run the docker image:
docker run --rm -p 8080:8080 jamiedovu/comphouse:1.0