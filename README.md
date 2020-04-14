# Omni reviews
## Requirements
* Ensure Go SDK 1.14 or higher is installed
* Ensure MongoDB database with name `omni` with collections `reviews` and `phrases` is running 
## How To  Run
### Review Crawler
* go to directory `review-crawler`
* change MongoDB connection to valid one in `main.go`
* run command `go run main.go`
### Review Phrase API
* go to directory `review-phrase-api`
* change MongoDB connection to valid one
* run command `go run main.go` in `main.go`
