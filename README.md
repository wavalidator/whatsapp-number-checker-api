# WhatsApp Number Checker API — check if a phone number has WhatsApp

[![API docs](https://img.shields.io/badge/docs-wavalidator.com-25D366)](https://wavalidator.com/api-docs/)
[![OpenAPI 3.0](https://img.shields.io/badge/OpenAPI-3.0-6BA539)](openapi.yaml)
[![Status](https://img.shields.io/badge/status-live-brightgreen)](https://wavalidator.com/status/)
[![License: MIT](https://img.shields.io/badge/license-MIT-blue)](LICENSE)

Official code examples for the **[WA Validator](https://wavalidator.com) WhatsApp number validation API**. Send one phone number, get back whether it is registered on WhatsApp — synchronously, in the same HTTP response.

- **Endpoint:** `POST https://wavalidator.com/api/v1/validate/`
- **Auth:** `Authorization: Bearer YOUR_API_KEY`
- **Cost:** 1 credit per number checked (server-side errors are not charged)
- **Response:** `valid` = has WhatsApp · `invalid` = no WhatsApp

> Need to check a list? Use the [WhatsApp Bulk Checker API](https://github.com/wavalidator/whatsapp-bulk-checker-api) — up to 100 numbers per request.

## Contents

- [Why use a WhatsApp number checker?](#why-use-a-whatsapp-number-checker)
- [Get an API key](#get-an-api-key)
- [Request](#request)
- [Response](#response)
- [Code examples](#code-examples)
- [Phone number format](#phone-number-format)
- [Errors and rate limits](#errors-and-rate-limits)
- [FAQ](#faq)

## Why use a WhatsApp number checker?

- **Validate leads at signup** — reject or flag phone numbers that have no WhatsApp account before they reach your CRM.
- **Clean WhatsApp marketing lists** — stop paying WhatsApp Business API conversation fees for numbers that can never receive a message.
- **Verify COD and e-commerce orders** — confirm the customer can be reached on WhatsApp before you ship.
- **Reduce fake signups** — a number with no WhatsApp account is a useful fraud and spam signal.
- **Route customers to the right channel** — WhatsApp if `valid`, SMS or email if `invalid`.

## Get an API key

1. [Create a free account](https://wavalidator.com/signup/) on wavalidator.com.
2. Open the [dashboard](https://wavalidator.com/dashboard/) and click **Generate API key**.
3. Keep the key on your server. Never put it in browser JavaScript or a mobile app.

## Request

```http
POST /api/v1/validate/ HTTP/1.1
Host: wavalidator.com
Authorization: Bearer YOUR_API_KEY
Content-Type: application/json

{ "phone_number": "14155551234" }
```

| Field | Type | Required | Description |
|---|---|---|---|
| `phone_number` | string | yes | Phone number in international format with country code. `+`, spaces and `-` are stripped automatically. |

## Response

```json
{
  "phone_number": "14155551234",
  "status": "valid",
  "credits_remaining": 99
}
```

| Field | Type | Description |
|---|---|---|
| `phone_number` | string | The number after normalization (digits only). |
| `status` | string | `valid` — registered on WhatsApp. `invalid` — not registered, or not a valid phone number. `error` — the check could not be completed; retry later, no credit charged. |
| `credits_remaining` | integer | Your credit balance after this request. |
| `cached` | boolean | Present and `true` when the result came from a check of the same number in the last few minutes. |

## Code examples

| Language | File | Run |
|---|---|---|
| cURL / Shell | [`examples/shell/check-number.sh`](examples/shell/check-number.sh) | `WAVALIDATOR_API_KEY=... ./check-number.sh 14155551234` |
| Python | [`examples/python/check_number.py`](examples/python/check_number.py) | `python check_number.py 14155551234` |
| Node.js | [`examples/nodejs/check-number.mjs`](examples/nodejs/check-number.mjs) | `node check-number.mjs 14155551234` |
| PHP | [`examples/php/check_number.php`](examples/php/check_number.php) | `php check_number.php 14155551234` |
| Go | [`examples/go/main.go`](examples/go/main.go) | `go run main.go 14155551234` |
| Java | [`examples/java/CheckNumber.java`](examples/java/CheckNumber.java) | `java CheckNumber.java 14155551234` |
| C# / .NET | [`examples/csharp/CheckNumber.cs`](examples/csharp/CheckNumber.cs) | `dotnet run 14155551234` |

Every example reads the key from the `WAVALIDATOR_API_KEY` environment variable.

### Python

```python
import os, requests

r = requests.post(
    "https://wavalidator.com/api/v1/validate/",
    headers={"Authorization": f"Bearer {os.environ['WAVALIDATOR_API_KEY']}"},
    json={"phone_number": "14155551234"},
    timeout=30,
)
print(r.json()["status"])  # "valid" or "invalid"
```

### Node.js

```js
const res = await fetch("https://wavalidator.com/api/v1/validate/", {
  method: "POST",
  headers: {
    Authorization: `Bearer ${process.env.WAVALIDATOR_API_KEY}`,
    "Content-Type": "application/json",
  },
  body: JSON.stringify({ phone_number: "14155551234" }),
});
console.log((await res.json()).status);
```

### PHP

```php
$ch = curl_init('https://wavalidator.com/api/v1/validate/');
curl_setopt_array($ch, [
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_POST => true,
    CURLOPT_HTTPHEADER => [
        'Authorization: Bearer ' . getenv('WAVALIDATOR_API_KEY'),
        'Content-Type: application/json',
    ],
    CURLOPT_POSTFIELDS => json_encode(['phone_number' => '14155551234']),
]);
echo json_decode(curl_exec($ch), true)['status'];
```

## Phone number format

Send the full international number: **country code + national number**, no leading `00`.

| Country | Local | Send |
|---|---|---|
| United States | (415) 555-1234 | `14155551234` |
| United Kingdom | 07911 123456 | `447911123456` |
| Morocco | 06 12 34 56 78 | `212612345678` |
| France | 06 12 34 56 78 | `33612345678` |
| Saudi Arabia | 050 123 4567 | `966501234567` |
| India | 98765 43210 | `919876543210` |
| Brazil | (11) 91234-5678 | `5511912345678` |

After stripping `+`, spaces and `-`, the number must be **7 to 15 digits**. Anything else returns `status: "invalid"` without a live check (and still costs 1 credit), so drop the national trunk `0` before you send.

## Errors and rate limits

| HTTP | Meaning | What to do |
|---|---|---|
| `400` | Invalid JSON or `phone_number` missing | Fix the request body. |
| `401` | Missing or invalid API key | Check the `Authorization` header. |
| `402` | Insufficient credits | [Top up](https://wavalidator.com/pricing/). |
| `405` | Method not allowed | Use `POST`. |
| `429` | Rate limit reached — body has `"status": "limit"` | Wait `retry_after` seconds, then retry. |
| `5xx` | Server error | Retry with exponential backoff. |

Each response carries `X-RateLimit-Limit`, `X-RateLimit-Remaining` and `X-RateLimit-Reset` headers. The default is **20 requests per minute** per API key, plus a daily validation quota that depends on your plan. For higher throughput, use the [bulk checker](https://github.com/wavalidator/whatsapp-bulk-checker-api) — 100 numbers per request.

## FAQ

**How do I check if a phone number has WhatsApp without saving it to my contacts?**
Send it to `POST /api/v1/validate/`. The API tells you whether the number is registered on WhatsApp — no phone, no contacts list, no WhatsApp Web session needed.

**Does the person know their number was checked?**
No message is sent to the number. The check only reports whether a WhatsApp account exists.

**Can I use this API from the browser?**
No. Call it from your backend so your API key stays secret.

**Is there an OpenAPI / Swagger spec?**
Yes — [`openapi.yaml`](openapi.yaml). Import it into Postman, Insomnia or any OpenAPI code generator.

**Can I use it from ChatGPT, Claude or Cursor?**
Yes, through the [WA Validator MCP server](https://github.com/wavalidator/whatsapp-checker-mcp-server).

## Related

- [WhatsApp Bulk Checker API](https://github.com/wavalidator/whatsapp-bulk-checker-api) — 100 numbers per request
- [WhatsApp Checker MCP Server](https://github.com/wavalidator/whatsapp-checker-mcp-server) — for AI assistants
- [Guides and resources](https://github.com/wavalidator/wavalidator-resources)
- [Full API documentation](https://wavalidator.com/api-docs/)

---

Results are point-in-time and are not identity verification or consent to contact. Only check numbers you are authorized to process. WhatsApp is a trademark of Meta Platforms, Inc.; WA Validator is not affiliated with Meta.
