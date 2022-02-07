import subprocess
import unittest
import os
from tracker import Tracker
from unittest import mock
from pprint import pprint
from io import StringIO
import contextlib
from bson.objectid import ObjectId


class TestTracker(unittest.TestCase):
    def setUp(self) -> None:
        print("starting test...")
        global tracker
        tracker = Tracker()

    # use play money for testing
    @mock.patch.dict(
        os.environ, {"url": "https://www.penguinmagic.com/p/3901", "dev": "true"}
    )
    def test_save(self):
        tracker.get_product_info()
        saved_product = tracker.save()
        #  if (
            #  saved_product
            #  == """{'appearances': 1,
    #  'average_discount': 50,
    #  'average_price': 4.95,
    #  'title': 'Play Money by Nick Diffatte (Instant Download)'}"""
        #  ):
            #  assert True
        #  else:
            #  assert False
        self.assertEqual(
            saved_product,
            {
                "_id": ObjectId("61dceb6228b23db27260d4e0"),
                "appearances": 1,
                "average_discount": 33.333333333333336,
                "average_price": 3.3000000000000003,
                "title": "Play Money by Nick Diffatte (Instant Download)",
            },
        )

        #  mock_print.assert_called_with("Product updated")
        #  mock_print.assert_called_with(
        #  """product updated

    #  {'appearances': 1,
    #  'average_discount': 50,
    #  'average_price': 4.95,
    #  'title': 'Play Money by Nick Diffatte (Instant Download)'}
    #  """
    #  )

    # # validates the output of the tracker script
    # def __validate_output(self, output):
    # print(output)
    # if (
    # "'appearances': 2".encode() and "'average_discount': 50.0".encode()
    # ) in output:
    # assert True
    # return True
    # else:
    # print("Outputs do no match...")
    # assert False

    # def test_ratings(self):
    # print("writing to end of file...")
    # # append to end of .env file
    # with open(".env", "a") as file:
    # file.write('url = "https://www.penguinmagic.com/p/3901"\n')
    # file.write("dev = True\n")

    # # capture output of script
    # process = subprocess.run(
    # "python3 tracker.py", shell=True, stdout=subprocess.PIPE
    # )
    # output = process.stdout

    # try:
    # self.__validate_output(output)
    # finally:
    # print("deleting last line from file...")
    # # delete the last 2 lines
    # os.system("sed -i '$ d' .env")
    # os.system("sed -i '$ d' .env")

    def tearDown(self) -> None:
        print("ending test...")


if __name__ == "__main__":
    unittest.main()
