# Test

```
curl -X POST http://localhost:8080/profiles \
     -H "Content-Type: application/json" \
     -d '{
       "traits": {
         "external_id": 8360300361548484476,
         "name": "Jane Doe",
         "info": "A passionate learner",
         "age": 30,
         "occupation": "Engineer",
         "personality": {
           "type"             : "adventurous",
           "humor_style"      : "dry"        ,
           "social_preference": "lone wolf"
         },
         "interests": {
           "primary_interests"  : ["coding" , "gaming"],
           "secondary_interests": ["reading"          ]
         },
         "background": {
           "hometown": "Berlin",
           "education_level": "PhD",
           "values": ["innovation"]
         },
         "communication_style": {
           "favorite_words": ["amazing", "challenging"],
           "formality_level": "informal",
           "conversation_tendency": "talker"
         },
         "social_connections": {
           "groups": ["Tech Enthusiasts"],
           "friendliness_rating": "friendly"
         }
       }
     }'
```
