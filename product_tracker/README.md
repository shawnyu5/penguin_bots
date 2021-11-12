# Product tracker

Keeps track of all penguin open box products, their **discounted price, discount
percentage and original price**.

`products.json` - contains all products that this script has kept track of over
time.

`validate(product)` - checks the product title stored in `current_product.json`,
if product title is the same, write current product info to file and return
False. Else write current product info to file and return true.

`current_products.json` - contains the product that was retrieved when script
was last ran.

`def to_file()` - loads all products from file, and parse the array to see if
current product exists

# Testing

We should run the script once on a known product.

Check if `current_product.json` contains the correct product

Then run the script on a different product

The previous product should be saved to `products.json`. `current_product.json`
should also change to reflect current product.
