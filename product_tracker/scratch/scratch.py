from pymongo import MongoClient
from pprint import pprint
import os
from dotenv import load_dotenv

load_dotenv()
client = MongoClient(os.getenv("key"))
db = client.penguin_magic
product = {
    "title": "Power Word: Fall by Matt Sconce (DVD + Gimmicks)",
    "discount_percent": 41,
    "price": 3.4899999999999998,
    "appearances": 5
}
pprint(product)
db.open_box.insert_one(product)

