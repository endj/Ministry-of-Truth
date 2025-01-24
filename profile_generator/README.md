# Generating Profiles

## TODO

* Move config.json up
* Add probability distributions to traits
* Move to new LLM model
* Find scriptable picture Model



Two step process:

## Creating the profiles

1. Define the traits users can have and some metadata profile_generator/config.json
2. Run template_generator.py to generate templates from traits
3. Start LLM using ollama or similar
4. Run profile_generator.py to generate profiles from templates
   4.1 generator picks up 1 template from templates files
   4.2 generates a profile
   4.3 creates a folder named after hash(info, name)
   4.4 writes profile and template to folder

## Creating profile pictures

TODO

