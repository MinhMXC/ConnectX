# Setup
Recommended IDE: GoLand
1. Clone the repository.
2. Open up ```db_credentials.yaml``` and put in your MySQL settings / configurations.
3. Inside the root folder, run ```go run . setup```. This will create a new database and seed it with some values. If the database has already existed, it will be deleted.
4. To start the server, run ```go run .```
