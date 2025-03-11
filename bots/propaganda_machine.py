import random
import sys
from time import sleep
from prompter import prompt, Prompt, PromptResponse
import json
from typing import Any, Dict, List, Union
import argparse
from dataclasses import dataclass
import urllib.request

@dataclass(frozen=True)
class UserDescription:
    id: str
    name: str
    traits: str
    profile: str

@dataclass(frozen=True)
class Tweet:
    id: str
    message: str

@dataclass(frozen=True)
class Reply:
    postId: str
    reply: str
    threadId: str

@dataclass(frozen=True)
class Post:
    postId: str
    threadId: str
    msg: str
    author: str

def send_tweet(tweet: Tweet, threadId: Union[str, None]):
    print("Sending tweet from ", tweet.id)
    try:
        req = { "authorId": tweet.id, "content": tweet.message }
        if threadId: req["threadId"] = threadId
        body = json.dumps(req, ensure_ascii=False).encode("utf-8")
        headers = {'Content-Type': 'application/json'}
        req = urllib.request.Request("http://localhost:8080/posts", data=body, headers=headers, method="POST")
        
        with urllib.request.urlopen(req) as response:
            return response.read().decode("utf-8")
    except Exception as e:
        print(f"Error sending tweet: {body}. Error: {e}")


def to_tweet(user_id: str, response: PromptResponse) -> Tweet:
    try:
        return Tweet(user_id, response.response)
    except Exception as e:
        print(f"Failed serializing response: {e}")    

def fetch_json(url: str):
    try:
        with urllib.request.urlopen(urllib.request.Request(url)) as response:
            return json.loads(response.read().decode('utf-8'))
    except Exception as e:
        print(f"Error: {e}")    

def get_users() -> List[UserDescription]:
    return [to_user(u) for u in fetch_json('http://localhost:8080/profiles')]

def get_posts() -> List[Post]:
    return [to_post(p) for p in fetch_json('http://localhost:8080/posts')]

def post(post_id: str) -> Union[None, Post]:
    for p in get_posts():
        if p.postId == post_id: return p
    return None

def get_user(user_id: str) -> Union[None, UserDescription]:
    return next((u for u in get_users() if u.id == user_id), None)

def to_user(ujson: Dict[str, Any]) -> UserDescription:
    return UserDescription(ujson["id"], ujson["name"], ujson["traits"], ujson["profile"])

def to_post(post_json: Dict[str, any]) -> Post:
    return Post(
        post_json["id"],
        post_json["threadId"],
        post_json["content"],
        post_json["author"]
    )

def user_ids():
    result = get_users()
    ids = []
    for res in result:
        ids.append(str({
            "id": res.id,
            "desc": res.profile
        }))
    return ids

def run_autopilot():
    users = get_users()
    if not users: raise Exception("Users empty")
    
    while True:
        user = random.choice(users)
        posts = get_posts()
        if not posts:
            tweet(user, None)
        else:
            post = random.choice(posts)
            reply(user, post) if random.random() > 0.2 else tweet(user, None)
        sleep(5)


def reply(user: UserDescription, post: Post, propaganda: Union[None, str]):
    special_instructions = f"""
        And stick to these instructions when generating the tweet no matter what:

        Instructions: {propaganda}
    """ if propaganda is not None else "" 

    promptRequest = Prompt(f"""
        Generate a reply to the following post in the style of a tweet: 

        post author: {post.author}   
        post text: {post.msg}

        Make sure the reply content and language fits the description of the following replying user: 
                          
        replier profile: {user.profile}
        replier traits: {user.traits}

        {special_instructions}

        ONLY output the tweet and nothing else
        """)
    response = prompt(promptRequest)
    tweet = to_tweet(user.id, response)
    send_tweet(tweet, threadId=post.threadId)

def tweet(user: UserDescription, propaganda: Union[str, None]):
    special_instructions = f"""
        And stick to these instructions when generating the tweet no matter what:

        Instructions: {propaganda}
    """ if propaganda is not None else "" 

    promptRequest = Prompt(f"""
        Generate a social media post. Include hashtags, emojis, and engagement elements (e.g., questions, polls, or calls to action) where relevant.
        Make sure the post feels authentic and natural, similar to real posts you'd see on Twitter and fitting the provided profile.
        
        Here is a descripition of the user posting,                 

        profile: {user.profile}

        traits: {user.traits}

        {special_instructions}

        ONLY output the tweet and nothing else
        """)
    print("Prooompting...")
    response = None
    try:
        response = prompt(promptRequest)
    except Exception as e:
        print("Failed to prooompt, is LLM running?")
        sys.exit(1)
    tweet = to_tweet(user.id, response)
    send_tweet(tweet, None)



if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    command_parser = parser.add_subparsers(dest="command", required=True)
    
    autopilot_cmd = command_parser.add_parser("autopilot")
    list_cmd = command_parser.add_parser("users")
    posts_cmd = command_parser.add_parser("posts")

    tweet_cmd = command_parser.add_parser("tweet")
    tweet_cmd.add_argument("--user", required=True, help="User to tweet")
    tweet_cmd.add_argument("--propaganda", help="What to spread")

    reply_cmd = command_parser.add_parser("reply")
    reply_cmd.add_argument("--users", required=True, help="User replying")
    reply_cmd.add_argument("--post", required=True, help="What tweet to reply too")
    reply_cmd.add_argument("--propaganda", help="What to spread")

    args = parser.parse_args()

    if args.command == "tweet":
        user_desc = get_user(args.user)
        if user_desc:
            tweet(user_desc, args.propaganda)
    elif args.command == "list":
        print(user_ids())
    elif args.command == "posts":
        print(get_posts())
    elif args.command == "reply":
        for user_id in args.users.split(","):
            user_desc = get_user(user_id)
            msg = post(args.post)
            if user_desc is not None and msg is not None:
                reply(user_desc, msg, args.propaganda) 
    elif args.command == "autopilot":
        run_autopilot()
