# set up pytest function
import pytest
import requests
from bs4 import BeautifulSoup
import utils

# set up
def setup_module(module):
    """
    Generate soup object for testing

    Args:
        module ():
    """
    print("setup")
    url = "https://www.penguinmagic.com/p/3901"
    page_html = requests.get(url).content
    global soup
    soup = BeautifulSoup(page_html, "html.parser")


class TestUtils:
    #  @pytest.fixture(autouse=True)
    #  def setUp(self):

    def test_get_discount_percentage(self):
        percentage = utils.get_discount_percentage(soup)
        assert percentage == 50

    def test_get_price(self):
        price = utils.get_price(soup)
        assert price == 10.00

    def test_get_discounted_price(self):
        discount_price = utils.get_discounted_price(soup)
        assert discount_price == 4.95

    def test_rating(self):
        rating = utils.get_rating(soup)
        assert rating == 4

    def test_title(self):
        title = utils.get_title(soup)
        assert title == "Play Money by Nick Diffatte (Instant Download)"

    def test_if_interested(self):
        title = "fjskdlkfja"
        assert utils.if_interested(title) == True
