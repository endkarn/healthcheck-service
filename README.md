*noted : I don't have the frontend skill to make the webpage so I do my best on backend service*

# Checking Site Up/Down Status API

## Requirements.

- [Golang](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Newman (Postman CLI)](https://www.npmjs.com/package/newman) *basically required npm to install

## API List.
-  [POST /check](#check-sites-status) Checking sites status

### Check Sites Status
The API that receive CSV file from multipart/form-data request and will check sites status up/down then return as response

#### Method : POST /check
#### Request Body :

`multipart/form-data
"file" : "./awesome-sites.csv"
`

#### Response Body:
***200 OK***
```json
{
    "total_up": 2,
    "total_down": 0,
    "sites": [
        {
            "Order": 0,
            "WebsiteURL": "https://facebook.com",
            "HTTPStatusCode": 200
        },
        {
            "Order": 1,
            "WebsiteURL": "https://google.com",
            "HTTPStatusCode": 200
        }
    ]
}
```
***400 Bad Request***
```
file err : request Content-Type isn't multipart/form-data
```

#### Example of Accepted CSV File

*Another format except this example might not be work*

***Example 1.***
```csv
https://facebook.com  
https://google.com  
https://github.com  
https://gmail.com  
https://stackoverflow.com  
https://youtube.com  
https://store.steampowered.com/  
https://discord.com/  
https://chordtabs.in.th/
```
***Example 2.***
```csv
https://facebook.com,https://google.com,https://github.com,https://gmail.com,https://stackoverflow.com,https://youtube.com,https://store.steampowered.com/,https://discord.com/,https://chordtabs.in.th/
```
---
## How To Run
All of the commands on this service is on `Makefile`

**Run Unit Test**

`
go test ./... -coverprofile=coverage.out
`

**Build Docker Image**

`
docker compose build
`

**Start Service**

`
docker compose up -d
`

**Run Acceptance Test / API**
```bash
cd atdd
newman run healthcheck.success_collection.json -e healthcheck-env.postman_environment.json -d healthcheck.data_success.json  
newman run healthcheck.unsuccess_collection.json -e healthcheck-env.postman_environment.json -d healthcheck.data_unsuccess.json
```