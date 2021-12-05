import sys
from bs4 import BeautifulSoup
import requests
import json
import os
from dotenv import load_dotenv

# return in the index which the search term appears in the array. -1 if nothing is found
def index(array, search_term) -> int:
    index = 0
    for current in array:
        if current["title"] == search_term["title"]:
            return index
        index = index + 1

    return -1


# add the current product to products.json
def to_file(current_product):
    # load all products
    with open("/home/shawn/python/web_scraping/penguin_bots/product_tracker/products.json", "r+") as file:
        product_arr = json.load(file)

    # load current product
    #  with open("/home/shawn/python/web_scraping/penguin_bots/product_tracker/current_product.json", "r+") as file:
        #  current_product = json.load(file)

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

        print("product saved:", current_product)

    with open("/home/shawn/python/web_scraping/penguin_bots/product_tracker/products.json", "w+") as file:
        json.dump(product_arr, file, indent=4)

