# FakeNet

TODO

* Create table for posting and replies
* Create API for creating post
* Create API for replying to post
* Create prompt API for generating posts and replies based on user profile and subject


## Schema
```json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "type": "object",
  "properties": {
    "metadata": {
      "type": "object",
      "properties": {
        "number_of_profiles": {
          "type": "integer",
          "minimum": 1
        }
      }
    },
    "traits": {
      "type": "object",
      "properties": {
        "demographics": {
          "type": "object",
          "properties": {
            "name_patterns": {
              "type": "array",
              "items": { "type": "string" }
            },
            "age_range": {
              "type": "array",
              "items": { "type": "integer" },
              "minItems": 2,
              "maxItems": 2
            },
            "occupations": {
              "type": "array",
              "items": { "type": "string" }
            }
          }
        },
        "personality": {
          "type": "object",
          "properties": {
            "types": {
              "type": "array",
              "items": { "type": "string" }
            },
            "humor_styles": {
              "type": "array",
              "items": { "type": "string" }
            },
            "social_preferences": {
              "type": "array",
              "items": { "type": "string" }
            }
          }
        },
        "interests": {
          "type": "object",
          "properties": {
            "primary": {
              "type": "array",
              "items": { "type": "string" }
            },
            "secondary": {
              "type": "array",
              "items": { "type": "string" }
            },
            "preferred_topics": {
              "type": "array",
              "items": { "type": "string" }
            },
            "disliked_topics": {
              "type": "array",
              "items": { "type": "string" }
            }
          }
        },
        "background": {
          "type": "object",
          "properties": {
            "hometowns": {
              "type": "array",
              "items": { "type": "string" }
            },
            "education_levels": {
              "type": "array",
              "items": { "type": "string" }
            },
            "values": {
              "type": "array",
              "items": { "type": "string" }
            }
          }
        },
        "communication": {
          "type": "object",
          "properties": {
            "favorite_words": {
              "type": "array",
              "items": { "type": "string" }
            },
            "formality_levels": {
              "type": "array",
              "items": { "type": "string" }
            },
            "conversation_tendencies": {
              "type": "array",
              "items": { "type": "string" }
            }
          }
        },
        "social": {
          "type": "object",
          "properties": {
            "groups": {
              "type": "array",
              "items": { "type": "string" }
            },
            "friendliness_ratings": {
              "type": "array",
              "items": { "type": "string" }
            }
          }
        }
      }
    }
  }
}


```
