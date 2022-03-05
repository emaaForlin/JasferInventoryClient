# Jasfer Inventory Software

### How to use it
Required env vars
- `SRV_ADDR` this must to be address of the server api and the path. For example `SRV_ADDR=http://192.168.0.10/products`
- `APIKEY` the key required for to connect the client with the server, it is `APIKEY=user:pass`. Must to be in the `authorized_users` table in the database.

### Changelog
- Added api authentication
- Show the products in the index page
- Add products
- Edit products
- Modify products
- Update the prices every month automatically
- Update every price with a percentil at the moment

## Note for my self: Clean the code before merge it into main 