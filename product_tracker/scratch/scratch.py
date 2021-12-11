from pymongo import MongoClient
import json
from pprint import pprint
import os
from dotenv import load_dotenv
import sys


def save():
    # load current product
    with open("current_product.json", "r") as file:
        current_product = json.load(file)

    # checks if current product is logged
    found = db.find_one({ "title": { "$eq": current_product["title"] }})
    old_data = found
    print("old data is ", old_data)

    # if product is logged, ...
    if found:

        # update appearances
        found["appearances"] = found["appearances"] + 1
        # update price
        found["average_price"] = (float(current_product["average_price"]) + found["average_price"]) / found["appearances"]
        # calculate average percentage
        found["average_discount"] = found["average_discount"] / found["appearances"];
        print("found is ", found)

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

load_dotenv()
client = MongoClient(os.getenv("key"))
db = client.penguin_magic.open_box

product = {
    "title": "Power Word: Fall by Matt Sconce (DVD + Gimmicks)",
    "average_discount": 41,
    "average_price": 3.4899999999999998,
    "appearances": 5
}
# pprint(product)
# db.open_box.insert_one(product)

print(os.path.dirname(sys.path[0]))
