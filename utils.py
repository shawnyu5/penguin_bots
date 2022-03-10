# !/usr/bin/env python3
# purpose of this file: A collection of utility functions to retrive information from penguinmagic
# Date: 2021-09-23
# ---------------------------------
import csv
from bs4 import BeautifulSoup
import requests

# check if current product is interesting, returns true or false
def if_interested(title: str) -> bool:
    """
    checks if the product title argument is interesting by reading the `not_interested_products.csv`

    Args:
        title (): str

    Returns:
        True: product is intersting
        False: product is not interesting

    """
    with open("/home/shawn/python/penguin_bots/not_interested_products.csv", newline = "") as file:
        # read products from cvs
        product = csv.reader(file, quotechar='|')
        product_titles = []
        for row in product:
            # add all product titles to list
            product_titles.extend(row)

        # convert all index of list to lower case for better comparsion
        for i in range(len(product_titles)):
            product_titles[i] = product_titles[i].lower()

        print("Not interested products:", *product_titles, sep = " \n- ")
        print("\n")
        print("Current product:", title)

        if(title.lower() in product_titles):
            print("Product is not intersting")
            return False
        else:
            print("Product is interesting")
            return True

# the price of product, without $
def get_price(soup) -> float:
    """
    get the price of the product from a soup object

    Args:
        soup (): BeautifulSoup

    Returns:
        float:
            price of product
    """
    try:
        price = soup.find("table", class_ = "product_price_details").strike.text.replace("$", "")
        return float(price)
    except:
        raise ValueError("Product has no price")

# returns the discount percentage without %
def get_discount_percentage(soup) -> int: # type: ignore
    try:
        # amount of discount in percentage
        discount_percent = soup.find("td", class_ = "yousave").text.strip()
        discount_percent = discount_percent.split("(")
        discount_percent = discount_percent[1]
        discount_percent = discount_percent.replace(")", "").replace("%", "")
        return int(discount_percent)
    except:
        raise ValueError("No product discount percentage")

# find discounted price without $
def get_discounted_price(soup) -> float:
    try:
        discounted_price = soup.find("td", class_ = "ourprice").text.strip().replace("$", "")
        # print(discounted_price)
        return float(discounted_price)
    except:
        raise ValueError("Product has no discount price")


# returns the number of stars
def get_rating(soup) -> int:
    try:
        review = soup.find("div", id = "review_summary").img["src"].split("/") #type: ignore
        review = review[-1]
        rating = int(review[0])
        return int(rating)
    except:
        raise ValueError("No product rating")

# find product title
def get_title(soup) -> str:
    try:
        title = soup.find("div", id = "product_name").h1.text #type:ignore
        title = title.replace("\"", '\\"')
        return str(title)
    except:
        raise ValueError("Product has no title")

# returns the description of the product
def get_description(soup) -> str:
    try:
        description = soup.find("div", class_ = "product_subsection").text.format() #type: ignore

        # escape all "
        description = description.replace("\"", '\\"')
        return description
    except:
        raise ValueError("Product has no description")

def add_not_interested(product_title) -> None:
    # Parse file to make sure current product is not on not interested list
    with open("/home/shawn/python/web_scraping/penguin_bots/not_interested_products.csv", "r") as file:
        contents = file.read()
        if product_title in contents:
            print(f"{product_title} is already on not interested list. Aborting")
            exit(1)

    choice = input(f"Would you like to add {product_title} to the not interested list?(y/n) ")
    if choice == ('y' or 'Y'):
        with open("/home/shawn/python/web_scraping/penguin_bots/not_interested_products.csv", "a") as file:
            file.write(product_title + "\n")
            print(f"{product_title} addded to not interested list")
    else:
        print("Aborted")
        exit(1)

def main():
    # Play Money by Nick Diffatte, 50% off
    url = "https://www.penguinmagic.com/p/3901"
    page_html = requests.get(url).content
    soup = BeautifulSoup(page_html, "html.parser")

    # add_not_interested("X Finger by Geoff Weber (DVD)")

if __name__ == "__main__":
    main()

