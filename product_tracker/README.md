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
