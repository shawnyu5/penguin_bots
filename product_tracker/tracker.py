# !/usr/bin/env python3
# purpose of this file:
# Date: 2021-10-13
# ---------------------------------
from pymongo import MongoClient
import sys
from bs4 import BeautifulSoup
import requests
import json
import os
from dotenv import load_dotenv
sys.path.insert(1, '/home/shawn/python/web_scraping/penguin_bots/') # utils
import utils  # type: ignore
from pprint import pprint

def to_object(title:str, discount:str, price:str):
    return {
            "title": title,
            "average_discount": discount,
            "average_price": price
            }

# check if current product is different from the product in current_product.json
def validate(product) -> bool:
    current_product_json = "/home/shawn/python/web_scraping/penguin_bots/product_tracker/current_product.json"
    # read current product from current_product.json
    with open(current_product_json, "W") as file:
        file_product = json.load(file)

    # save current product to current_product.json
    with open(current_product_json, "W") as file:
        json.dump(product, file, indent=4)

    # if product from file is same as current product on site, return false
    if file_product["title"] == product["title"]:
        return False
    else:
        return True

# return in the index which the search term appears in the array. -1 if nothing is found
def index(array, search_term) -> int:
    index = 0
    for current in array:
        if current["title"] == search_term["title"]:
            return index
        index = index + 1
    return -1

# add the current product to products.json
def save():
    # load current product
    with open("current_product.json", "r") as file:
        current_product = json.load(file)

    # checks if current product is logged
    found = db.find_one({ "title": { "$eq": current_product["title"] }})
    old_data = found
    # print("old data is ", old_data)

    # if product is logged, ...
    if found:

        # update appearances
        found["appearances"] = found["appearances"] + 1
        # update price
        found["average_price"] = (float(current_product["average_price"]) + found["average_price"]) / found["appearances"]
        # calculate average percentage
        found["average_discount"] = found["average_discount"] / found["appearances"];
        # print("found is ", found)

        db.update_one({ "_id": old_data["_id"] }, { #type: ignore
            "$set": {
                "appearances": found["appearances"],
                "average_price": found["average_price"],
                "average_discount": found["average_discount"]
                }
            })
        print("product updated")
        pprint(found)
    else:
        db.insert_one(current_product)
        print("product saved:")
        pprint(current_product)


def main():
    load_dotenv()
    url = str(os.getenv("url"))
    print(url)

    client = MongoClient(os.getenv("key"))
    global db
    db = client.penguin_magic.open_box

    html_page = requests.get(url).text
    soup = BeautifulSoup(html_page, "html.parser")

    product_title = utils.get_title(soup)
    product_discount_percentage = utils.get_discount_percentage(soup)
    product_discount_price = utils.get_discounted_price(soup)

    product = to_object(product_title, product_discount_percentage, product_discount_price) #type: ignore

    # if product has not changed, save product to file and don't parse json
    # file
    if not validate(product):
        print("Product has not changed")
        print(product)
        exit(0)

    save()

if __name__ == "__main__":
    main()

