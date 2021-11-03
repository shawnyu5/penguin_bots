# !/usr/bin/env python3
# purpose of this file:
# Date: 2021-11-01
# ---------------------------------

def main():
    product = {
            "title": "Jar of pickles",
            "discount_percentage": 50,
            "price": 100
            }

    product["appearence"] = 10
    print(product["appearence"])

if __name__ == "__main__":
    main()

