#purpose of this file: scrape penguin magic open box section for coin products
#Date: 2021-09-03
#---------------------------------
from dotenv import load_dotenv
import requests
from bs4 import BeautifulSoup
import os
import sys

sys.path.insert(1, '/home/shawn/python//penguin_bots/')   # add directory to system path in this program
import utils # type: ignore

# gets url from .env and returns a soup object
def get_webpage():
    load_dotenv()
    url = str(os.environ.get("url"))

    # get web html
    html_page = requests.get(url).text
    return BeautifulSoup(html_page, "html.parser")

def validate(product):
    # make sure it's a coin product
    not_coin_product = lambda title: (
            print(title, "is not a coin product"),
            exit(1)
            )

    # print(product["description"].lower())
    if ("coin" or "coins") not in product["description"].lower() and ("coin" or "coins") not in product["title"].lower():
        not_coin_product(product["title"])

    with open("product_info.txt", "+r") as file:
        # if file.read() == product_title:
        if file.read() != product["title"]:
            print("product changed")

            # overwrite current product title
            with open("product_info.txt", "w") as file:
                file.write(product["title"])

            send_email(product)
        else:
            print("product has not changed")

def send_email(product: dict):
    load_dotenv()
    recipients = os.environ.get("recipients")

    title = product["title"]
    description = product["description"]
    url = product["url"]

    command = (f"printf \"***{title}***is available at penguin right now!!!\n\nDescription: {description}\n\nURL: {url}\" | neomutt -s \"Coin open box product right now!!!\" \" {recipients}\" &> /dev/null")
    # print("command is ", command)

    os.system(command)
    print("email sent!!!")

def get_product_info(soup, product):
    try:
        # find product title
        product["title"] = soup.find("div", id = "product_name").h1.text #type:ignore
        # product description
        product["description"] = soup.find("div", class_ = "product_subsection").text.format() #type: ignore
    except:
        print("There are no open box product currently")

    # escape all "
    product["description"] = product["description"].replace("\"", '\\"')

def main():

    product = {
        "title": None,
        "description": None,
        "url": None
    }

    load_dotenv()
    product["url"] = str(os.environ.get("url")) #type: ignore

    # create soup object
    soup = get_webpage()

    # get product title and description
    get_product_info(soup, product)

    if not utils.if_interested(product["title"]):
        exit(1)

    # make sure current product is different from previous product sent in email
    validate(product)

if __name__ == "__main__":
    main()

