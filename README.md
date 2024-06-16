# Exoplanet Service

This microservice provides functionalities to manage exoplanets.

## Endpoints

- `POST /exoplanets` - Add a new exoplanet
- `GET /exoplanets` - List all exoplanets
- `GET /exoplanets/{id}` - Get an exoplanet by ID
- `PUT /exoplanets/{id}` - Update an exoplanet
- `DELETE /exoplanets/{id}` - Delete an exoplanet
- `GET /exoplanets/{id}/fuel?crew_capacity=5` - Get fuel estimation for a trip

## Build and Run

1. Build the Docker image:
    ```sh
    docker build -t exoplanet-service .
    ```

2. Run the Docker container:
    ```sh
    docker run -p 8080:8080 exoplanet-service
    ```

3. Access the service at `http://localhost:8080`

## Example
GetAll Exo planet
![image](https://github.com/AviralDixit-star/exoplanet-service/assets/61451663/7ea38a5f-f776-4be4-8e67-5b3c1b448cb8)
Create Exo Planet
![image](https://github.com/AviralDixit-star/exoplanet-service/assets/61451663/9e696b79-4a3f-45f9-897c-2119b48a99f7)
Get by Id
![image](https://github.com/AviralDixit-star/exoplanet-service/assets/61451663/c69690bc-48c5-4308-8759-1ab5f127959c)
Update
![image](https://github.com/AviralDixit-star/exoplanet-service/assets/61451663/42027679-fb88-4701-b891-cbe8038e5d81)
Delete
![image](https://github.com/AviralDixit-star/exoplanet-service/assets/61451663/061788ea-83b2-4ccc-8ce2-e7fa809de372)
Fuel Estimate
![image](https://github.com/AviralDixit-star/exoplanet-service/assets/61451663/d736c871-4b67-4e78-962f-9c9bf98c0df1)





