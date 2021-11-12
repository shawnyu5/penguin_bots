from dotenv import load_dotenv
import os

load_dotenv();

name = os.getenv("name")
name = "adam"

print(name)

