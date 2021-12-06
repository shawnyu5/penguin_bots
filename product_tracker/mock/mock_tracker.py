import os
import sys
from bs4 import BeautifulSoup
import requests
import json
sys.path.insert(1, '/home/shawn/python/web_scraping/penguin_bots/') # utils
import utils  # type: ignore
from pprint import pprint
from dotenv import load_dotenv
from pymongo import MongoClient

# return in the index which the search term appears in the array. -1 if nothing is found
def index(array, search_term) -> int:
    index = 0
    for current in array:
        if current["title"] == search_term["title"]:
            return index
        index = index + 1

    return -1

def validate(product) -> bool:

    # read current product from current_product.json
    with open("current_product.json", "r") as file:
        file_product = json.load(file)

    # save current product to current_product.json
    with open("current_product.json", "w") as file:
        json.dump(product, file, indent=4)

    # if product from file is same as current product on site, return false
    if file_product["title"] == product["title"]:
        return False
    else:
        return True


# add the current product to products.json
def to_file():
    # load all products
    with open("products.json", "r+") as file:
        product_arr = json.load(file)

    with open("current_product.json", "r+") as file:
        current_product = json.load(file)

    # checks if current product is logged and return the index.
    found_index = index(product_arr, current_product)

    # if product is logged, remove the product from list
    if found_index != -1:
        removed_product = product_arr.pop(found_index)

        # not all products saved has a appearance attribute
        if "appearances" not in removed_product:
            current_product["appearances"] = 1
        else:
            current_product["appearances"] = removed_product["appearances"] + 1
            current_product["price"] =  current_product["price"] / current_product["appearances"]

        # calculate average percentage
        current_product["discount_percent"] = (current_product["discount_percent"] + removed_product["discount_percent"]) / current_product["appearances"];

    # add current product to list
    product_arr.append(current_product)

    pprint("product saved:", current_product)

    with open("products.json", "w+") as file:
        json.dump(product_arr, file, indent=4)


load_dotenv()
client = MongoClient(os.getenv("key"))
db = client.penguin_magic.open_box

current_product = {
    "title": "Jar of Pickles",
    "discount_percent": 76,
    "price": 11.77,
    "appearances": 1
}

if not validate(current_product):
    print("Product has not changed:")
    pprint(current_product)
    exit(0)

to_file()

# html_page = requests.get("https://www.penguinmagic.com/p/3901").text
# soup = BeautifulSoup(html_page, "html.parser")

# product_title = utils.get_title(soup)
# product_discount_percentage = utils.get_discount_percentage(soup)
# product_discount_price = utils.get_discounted_price(soup)
