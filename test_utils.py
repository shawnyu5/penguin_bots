# set up pytest function
import pytest

# before each
@pytest.fixture(autouse=True)
def run_around_tests():
    print("setup")

def test_1():
    print("test 1")

def test_2():
    print("test2")

