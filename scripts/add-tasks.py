import requests

baseurl = "https://json.tesc.farm"

tasks = [
    ("Apples", "preharvest"),
    ("Artichokes", "preharvest"),
    ("Asian Greens", "preharvest"),
    ("Asian Pears", "preharvest"),
    ("Asparagus", "preharvest"),
    ("Basil", "preharvest"),
    ("Beans", "preharvest"),
    ("Blueberries", "preharvest"),
    ("Bok Choy", "preharvest"),
    ("Broccoli", "preharvest"),
    ("Brussels Sprouts", "preharvest"),
    ("Cabbage", "preharvest"),
    ("Carrots", "preharvest"),
    ("Celeriac", "preharvest"),
    ("Celery", "preharvest"),
    ("Chives", "preharvest"),
    ("Cilantro", "preharvest"),
    ("Collards", "preharvest"),
    ("Cucumbers", "preharvest"),
    ("Dill", "preharvest"),
    ("European Pears", "preharvest"),
    ("Figs", "preharvest"),
    ("Flowers", "preharvest"),
    ("Grapes", "preharvest"),
    ("Kale", "preharvest"),
    ("Kiwiberries", "preharvest"),
    ("Parsley", "preharvest"),
    ("Peas", "preharvest"),
    ("Peppers, Hot", "preharvest"),
    ("Peppers, Sweet", "preharvest"),
    ("Perennial Cut Flowers", "preharvest"),
    ("Plums", "preharvest"),
    ("Potatoes", "preharvest"),
    ("Radicchio", "preharvest"),
    ("Radishes", "preharvest"),
    ("Raspberries", "preharvest"),
    ("Rhubarb", "preharvest"),
    ("Romanesco", "preharvest"),
    ("Salad Mix", "preharvest"),
    ("Strawberries", "preharvest"),
    ("Summer Squash", "preharvest"),
    ("Tomatoes, Cherry", "preharvest"),
    ("Tomatoes, Slicing", "preharvest"),
    ("Winter Squash", "preharvest"),
    ("1990's Orchard", "mu"),
    ("2017 Hedgerow", "mu"),
    ("BLT", "mu"),
    ("Food Forest", "mu"),
    ("Jacob's Orchard", "mu"),
    ("MU1", "mu"),
    ("MU2", "mu"),
    ("MU2a", "mu"),
    ("MU3", "mu"),
    ("MU6", "mu"),
    ("MU8", "mu"),
    ("MU9", "mu"),
    ("General Maintenance", "maintenance")
]

for task in tasks:
    payload = {}
    payload['name'] = task[0]
    payload['tags'] = []
    if task[1] == 'preharvest':
        payload['tags'].append({'name': 'Preharvest'})
        newcrop = {'name': task[0]}
        parts = task[0].split(", ")
        if len(parts) == 2:
            payload['name'] = f"{parts[1]} {parts[0]}"
            newcrop['name'] = parts[0]
            newcrop['variety'] = parts[1]
        newcrop['tags'] = [{'name': parts[0]}]
        planting = {}
        response = requests.post(baseurl + "/crop/new", json=newcrop)
        if response.status_code == 200:
            data = response.json()
            print("Response:", data)
            planting['crop_id'] = data['ID']
        else:
            print("Error:", response.text)
            quit()
        response = requests.post(baseurl + "/planting/new", json=planting)
        if response.status_code == 200:
            data = response.json()
            print("Response:", data)
            payload['planting_id'] = data['ID']
        else:
            print("Error:", response.text)
            quit()
        payload['description'] = "Preharvest " + payload['name']
    elif task[1] == "mu":
        payload['tags'].append({'name': 'Management Unit'})
        payload['description'] = 'Work in ' + payload['name']
    else:
        payload['tags'].append({'name': 'Maintenance'})
    print(payload)
    response = requests.post(baseurl + "/task/new", json=payload)
    if response.status_code == 200:
        data = response.json()
        print("Response:", response.json())
    else:
        print("Error:", response.text)
        quit()
    break # Remove after testing
