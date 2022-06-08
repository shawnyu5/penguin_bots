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

sys.path.insert(1, os.path.dirname(sys.path[0]))  # utils
import utils  # type: ignore
from pprint import pprint


class Tracker:
    def __init__(self):
        self.product = {
            "title": "",
            "average_discount": "",
            "average_price": "",
            "appearances": 1,
        }
        load_dotenv()
        # url of penguin open box website
        url = "https://www.penguinmagic.com/openbox/"
        html_page = requests.get(url).text
        self.soup = BeautifulSoup(html_page, "html.parser")

        client = MongoClient(os.getenv("key"))
        # reference to the database collection
        self.db = client.penguin_magic.open_box

    # turns the paramers passed in into an object.
    def __to_object(self, title: str, discount_percentage: str, discount_price: str):
        return {
            "title": title,
            "average_discount": discount_percentage,
            "average_price": discount_price,
            "appearances": 1,
        }

    # retrieves product info from penguin and returns an object.
    def get_product_info(self):
        title = utils.get_title(self.soup)
        discount_percentage = utils.get_discount_percentage(self.soup)
        discounted_price = utils.get_discounted_price(self.soup)
        self.product = self.__to_object(title, discount_percentage, discounted_price)

    def valid(self) -> bool:
        """
        check if current product from penguin is different from the product in

        Returns:

            True: if current product is different from the product in `current_product.txt`

            False: if current product is the same as the product in `current_product.txt`

        """
        # read current product from current_product.json
        with open("current_product.txt", "r") as file:
            file_product = file.read()

        # if product from file is same as current product on site, return false
        if file_product == self.product["title"]:
            with open("current_product.txt", "w") as file:
                # update product in current_product.txt with product title
                file.write(self.product["title"])
            return False
        else:
            return True

    def save(self):
        """
        save current product to database
        """
        # load current product
        #  with open("current_product.json", "r") as file:
        #  current_product = json.load(file)
        # print("current product is ", current_product)

        with open("current_product.txt", "w") as file:
            # save current product from penguin to file
            file.write(self.product["title"])

        # checks if current product is logged
        found = self.db.find_one({"title": {"$eq": self.product["title"]}})
        old_data = found

        # if product is logged, ...
        if found:
            # update appearances
            found["appearances"] = found["appearances"] + 1

            average: float = found["average_price"] * (found["appearances"] - 1)
            average = (average + self.product["average_price"]) / found["appearances"]

            # update average price
            found["average_price"] = average

            # update average discount
            discount: float = found["average_discount"] * (found["appearances"] - 1)
            discount = (discount + self.product["average_discount"]) / found[
                "appearances"
            ]
            found["average_discount"] = discount

            # only save to database when dev is false
            if os.getenv("dev") == "False":
                self.db.update_one(
                    {"title": old_data["title"]},
                    {  # type: ignore
                        "$set": {
                            "appearances": found["appearances"],
                            "average_price": found["average_price"],
                            "average_discount": found["average_discount"],
                        }
                    },
                )

            print("product updated")
            pprint(found)
        else:
            print("product saved:")
            pprint(self.product)

            self.db.insert_one(
                self.product
            )  # NOTE: insert one adds "_id" attribute to object, which screws up writing to file

        return self.product

    @staticmethod
    def run():
        tracker = Tracker()

        tracker.get_product_info()
        if not tracker.valid():
            print("Product has not changed:")
            pprint(tracker.product)
            return

        tracker.save()


def main():
    Tracker.run()


if __name__ == "__main__":
    main()
