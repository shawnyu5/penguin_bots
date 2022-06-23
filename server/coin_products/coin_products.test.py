import unittest
from unittest import TestCase, mock
import os


class TestTracker(unittest.TestCase):
    def setUp(self) -> None:
        print("Starting test...")

    def tearDown(self) -> None:
        print("Ending test...")

    @mock.patch.dict(os.environ, {"FROBNICATION_COLOUR": "ROUGE"})
    def test_description(self):
        pass

if __name__ == "__main__":
    unittest.main()

