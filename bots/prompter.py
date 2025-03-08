from dataclasses import dataclass
import json
from typing import List
import http.client


@dataclass
class PromptResponse:
    model: str
    created_at: str
    response: str
    done: bool
    done_reason: str
    context: List[int]
    total_duration: int
    load_duration: int
    prompt_eval_count: int
    prompt_eval_duration: int
    eval_count: int
    eval_duration: int

@dataclass
class Prompt:
    prompt: str
    model: str  = "llama3.2"

def prompt(prompt: Prompt) -> PromptResponse:
    body = json.dumps({
        "model": prompt.model,
        "prompt": prompt.prompt,
        "stream": False
    })

    conn = http.client.HTTPConnection("localhost", 11434)
    conn.request("POST", "/api/generate", body, {"Content-Type": "application/json"})

    response = conn.getresponse()
    response_data = response.read()
    response_body = response_data.decode()

    conn.close()

    return PromptResponse(
        **json.loads(response_body)
    )