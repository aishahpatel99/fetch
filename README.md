To run the code, first build the app:
`go build -o receipt-app .`

Then run the app:
`./receipt-app`

Then, hit the API endpoints (receipts/process or receipts/{id}/points) to test.

If using Curl to hit the endpoints, an example Curl request is:

for the process endpoint:
`curl -X POST http://localhost:8080/receipts/process \
     -H "Content-Type: application/json" \
     -d '{
       "retailer": "Target",
       "purchaseDate": "2022-01-01",
       "purchaseTime": "13:01",
       "items": [
         {
           "shortDescription": "Mountain Dew 12PK",
           "price": "6.49"
         },
         {
           "shortDescription": "Emils Cheese Pizza",
           "price": "12.25"
         },
         {
           "shortDescription": "Knorr Creamy Chicken",
           "price": "1.26"
         },
         {
           "shortDescription": "Doritos Nacho Cheese",
           "price": "3.35"
         },
         {
           "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
           "price": "12.00"
         }
       ],
       "total": "35.35"
     }'
`
You should receive a JSON object with an id for this receipt. 
Then, call the /points endpoint with that id:

`curl -X GET http://localhost:8080/receipts/{id}/points`

The default port to run this application on is 8080. If you need to change the port, you can change it on line 18 in main.go
