from bs4 import BeautifulSoup
import requests
from dotenv import load_dotenv
import json

url = "https://www.penguinmagic.com/p/266"
print(url)
#  return

html_page = requests.get(url).text
soup = BeautifulSoup(html_page, "html.parser")

title = soup.find("div", { "id": "product_main_details" }).text
print(title)

