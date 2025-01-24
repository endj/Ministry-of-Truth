import os
import json
import requests

def process_folder(folder_path):
    """
    Processes a folder containing profile.json and template.json.

    Args:
        folder_path: Path to the folder.

    Returns:
        A dictionary containing 'traits' and 'profiles' keys.
    """

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
    """
    Processes all folders in the current directory and sends POST requests.
    """

    folder_path = "."  # Process folders in the current directory

    for root, dirs, files in os.walk(folder_path):
        if "profile.json" in files and "template.json" in files:
            folder_data = process_folder(root)
            if folder_data:
                try:
                    response = requests.post("http://localhost:8080/profiles", json=folder_data)
                    response.raise_for_status()  # Raise an exception for bad status codes
                    print(f"Successfully sent data for folder: {root}")
                except requests.exceptions.RequestException as e:
                    print(f"Error sending request for folder: {root}. Error: {e}")

if __name__ == "__main__":
    main()
