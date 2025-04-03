import csv
import requests

csvinput = "tasks.csv"
#apiurl = "https://json.tesc.farm"
apiurl = "http://127.0.0.1:8078"

with open(csvinput, 'r') as file:
    csv_reader = csv.reader(file)
    for task in csv_reader:
        print(task)
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
            response = requests.post(apiurl + "/crop/new", json=newcrop)
            if response.status_code == 200:
                data = response.json()
                print("Response:", data)
                planting['crop_id'] = data['ID']
            else:
                print("Error:", response.text)
                quit()
            response = requests.post(apiurl + "/planting/new", json=planting)
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
        response = requests.post(apiurl + "/task/new", json=payload)
        if response.status_code == 200:
            data = response.json()
            print("Response:", response.json())
        else:
            print("Error:", response.text)
            quit()
