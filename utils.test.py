import unittest
import requests
from bs4 import BeautifulSoup
import utils


class TestUtils(unittest.TestCase):

    """Test the utils api"""

    def setUp(self):
        url = "https://www.penguinmagic.com/p/3901"
        page_html = requests.get(url).content
        global soup
        soup = BeautifulSoup(page_html, "html.parser")

    def tearDown(self):
        pass

    def test_get_discount_percentage(self):
        percentage = utils.get_discount_percentage(soup)
        self.assertEqual(percentage, 50, "Discount percentage failed")

    def test_get_price(self):
        price = utils.get_price(soup)
        self.assertEqual(price, 10.00, "get price failed")

    def test_get_discounted_price(self):
        discount_price = utils.get_discounted_price(soup)
        self.assertEqual(discount_price, 4.95, "discounted price falied")

    def test_rating(self):
        rating = utils.get_rating(soup)
        self.assertEqual(rating, 4)

    def test_title(self):
        title = utils.get_title(soup)
        self.assertEqual(title, "Play Money by Nick Diffatte (Instant Download)")

    def test_if_interested(self):
        title = "fjskdlkfja"
        self.assertFalse(utils.if_interested(title))


if __name__ == "__main__":
    unittest.main()
