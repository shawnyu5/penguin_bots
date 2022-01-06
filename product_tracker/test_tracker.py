import json
from pprint import pprint
import subprocess
import unittest
import os
import tracker
from dotenv import load_dotenv


class TestTracker(unittest.TestCase):
    def setUp(self) -> None:
        print("starting test...")

    # validates the output of the tracker script
    def __validate_output(self, output):
        #  try:
        #  output.index("'appearances': 2")
        #  except ValueError:
        #  print("Not found")
        #  else:
        #  print("Found!!")
        #  #  if "'appearances': 2" and "'average_discount': 50.0" in output:
        str = "'appearances': 2"
        #  output = ""
        #  output = json.load(output)
        print(output)
        if str.encode() in output:
            assert True
            return True
        else:
            assert False

    def test_ratings(self):
        print("writing to end of file...")
        # append to end of .env file
        with open(".env", "a") as file:
            file.write('url = "https://www.penguinmagic.com/p/3901"\n')
            file.write("dev = True")

        process = subprocess.run(
            "python3 tracker.py", shell=True, stdout=subprocess.PIPE
        )
        output = process.stdout

        #  print(output)

        try:
            self.__validate_output(output)
        finally:
            print("deleting last line from file...")
            # delete the last 2 lines
            os.system("sed -i '$ d' .env")
            os.system("sed -i '$ d' .env")

    def tearDown(self) -> None:
        print("ending test...")


if __name__ == "__main__":
    unittest.main()
