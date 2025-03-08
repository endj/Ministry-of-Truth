from dataclasses import dataclass
from typing import Tuple, List, Any
import json
import sys


def exit(msg: str, code = 1):
    print(msg)
    sys.exit(code)

def read_conf(file_path: str) -> Any:
    try:
        with open(file_path) as file:
            return json.load(file)
    except json.JSONDecodeError:
        exit(f'File {file_path} could not be decoded as json')
    except FileNotFoundError:
        exit(f'File {file_path} not found')

@dataclass
class ConfigMetadata:
    number_of_profiles: int

@dataclass
class ConfigDemographics:
    name_patterns: List[str]
    age_range: Tuple[int, int]
    occupations: List[str]

@dataclass
class ConfigPersonality:
    types: List[str]
    humor_styles: List[str]
    social_preferences: List[str]

@dataclass
class ConfigInterests:
    primary: List[str]
    secondary: List[str]
    preferred_topics: List[str]
    disliked_topics: List[str]

@dataclass
class ConfigBackground:
    hometowns: List[str]
    education_levels: List[str]
    values: List[str]

@dataclass
class ConfigCommunication:
    favorite_words: List[str]
    formality_levels: List[str]
    conversation_tendencies: List[str]

@dataclass
class ConfigSocial:
    groups: List[str]
    friendliness_ratings: List[str]

@dataclass
class ConfigTraits:
    demographics: ConfigDemographics
    personality: ConfigPersonality
    interests: ConfigInterests
    background: ConfigBackground
    communication: ConfigCommunication
    social: ConfigSocial

@dataclass
class Config:
    metadata: ConfigMetadata
    traits: ConfigTraits

@dataclass
class ProfilePersonality:
    type: str
    humor_style: str
    social_preference: str

@dataclass
class ProfileInterests:
    primary_interests: List[str]
    secondary_interests: List[str]
    preferred_topics: List[str]
    disliked_topics: List[str]

@dataclass
class ProfileBackground:
    hometown: str
    education_level: str
    values: List[str]

@dataclass
class ProfileCommunicationStyle:
    favorite_words: List[str]
    formality_level: str
    conversation_tendency: str

@dataclass
class ProfileSocialConnections:
    groups: List[str]
    friendliness_rating: str

@dataclass
class Profile:
    name: str
    age: int
    occupation: str
    personality: ProfilePersonality
    interests: ProfileInterests
    background: ProfileBackground
    communication_style: ProfileCommunicationStyle
    social_connections: ProfileSocialConnections
