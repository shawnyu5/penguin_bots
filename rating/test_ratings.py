import unittest
import ratings  # type: ignore
import requests
from bs4 import BeautifulSoup
import sys
sys.path.insert(1, '/home/shawn/python/web_scraping/penguin')   # add directory to system path in this program
import utils as penguin # type: ignore


class TestRatings(unittest.TestCase):

    def setUp(self):
        print("starting test")

    def test_valid_product(self):

        print("test valid product")
        # open box
        # url = "https://www.penguinmagic.com/openbox"

        # Play Money by Nick Diffatte (Instant Download)
        url = "https://www.penguinmagic.com/p/3901"
        # get web html
        html_page = requests.get(url).text
        soup = BeautifulSoup(html_page, "html.parser")

        product_price = penguin.get_price(soup)
        product_title = penguin.get_title(soup)
        product_rating = penguin.get_rating(soup)
        product_discounted_price = penguin.get_discounted_price(soup)
        # product_description = penguin.get_description(soup)
        product_discount = penguin.get_discount_percentage(soup)

        try:
            if sys.argv[1] == "noti":
                penguin.add_not_interested(product_title)
                exit(0)
        except:
            pass

        if product_rating < 4:
            print(f"Product rating does not fit requirements\nProduct rating: {product_rating}")
            assert False

        if product_discount < 50:
            print(f"Product discount does not fit requirements\nProduct discount: {product_discount}%")
            assert False

        # if product_discount
        if float(product_discounted_price) > 20:
            print(f"Product price not fit requirements\nProduct discounted price: {product_price}")
            assert False

        if not penguin.if_interested(product_title):
            assert False

        assert True

    def test_not_valid_product(self):

        print("testing not valid product")

        # Chi Touch by Christopher Taylor (Instant Download
        url = "https://www.penguinmagic.com/p/15548"

        # get web html
        html_page = requests.get(url).text
        soup = BeautifulSoup(html_page, "html.parser")

        product_price = penguin.get_price(soup)
        product_title = penguin.get_title(soup)
        product_rating = penguin.get_rating(soup)
        product_discounted_price = penguin.get_discounted_price(soup)
        # product_description = penguin.get_description(soup)
        product_discount = penguin.get_discount_percentage(soup)

        print("discount is", product_discount)

        try:
            if sys.argv[1] == "noti":
                penguin.add_not_interested(product_title)
                exit(0)
        except:
            pass

        if product_rating < 4:
            print(f"Product rating does not fit requirements\nProduct rating: {product_rating}")
            assert False

        if product_discount < 50:
            print(f"Product discount does not fit requirements\nProduct discount: {product_discount}%")
            assert True
            return

        # if product_discount
        if float(product_discounted_price) > 20:
            print(f"Product price not fit requirements\nProduct discounted price: {product_price}")
            assert False

        assert False

    def tearDown(self):
        pass

if __name__ == "__main__":
    unittest.main()
