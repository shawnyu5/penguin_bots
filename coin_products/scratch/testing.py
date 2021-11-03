#purpose of this file: test system commands
#Date: 2021-09-03
#---------------------------------
import os

def main():
    product = "Coinductor"
    if ("coin " or "coins ") in product.lower():
        print("yessss")
    else:
        print("noooo")

if __name__ == "__main__":
    main()

