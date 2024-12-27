Provide instructions to run the application:

markdown
Copy code
# YouTube Fetcher

## Prerequisites
- Go installed
- PostgreSQL database setup

## Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/its-kundan/task-yt.git
   cd youtube-fetcher
Create a .env file:

env
Copy code
DB_URL=postgres://username:password@localhost:5432/youtube_fetcher?sslmode=disable
YOUTUBE_API_KEY_1=YOUR_API_KEY_1
YOUTUBE_API_KEY_2=YOUR_API_KEY_2
Run the application:

bash
Copy code
go run main.go
Access the API:

GET /videos?page=1&limit=10
yaml
Copy code

---

This setup fulfills the requirements. For the bonus, a dashboard can be added using a frontend framework like React, and served through a Gin route.





