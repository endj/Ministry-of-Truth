import sys, os, json
from typing import List
from dataclasses import asdict
from random import choice, randint, sample
from common import Config,Profile,ProfilePersonality, ProfileInterests, ProfileBackground, ProfileCommunicationStyle, ProfileSocialConnections,ConfigMetadata, ConfigTraits, ConfigSocial, ConfigInterests, ConfigBackground, ConfigCommunication, ConfigPersonality, ConfigDemographics
from common import exit, read_conf



def generate_random_profiles(config: Config) -> List[Profile]:
    profiles = []
    print(config.metadata.number_of_profiles)
    for _ in range(config.metadata.number_of_profiles):
        name = choice(config.traits.demographics.name_patterns)

        age = randint(config.traits.demographics.age_range[0], config.traits.demographics.age_range[1])

        occupation = choice(config.traits.demographics.occupations)

        personality = ProfilePersonality(
            type=choice(config.traits.personality.types),
            humor_style=choice(config.traits.personality.humor_styles),
            social_preference=choice(config.traits.personality.social_preferences)
        )

        interests = ProfileInterests(
            primary_interests= sample(config.traits.interests.primary, k=randint(1, len(config.traits.interests.primary))),
            secondary_interests= sample(config.traits.interests.secondary, k=randint(1, len(config.traits.interests.secondary))),
            preferred_topics= sample(config.traits.interests.preferred_topics, k=randint(1, len(config.traits.interests.preferred_topics))),
            disliked_topics=sample(config.traits.interests.disliked_topics, k=randint(1, len(config.traits.interests.disliked_topics)))
        )

        background = ProfileBackground(
            hometown=choice(config.traits.background.hometowns),
            education_level=choice(config.traits.background.education_levels),
            values=sample(config.traits.background.values, k=randint(1, len(config.traits.background.values)))
        )

        communication_style = ProfileCommunicationStyle(
            favorite_words=sample(config.traits.communication.favorite_words, k=randint(1, len(config.traits.communication.favorite_words))),
            formality_level=choice(config.traits.communication.formality_levels),
            conversation_tendency=choice(config.traits.communication.conversation_tendencies)
        )

        social_connections = ProfileSocialConnections(
            groups=sample(config.traits.social.groups, k=randint(1, len(config.traits.social.groups))),
            friendliness_rating=choice(config.traits.social.friendliness_ratings)
        )

        profile = Profile(
            name=name,
            age=age,
            occupation=occupation,
            personality=personality,
            interests=interests,
            background=background,
            communication_style=communication_style,
            social_connections=social_connections
        )

        profiles.append(profile)

    return profiles

def generate_profiles():
    if len(sys.argv) != 2:
        exit(f'Expected python3 template_generator.py <config_file>, got {sys.argv}')

    [_, config_file] = sys.argv
    print(_, config_file)

    if not os.path.isfile(config_file):
        exit(f'{config_file} is not a file')

    config_data = read_conf(config_file)

    config = Config(
        metadata=ConfigMetadata(**config_data["metadata"]),
        traits=ConfigTraits(
            demographics=ConfigDemographics(**config_data["traits"]["demographics"]),
            personality=ConfigPersonality(**config_data["traits"]["personality"]),
            interests=ConfigInterests(**config_data["traits"]["interests"]),
            background=ConfigBackground(**config_data["traits"]["background"]),
            communication=ConfigCommunication(**config_data["traits"]["communication"]),
            social=ConfigSocial(**config_data["traits"]["social"])
        )
    )
    profiles = generate_random_profiles(config)
    profiles_dict = [asdict(p) for p in profiles]
    with open("generated_templates/templates.json", "w") as f:
        json.dump(profiles_dict, f, indent=2)
    print(f'Wrote {len(profiles)} profiles to generated_templates/templates.json')

if __name__ == "__main__":
    generate_profiles()

