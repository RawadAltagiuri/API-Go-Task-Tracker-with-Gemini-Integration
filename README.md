## Go Task Tracker with Gemini Integration
 This Go project provides a streamlined way to record your daily tasks, leverage Gemini's powerful sorting capabilities, and visualize them in a structured format.


## Key Features
* HTTP API: Send a POST request with your daily task summary text.
* Gemini Integration: Tasks are intelligently ordered based on your Gemini configuration.
* Structured Output: The API response presents tasks in a tabular format for easy analysis or integration with spreadsheets.

## Getting Started
* Clone the Repository: git clone https://github.com/your-username/go-task-tracker
* Install Dependencies: go mod download
* Configure Credentials (Gemini API key, Google API Credetentials, Google Sheets ID): Refer to the code for instructions if necessary.
* Run the Application: go run main.go
* Send a POST Request: Make a POST request to the specified endpoint (see code for details) with your daily task summary in the request body.

## Example Usage
### Request Body:

```json
{
  "Text": "Meeting at 10am, Worked on feature X, Fixed bug Y"
}
```

### Result (In google sheet):

| Task                 | Start Date | End Date | Time Spent | Notes | Assignee | Status |
|-----------------------|------------|----------|------------|-------|---------|---------|
| Meeting at 10am       | -          | -        |            |        | Rawad    | Done     |
| Fixed bug Y           | -          | -        |            |        | Rawad    | Done     |
| Worked on feature X   | -          | -        |            |        | Rawad    | In Progress|

## Customize Output Format:

Modify the vr.Values section in the code to adjust the output fields (task, start date, end date, time spent, notes, assignee, status) according to your preferences.

## Further Enhancements
* Implement authentication and authorization
* Add filtering or searching based on task attributes
* Integrate with popular project management tools
