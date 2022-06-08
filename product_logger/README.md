# Product tracker

Keeps track of all penguin open box products, their **discounted price, discount
percentage and original price**.

A single product:

```json
{
   "title": "A jar of pickles",
   "price": 100,
   "discount_precentage": 50,
   "appearences": 1
}
```

`products.json` - contains all products that this script has kept track of over
time.

`def validate(product)` - checks the product title stored in
`current_product.json`, if product title is the same as current product on site,
write current product info to file and return False. Else write current product
info to file and return true.

`current_products.json` - contains the product that was retrieved when script
was last ran.

`def to_file()` - loads all products from file, and parse the array to see if
current product exists. If it does, remove the product stored.

-  Calculate average price and discount percentage
