import subprocess
import unittest
import os


class TestTracker(unittest.TestCase):
    def setUp(self) -> None:
        print("starting test...")

    # validates the output of the tracker script
    def __validate_output(self, output):
        print(output)
        if (
            "'appearances': 2".encode() and "'average_discount': 50.0".encode()
        ) in output:
            assert True
            return True
        else:
            print("Outputs do no match...")
            assert False

    def test_ratings(self):
        print("writing to end of file...")
        # append to end of .env file
        with open(".env", "a") as file:
            file.write('url = "https://www.penguinmagic.com/p/3901"\n')
            file.write("dev = True\n")

        # capture output of script
        process = subprocess.run(
            "python3 tracker.py", shell=True, stdout=subprocess.PIPE
        )
        output = process.stdout

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
