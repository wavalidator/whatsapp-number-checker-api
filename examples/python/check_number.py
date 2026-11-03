"""Check if a phone number has WhatsApp with the WA Validator API.

Usage:
    pip install requests
    export WAVALIDATOR_API_KEY=your_key
    python check_number.py 14155551234
"""
import os
import sys
import time

import requests

API_URL = "https://wavalidator.com/api/v1/validate/"


def check_whatsapp_number(phone_number: str, api_key: str, retries: int = 3) -> dict:
    for attempt in range(retries):
        response = requests.post(
            API_URL,
            headers={
                "Authorization": f"Bearer {api_key}",
                "Content-Type": "application/json",
            },
            json={"phone_number": phone_number},
            timeout=30,
        )
        if response.status_code == 429:
            time.sleep(int(response.json().get("retry_after", 2 ** attempt)))
            continue
        response.raise_for_status()
        return response.json()
    raise RuntimeError("Rate limit still reached after retries")


if __name__ == "__main__":
    number = sys.argv[1] if len(sys.argv) > 1 else "14155551234"
    result = check_whatsapp_number(number, os.environ["WAVALIDATOR_API_KEY"])

    if result["status"] == "valid":
        print(f"{result['phone_number']}: has WhatsApp")
    elif result["status"] == "invalid":
        print(f"{result['phone_number']}: no WhatsApp")
    else:
        print(f"{result['phone_number']}: check failed, try again later")
    print(f"Credits remaining: {result['credits_remaining']}")
