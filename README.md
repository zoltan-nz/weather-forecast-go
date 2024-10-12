# Weather Forecast Web Application

## 1. Project Summary

We're building a web-based weather forecast application that allows users to search for weather information by city name. The application will have both a user interface and an API backend.

### Key Features:
- City search with autocomplete
- Current weather display
- 5-day forecast
- Recent searches history
- Admin statistics page (protected by authentication)

## 2. Wireframes

### a) Home Page (index.html):
```
+----------------------------------+
|        Weather Forecast          |
|                                  |
|  +----------------------------+  |
|  |     Enter city name        |  |
|  +----------------------------+  |
|           [Search]               |
|                                  |
+----------------------------------+
```

### b) Weather Results Page (weather.html):
```
+----------------------------------+
|   Weather for [City Name]        |
|                                  |
|   Current Weather:               |
|   Temperature: XX°C              |
|   Humidity: XX%                  |
|   Wind Speed: XX km/h            |
|                                  |
|   5-Day Forecast:                |
|   +----------------------------+ |
|   | Date | Temp | Conditions   | |
|   |------|------|--------------|
|   | Day1 | XX°C | Sunny        | |
|   | Day2 | XX°C | Cloudy       | |
|   | Day3 | XX°C | Rainy        | |
|   | Day4 | XX°C | Partly Cloudy| |
|   | Day5 | XX°C | Clear        | |
|   +----------------------------+ |
|                                  |
+----------------------------------+
```

### c) Admin Statistics Page (stats.html):
```
+----------------------------------+
|        Admin Statistics          |
|                                  |
|   Recent Searches:               |
|   +----------------------------+ |
|   | City      | Search Count   | |
|   |-----------|----------------|
|   | City1     | XX             | |
|   | City2     | XX             | |
|   | City3     | XX             | |
|   | City4     | XX             | |
|   | City5     | XX             | |
|   +----------------------------+ |
|                                  |
|   Total Searches: XXXX           |
|                                  |
+----------------------------------+
```

## 3. API Endpoints

- GET /api/weather?city={cityname}
    - Returns weather data for the specified city
- GET /api/stats
    - Requires authentication
    - Returns statistics about recent searches

## 4. Todo List

1. Project Setup
    - [x] Initialize Go module
    - [x] Create basic project structure

2. Geo Location Service
    - [x] Implement FetchLatLong function (TDD)
    - [x] Add error handling and input validation

3. Weather Service
    - [ ] Implement FetchWeather function (TDD)
    - [ ] Parse and structure weather data (current and 5-day forecast)

4. API Server
    - [ ] Set up Gin web framework
    - [ ] Create /api/weather endpoint
    - [ ] Integrate Geo Location and Weather services

5. Database Integration
    - [ ] Set up PostgreSQL connection
    - [ ] Implement functions to store and retrieve recent searches

6. HTML Templates
    - [ ] Create index page with search form and autocomplete
    - [ ] Create weather results page with current weather and 5-day forecast
    - [ ] Create admin statistics page
    - [ ] Implement template rendering in API server

7. Authentication
    - [ ] Implement basic auth for /api/stats endpoint
    - [ ] Create /api/stats endpoint to show recent searches and total count

8. User Interface
    - [ ] Implement city search autocomplete functionality
    - [ ] Style pages with CSS for a better user experience

9. Error Handling and Logging
    - [ ] Implement consistent error handling across the application
    - [ ] Add logging for important events and errors

10. Testing
    - [ ] Unit tests for all main functions
    - [ ] Integration tests for API endpoints

11. Documentation
    - [ ] Write API documentation
    - [ ] Add usage instructions in README.md

## 5. Data Structures

- LatLong: {Latitude: float64, Longitude: float64}
- CurrentWeather: {Temperature: float64, Humidity: float64, WindSpeed: float64}
- Forecast: {Date: string, Temperature: float64, Conditions: string}
- WeatherData: {City: string, CurrentWeather: CurrentWeather, Forecast: []Forecast}
- SearchStat: {City: string, Count: int}

## 6. External APIs

- Geocoding API: https://geocoding-api.open-meteo.com/v1/search
- Weather API: https://api.open-meteo.com/v1/forecast