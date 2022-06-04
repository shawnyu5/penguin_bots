## Steps:

1. Get url from .env file

   -  `def get_webPage` - `load_dotenv` and create and return soup object

2. Get title of product

   -  `def get_title` - takes in soup object. Find product title

3. Get and validate title of product

   -  `def get_title` - takes in soup object. Find description and formate text.

4. Validate product to ensure it's a coin product - `def validate_product(product)`. Read from `product_info.txt` file. Make sure current product
   title is different from one in file

5. If all above condition is satisfied, read previous product from
   `product_info.txt`. And ensure the current title is different from title
   read. If it is, exit. `def validate(product_title)`

6. Get recipients from `.env` and Send email

**Extra**: Escape all `"` -> `def escape_quotes`
