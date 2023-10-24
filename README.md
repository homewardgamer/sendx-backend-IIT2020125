# SendX - Web Crawler Service

This is a web crawler service built in Go, utilizing the Gin framework. The project provides API endpoints to facilitate web crawling and offers functionality to manage worker counts, rate limits, and retrieve current configurations.

## Project Structure

- **api**: Contains the API handlers and middlewares.
- **pkg**: Houses various packages such as:
  - `middleware`: Middleware functionalities.
  - `crawler`: Web crawling functionalities.
  - `storage`: Data storage mechanisms.
  - `worker`: Background processing or tasks.
- **assets**: Static assets including:
  - `css`: Stylesheets.
  - `images`: Image assets.
  - `js`: JavaScript files.
- **template**: HTML templates.
  - `index.html`: The main or home page template.
- **config**: Configuration-related functionalities.

## Functionality Overview

1. **Crawl Request Handling**:
   - Accepts a URL to be crawled.
   - Retrieves previously stored pages, if available.
   - Queues jobs for new URLs, stores and returns results.

2. **Worker Management**:
   - Update the count of workers for both paying and non-paying customers.

3. **Rate Limit Management**:
   - Adjust the crawl rate limit (e.g., pages per hour).

4. **Configuration Retrieval**:
   - Retrieve current configurations like worker counts and rate limits.

## API Routes

Detailed routes and their respective functionalities are defined within the project. Please refer to the source code for specifics.

## Setup & Usage

- Make sure you have Go installed in your local environment by running `go` in your terminal
- Install the gin framework by using the command `go get -u github.com/gin-gonic/gin`
- Install necesssary dependencies by `go install`
- Run the server by `go run .`
---