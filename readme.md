# Mini Project Alta Batch 3

## go-shortener
A URL shortener service written in Go. This service is a part of the Mini Project Alta Kampus Merdeka Batch 3.

### How to run
1. Clone this repository
2. Create a `.env` file in the root directory of the project
3. Copy the content of `.env.example` to `.env` file
4. Fill the `.env` file with your own configuration
5. Run `go run main.go`
6. Open your browser and go to `localhost:8080`
7. Enjoy!

### Features
1. URL
    - [x] Create a short URL
    - [x] Redirect to the original URL
    - [x] Delete a short URL
    - [X] Update the URL
    - [X] Set Expired Date
    - [x] Get all short URLs by User ID
    - [x] Get a short URL details `New Feature!`
2. User
   - [x] Register 
   - [x] Login

### API Endpoints
1. URL
   -  Create a short URL
     ```
    POST {{HOST}}/urls
    ```
   - Update a short URL
   ```
    PUT {{HOST}}/urls/:short
    ```
    - Delete a short URL
    ```
    DELETE {{HOST}}/urls/:short
    ```
    - Redirect to the original URL
    ```
    GET {{HOST}}/:short
    ```
    - Get all short URLs by User ID
    ```
    GET {{HOST}}/urls
    ```
   - Get a short URL details
    ```
    GET {{HOST}}/urls/details/:short
    ```
2. User
   - Register
   ```
    POST {{HOST}}/register
    ```
   - Login
   ```
    POST {{HOST}}/login
    ```
   
### Tech Stack & Tools
1. Go
2. Echo
3. GORM
4. MySQL
5. Docker
6. Postman
7. Git
8. Validator
9. JWT
10. godotenv