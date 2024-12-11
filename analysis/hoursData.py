import json

# Define the path to your JSON file
file_path = 'testHours.json'

# Function to load data from JSON
def load_data(file_path):
    try:
        with open(file_path, 'r') as file:
            data = json.load(file)
        # Ensure the JSON is a list
        if not isinstance(data, list):
            data = [data]
        return data
    except FileNotFoundError:
        print(f"Error: The file '{file_path}' was not found.")
        return []
    except json.JSONDecodeError as e:
        print(f"Error decoding JSON: {e}")
        return []

# Function to summarize worker hours and tasks
def summarize_worker_hours_and_tasks(data):
    summary = {}

    for entry in data:
        # Extract worker details
        worker = entry.get("worker", {})
        worker_name = worker.get("name", "Unknown Worker")
        duration = entry.get("duration", 0)

        # Initialize task list
        task_names = []

        # Extract task details from 'process -> harvest -> crop'
        process = entry.get("process", {})
        harvest_in_process = process.get("harvest", {})
        crop_in_process = harvest_in_process.get("crop", {})
        if crop_in_process:
            process_task_name = crop_in_process.get("name")
            if process_task_name:
                task_names.append(f"Processing {process_task_name}")

        # Extract task details from 'harvest -> crop'
        harvest = entry.get("harvest", {})
        crop_in_harvest = harvest.get("crop", {})
        if crop_in_harvest:
            harvest_task_name = crop_in_harvest.get("name")
            if harvest_task_name:
                task_names.append(f"Harvesting {harvest_task_name}")

        # Fallback if no task names are found
        if not task_names:
            task_names.append("Unknown Task")

        # Initialize worker in the summary if not already present
        if worker_name not in summary:
            summary[worker_name] = {"total_time": 0, "tasks": []}

        # Update worker's total time and task list
        summary[worker_name]["total_time"] += duration
        for task_name in task_names:
            summary[worker_name]["tasks"].append({"task": task_name, "time": duration})

    return summary

# Function to get worker hours and tasks for a specific worker
def get_worker_hours(worker_name, data):
    total_hours = 0
    task_list = []

    for entry in data:
        if entry['worker']['name'] == worker_name:
            total_hours += entry['duration']
            
            # Add the process or harvest task details
            task = ""
            if 'process' in entry:
                task += f"Process: {entry['process']['product']['name']} (Qty: {entry['process']['qty']})"
            if 'harvest' in entry:
                task += f" | Harvest: {entry['harvest']['crop']['name']} (Weight: {entry['harvest']['weight']} kg)"
            
            task_list.append(f"{task} - Duration: {entry['duration']} hours")

    return total_hours, task_list

# Main script to load data and summarize specific worker's hours and tasks
data = load_data(file_path)

# Example: Get total hours worked and tasks for "Ginny Gambie" will make this an input you can do from command line
worker_name = "Ginny Gambie"
total_hours, tasks = get_worker_hours(worker_name, data)

# Display the results for Ginny Gambie
print(f"Total hours worked by {worker_name}: {total_hours:.2f} hours")
print("Tasks:")
for task in tasks:
    print(f"  - {task}")

# # Summarize and display all workers' hours and tasks
# summary = summarize_worker_hours_and_tasks(data)
# for worker, details in summary.items():
#     print(f"\nWorker: {worker}")
#     print(f"Total Time: {details['total_time']:.2f} hours")
#     print("Tasks:")
#     for task in details["tasks"]:
#         print(f"  - {task['task']}: {task['time']:.2f} hours")

# TODO: add function to look up crop harvests and time put into the crop and who worked on a specific crop