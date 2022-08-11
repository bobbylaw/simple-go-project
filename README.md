# Overview

This is a simple go project that is part of GOVTECH TAP 2023 GDS ACE Tech Assessment. The main goal is to be able to calculate the score of the teams and showcase which team advances into the next round of the championship based on the requirements.

## Set-Up
- Clone project into your local directory
- Download Golang and run `go build` followed by `go run main.go` on command line to launch the web application
- It will display `Starting server at port 8080` upon success

## Database
- This application is utilizing SQLite3 to persist database upon server restart / failure
- Upon running the program, the database will be generated with the following tables in `database/database.db`
  - `team`
  - `group_record`
  
## Usage
- The application consist of three pages
  - Registration of teams
  - Input of match results
  - Viewing of the final results

## Remarks
- In any event of failure, please delete `database.db` and relaunch the application.
- This application currently do not validate the inputs from user.
- The server runs on port `8080` by default.
