#!/usr/bin/env python3
#purpose of this file: scrape penguin open box products for discount more than 50% and more then 4 star rating
#Date: 2021-09-09
#---------------------------------
import requests
from bs4 import BeautifulSoup
import os
from dotenv import load_dotenv
import sys
import time

sys.path.insert(1, '/home/shawn/python/web_scraping/penguin')   # add directory to system path in this program
import utils as penguin # type: ignore

# makes sure the current product is different from the product stored in file
def validate_product(product_title):
    # store product title to file
    try:
        with open("/home/shawn/python/web_scraping/penguin/rating/product_info.txt", "r+") as file:
            # read title from file
            file_title = file.read()
    except:
        file_title = "xxxxx"
        pass

    if file_title == product_title:
        print("product has not changed")
        exit(1)

    # if current title is different from previous, write to file
    with open("/home/shawn/python/web_scraping/penguin/rating/product_info.txt", "w+") as file:
        # print(f"new product title {product_title} written")
        file.write(product_title)

def main():

    start = time.perf_counter()
    load_dotenv()
    url = str(os.environ.get("url"))

    # get web html
    html_page = requests.get(url).text
    soup = BeautifulSoup(html_page, "html.parser")

    try:
        product_price = penguin.get_price(soup)
        product_title = penguin.get_title(soup)
        product_rating = penguin.get_rating(soup)
        product_discounted_price = penguin.get_discounted_price(soup)
        product_description = penguin.get_description(soup)
        product_discount = penguin.get_discount_percentage(soup)
    except Exception as error:
        print(error)
        exit(1)

    not_interested = False

    try:
        if sys.argv[1] == "noti":
            not_interested = True
    except:
        pass

    if not_interested:
        penguin.add_not_interested(product_title)
        exit(0)

    if product_rating < 4:
        print(f"Product rating does not fit requirements\nProduct rating: {product_rating}")
        exit(1)

    if product_discount < 50:
        print(f"Product discount does not fit requirements\nProduct discount: {product_discount}%")
        exit(1)

    # if product_discount
    if float(product_discounted_price) > 20:
        print(f"Product price not fit requirements\nProduct discounted price: {product_price}")
        exit(1)

    if not penguin.if_interested(product_title):
        exit(1)

    # make sure the current product is different than the one sent by email
    validate_product(product_title)

    # retrieve recipients from .env file
    recipients = str(os.getenv("recipients"))

    # Note: must escape $ in bash echo command
    command = (f"echo \"***{product_title}***is available at penguin right now!!!\n\nDiscount percentage: {product_discount}%\nOrigional price: \${product_price}\nCurrent price: \${product_discounted_price}\nStars: {product_rating}\n\nDescription: {product_description}\n\nURL: {url}\" | neomutt -s \"Heavily discounted product at penguin right now!!!\" \"{recipients}\" &> /dev/null")

    os.system(command)
    print("good product")
    print("email sent")

    end = time.perf_counter()

    print(f"this scrip took {end - start} seconds:")
if __name__ == "__main__":
    main()
