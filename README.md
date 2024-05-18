# Grower's Diary microservice

PURPOSE - TRACKING YOUR HAPPY PLANTS GROWTH

## API Endpoints

### Strain
- **Create a new strain:** `POST /strains/create`
- **Update an existing strain:** `POST /strains/update`
- **Get information about a strain by its ID:** `GET /strains/{id}`
- **Delete a strain by its ID:** `DELETE /strains/{id}`
- **Get a list of all strains:** `GET /strains/list`

### GrowLog
- **Create a new grow log:** `POST /growlogs/create`
- **Update an existing grow log:** `POST /growlogs/update`
- **Get information about a grow log by its ID:** `GET /growlogs/{id}`
- **Delete a grow log by its ID:** `DELETE /growlogs/{id}`
- **Get a list of all grow logs:** `GET /growlogs/list`

### LogEntry
- **Create a new log entry:** `POST /logentries/create`
- **Update an existing log entry:** `POST /logentries/update`
- **Get information about a log entry by its ID:** `GET /logentries/{id}`
- **Delete a log entry by its ID:** `DELETE /logentries/{id}`
- **Get a list of all log entries:** `GET /logentries/list`
- **Upload photos for a log entry:** `POST /logentries/photos/upload`

## Usage

1. **Clone the repository:** `git clone <repository_url>`
2. **Navigate to the project directory:** `cd <project_directory>`
3. **Install dependencies:** `go mod tidy`
4. **Set up the PostgreSQL database:** 
   - Install PostgreSQL and create a new database.
   - Update the database configuration in the project.
5. **Run the server:** `go run main.go`
6. **Access the API endpoints using a REST client or web browser.**

