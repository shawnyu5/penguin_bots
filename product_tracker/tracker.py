# !/usr/bin/env python3
# purpose of this file: penguin open box tracker
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

class Tracker:
    def __init__(self):
        self.product = {
                "title": "",
                "average_discount": "",
                "average_price": ""
                }
        load_dotenv()
        # url of penguin open box website
        url = str(os.getenv("url"))
        html_page = requests.get(url).text
        self.soup = BeautifulSoup(html_page, "html.parser")

        client = MongoClient(os.getenv("key"))
        # reference to the database collection
        self.db = client.penguin_magic.open_box


    # turns the paramers passed in into an object.
    def __to_object(self, title:str, discount_percentage:str, discount_price:str):
        return {
                "title": title,
                "average_discount": discount_percentage,
                "average_price": discount_price
                }

    # retrieves product info from penguin and returns an object.
    def get_product_info(self):
        title = utils.get_title(self.soup)
        discount_percentage = utils.get_discount_percentage(self.soup)
        discounted_price = utils.get_discounted_price(self.soup)
        self.product = self.__to_object(title, discount_percentage, discounted_price)

    # check if current product from penguin is different from the product in
    # `current_product.json`
    def valid(self) -> bool:
        current_product_json = "/home/shawn/python/web_scraping/penguin_bots/product_tracker/current_product.json"
        # read current product from current_product.json
        with open(current_product_json, "r") as file:
            file_product = json.load(file)

        with open(current_product_json, "w") as file:
            # save current product to file
            json.dump(self.product, file, indent=4)

        # if product from file is same as current product on site, return false
        if file_product["title"] == self.product["title"]:
            return False
        else:
            return True

    # add the current product to products.json
    def save(self):
        # load current product
        with open("current_product.json", "r") as file:
            current_product = json.load(file)

        # checks if current product is logged
        found = self.db.find_one({ "title": { "$eq": current_product["title"] }})
        old_data = found

        # if product is logged, ...
        if found:
            # update appearances
            found["appearances"] = found["appearances"] + 1
            # update price
            found["average_price"] = (float(current_product["average_price"]) + found["average_price"]) / found["appearances"]
            # calculate average percentage
            found["average_discount"] = found["average_discount"] / found["appearances"];
            # print("found is ", found)

            self.db.update_one({ "_id": old_data["_id"] }, { #type: ignore
                "$set": {
                    "appearances": found["appearances"],
                    "average_price": found["average_price"],
                    "average_discount": found["average_discount"]
                    }
                })
            print("product updated")
            pprint(found)
        else:
            self.db.insert_one(current_product)
            print("product saved:")
            pprint(current_product)

    @staticmethod
    def run():
        tracker = Tracker()
        tracker.get_product_info()
        if not tracker.valid():
            print("Product has not changed:")
            pprint(tracker.product)

        tracker.save()


def main():
    Tracker.run()
    # load_dotenv()
    # url = str(os.getenv("url"))
    # print(url)

    # client = MongoClient(os.getenv("key"))
    # global db
    # db = client.penguin_magic.open_box

    # html_page = requests.get(url).text
    # soup = BeautifulSoup(html_page, "html.parser")

    # product_title = utils.get_title(soup)
    # product_discount_percentage = utils.get_discount_percentage(soup)
    # product_discount_price = utils.get_discounted_price(soup)

    # product = to_object(product_title, product_discount_percentage, product_discount_price) #type: ignore

    # if product has not changed, save product to file and don't parse json
    # file
    # if not validate(product):
        # print("Product has not changed")
        # print(product)
        # exit(0)

    # save()

if __name__ == "__main__":
    main()

