# stock-service
Service to handle the stock of products.

## Notes
It was implemented only the POST and the GET endpoint because I had just a few hours to do this assessment.

### Endpoints

| Endpoint                         | Method | Description                                 
|----------------------------------|--------|---------------------------------------------|
| `/stock`                         | POST   | To create new stock                         |
| `/stock/:product_code`           | GET    | Returns the stock information for that code |
| `/stock/:product_code`           | DELETE | Deletes the stock information               |
| `/stock/:product_code`           | PUT    | Update the stock information                |
| `/stock/:product_code`           | PATCH  | Modifify the stock information              |


## How to run the application:
Clone the repository and move to the project folder:
```
git clone https://github.com/juanmabaracat/stock-service.git
cd stock-service
```
Run the application:
```
make run
```

Run the tests:

```
make test
```

## Examples
### Create a new stock
### POST `/stock`

| Code  | Description   |
| ----  | ------------- |
| 201   | Game created  |
| 400   | Bad request   |
| 500   | Server error  |

```
curl -i -X POST http://localhost:8080/stock -d '{"product_name":"Coca-Cola", "stock_quantity":500}'
```

#### Response
```
{"product_code":"a84860b8-f058-4fc4-8140-022dc1b5437e"}
```

### Get a stock information
### GET `/stock/a84860b8-f058-4fc4-8140-022dc1b5437e`

| Code | Description  |
|------|--------------|
| 200  | Status OK    |
| 400  | Bad request  |
| 500  | Server error |

```
curl -i http://localhost:8080/stock/a84860b8-f058-4fc4-8140-022dc1b5437e
```

#### Response
```
{
"ProductCode": "75e71c26-93cd-4d06-8ed3-67a5e13acf84",
"ProductName": "Coca-Cola",
"StockQuantity": 500
}
```