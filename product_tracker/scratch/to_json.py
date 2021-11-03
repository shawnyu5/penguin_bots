import json

arr = []
arr.append({"title": "Jar of pickles",
            "price": 10000
            })
arr.append({"title": "Picke sandwich",
            "price": 10
            })
with open("file.json", "w") as file:
    json.dump(arr, file, indent=4)
