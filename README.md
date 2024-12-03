# TESC-Farm
The Evergreen State College Organic Farm Project

## API
Compiled API written in Go to collect data for labor, harvests, and sales.

### Labor Endpoints:
- GET  /hours
- GET  /hours/:id
- POST /hours/new
- POST /hours/:id/update
- POST /hours/:id/delete
- GET  /workers
- GET  /worker/:id
- GET  /worker/:id/hours
- POST /worker/new
- POST /worker/:id/update
- POST /worker/:id/delete

### Harvest Endpoints:
- GET  /crops
- GET  /crop/:id
- GET  /crop/:id/harvests
- GET  /crop/:id/processing
- POST /crop/new
- POST /crop/:id/update
- POST /crop/:id/delete
- GET  /harvests
- GET  /harvest/:id
- GET  /harvest/:id/processing
- POST /harvest/new
- POST /harvest/:id/update
- POST /harvest/:id/delete
- GET  /processing
- GET  /process/:id
- POST /process/new
- POST /process/:id/update
- POST /process/:id/delete
- GET  /bin/:bin/harvest

### Sales Endpoints:
- GET  /products
- GET  /product/:id
- GET  /product/:id/sales
- POST /product/new
- POST /product/:id/update
- POST /product/:id/delete
- GET  /sales
- GET  /sale/:id
- POST /sale/new
- POST /sale/:id/update
- POST /sale/:id/delete
