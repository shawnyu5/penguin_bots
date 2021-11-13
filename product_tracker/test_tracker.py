import unittest
import os
import json

class TestTracker (unittest.TestCase):

    def setUp(self) -> None:
        print("starting test")

    def test_tracker(self):
        print("adding new url...")
        # Buddha Money Mystery
        os.system("echo \"url = https://www.penguinmagic.com/p/266\" >> .env")
        # sed "$ d" FILE

        print("making a copy of current_product.json")
        os.system("cp -v current_product.json _current_product.json")

        print("running script...")
        # excute script
        os.system("python3 tracker.py")

        # delete newly added URL
        os.system("sed -i '$ d' .env")

        print("checking current_product.json...")
        with open("current_product.json") as file:
            current_product = json.load(file)

            print("reverting current_product.json...")
            os.system("mv _current_product.json current_product.json")

            if current_product["title"].lower() == "buddha money mystery":
                assert True
                return
            else:
                assert False

    def tearDown(self) -> None:
        print("ending test")

if __name__ == "__main__":
    unittest.main()
