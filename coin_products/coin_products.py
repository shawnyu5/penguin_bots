# purpose of this file: scrape penguin magic open box section for coin products
# Date: 2021-09-03
# ---------------------------------
import requests
from bs4 import BeautifulSoup
import os
import sys
import json

sys.path.insert(
    1, "/home/shawn/python/penguin_bots/"
)  # add directory to system path in this program
import utils  # type: ignore

# gets url from .env and returns a soup object
def get_webpage(product: dict):
    url = product["url"]
    # get web html
    html_page = requests.get(url).text
    return BeautifulSoup(html_page, "html.parser")


def validate(product):
    # make sure it's a coin product
    not_coin_product = lambda title: (print(title, "is not a coin product"), exit(1))

    if ("coin" or "coins") not in product["description"].lower() or (
        "coin" or "coins"
    ) not in product["title"].lower():
        not_coin_product(product["title"])

    with open("product_info.txt", "+r") as file:
        if file.read() != product["title"]:
            print("product changed")

            # overwrite current product title
            with open("product_info.txt", "w") as file:
                file.write(product["title"])
        else:
            print("product has not changed")


def get_product_info(soup, product):
    try:
        # find product title
        product["title"] = soup.find("div", id="product_name").h1.text  # type:ignore
        # product description
        product["description"] = soup.find("div", class_="product_subsection").text.format()  # type: ignore
    except:
        print("There are no open box product currently")

    # escape all "
    product["description"] = product["description"].replace('"', '\\"')


def main():

    product = {"title": str, "description": None, "url": None}

    product["url"] = "https://www.penguinmagic.com/openbox/"

    # create soup object
    soup = get_webpage(product=product)

    # get product title and description
    get_product_info(soup, product)

    if not utils.if_interested(product["title"]):
        print("Product is not interesting...")
        exit(1)

    # make sure current product is different from previous product sent in email
    validate(product)

    print(json.dumps(product))


if __name__ == "__main__":
    main()
