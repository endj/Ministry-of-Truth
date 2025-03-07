import os
import json
import urllib.request
import urllib.response

def process_folder(folder_path):
    try:
        with open(os.path.join(folder_path, "template.json"), "r") as f:
            traits = json.load(f)
        with open(os.path.join(folder_path, "profile.json"), "r") as f:
            profiles = json.load(f)
        return {"traits": traits, "profile": profiles}
    except FileNotFoundError as e:
        print(f"Error: {e}. Skipping folder: {folder_path}")
        return None

def main():
    folder_path = "."  # Process folders in the current directory

    for root, dirs, files in os.walk(folder_path):
        if "profile.json" in files and "template.json" in files:
            folder_data = process_folder(root)
            if folder_data:
                try:
                    data = json.dumps(folder_data).encode("utf-8")
                    headers = {'Content-Type': 'application/json'}
                    print("Sending data", data)
                    req = urllib.request.Request(
                        "http://localhost:8080/profiles",
                        data=data,
                        headers=headers,
                        method="POST"
                    )
                    with urllib.request.urlopen(req) as response:
                        result = response.read().decode("utf-8")
                        print(f"Successfully sent data for folder: {root} " + result)
                except Exception as e:
                    print(f"Error sending request for folder: {root}. Error: {e}")

if __name__ == "__main__":
    main()
