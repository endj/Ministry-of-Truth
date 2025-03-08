from dataclasses import dataclass, asdict
import sys, json
from typing import Any, Union
from common import Profile, ProfilePersonality, ProfileInterests, ProfileBackground, ProfileCommunicationStyle, ProfileSocialConnections
from common import exit, read_conf
import os
from bots.prompter import Prompt, prompt

@dataclass
class GeneratedProfile:
    name: str
    info: str

def try_parse(response: str) -> Union[GeneratedProfile, None]:
    # Expects format name=<name>,profile=<profile>
    name_end = response.find(",profile=")

    if name_end == -1:
        print("Failed to find profile in response:", response)
        return None

    # name=<name>, profile=<profile>
    name_part = response[:name_end]
    profile_part = response[name_end + len(",profile="):]

    if not name_part.startswith("name="):
        print("Invalid format for name:", name_part)
        return None
    name = name_part[len("name="):]

    profile = profile_part.replace('"', '')  # Remove all quotes from profile
    return GeneratedProfile(name, profile)


def profile_from_dict(profile_data: Any) -> Profile:
    return Profile(
        name=profile_data["name"],
        age=profile_data["age"],
        occupation=profile_data["occupation"],
        personality=ProfilePersonality(**profile_data["personality"]),
        interests=ProfileInterests(**profile_data["interests"]),
        background=ProfileBackground(**profile_data["background"]),
        communication_style=ProfileCommunicationStyle(**profile_data["communication_style"]),
        social_connections=ProfileSocialConnections(**profile_data["social_connections"])
    )

def generate_prompt(profile_data: Any):
    prompt = f'''
        Given the following user information, you need to write a social media profile as if a user was writing it
        Make up a name for the user that fits their traits.
        Keep it below 240 characters

        Only output name and profile information in the following format
        name=<name>,profile=<profile>

        Here is the user traits:
            ${json.dumps(profile_data)}
    '''
    return prompt


def next_profile(profile_template: str) -> Union[GeneratedProfile, None]:
    profile = profile_from_dict(profile_template)
    print("Generating profile for user", profile)

    prumpt = generate_prompt(profile_template)
    print("[PROMPT]", prumpt)
    answer = prompt(Prompt(prumpt))
    print("[RESPONSE]",answer.response)
    outcome = try_parse(answer.response)
    print("[OUTCOME]",outcome)
    return outcome


def generate_profiles():
    if len(sys.argv) != 2:
        exit(f'Expected python3 profile_generator <profile_template_file>, expceted 2, got 1 argument, {sys.argv}')

    [_, profiles_file] = sys.argv

    while True:
        profiles = read_conf(profiles_file)
        if not profiles:
            exit(f"All profiles in {profiles_file} generated", 0)
        profile_template = profiles.pop()
        outcome = next_profile(profile_template)

        if not outcome:
            print("Failed to generate profile", profile_template)
            continue

        id = hash((outcome.info, outcome.name))
        folder_name = str(abs(id))
        folder_path = f'generated_profiles/{folder_name}'
        os.mkdir(folder_path)
        with open(f'{folder_path}/profile.json', "w+") as f:
            json.dump(asdict(outcome), f, indent=2)
        with open(f'{folder_path}/template.json', "w+") as f:
            json.dump(profile_template, f, indent=2)
        with open(profiles_file, "w") as f:
            json.dump(profiles, f, indent=2)

if __name__ == "__main__":
    generate_profiles()