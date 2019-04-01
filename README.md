# yorha-api
Simple api handled with dependency injection with a Nier:Automata tematic.

#### Technologies used in this project.
* Golang: Language used across the project.
* Iris: Golang framework.
* Hero: Dependency injection helper for iris.
* Gorm: Orm helper for golang.

### Usage:
Create a mysql database called <b>yorha_db</b> with user <b>yorha_user</b> and pass <b>yorha_pass</b>
#### In order to populate the database run <a href="https://gist.github.com/DarkoVR/cfec38abbada16e0cffa514a86a9228f">yorha-migration.go</a> that will create everything in the database
`go run yorha-migration.go`
#### Run the project
`go run main.go`
#### Make petitions by yourself, check <a href="https://github.com/DarkoVR/yorha-api/wiki">wiki</a> for further references
