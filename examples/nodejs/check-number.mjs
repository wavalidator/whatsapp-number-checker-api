// Check if a phone number has WhatsApp with the WA Validator API.
// Node.js 18+ (built-in fetch). Usage:
//   export WAVALIDATOR_API_KEY=your_key
//   node check-number.mjs 14155551234

const API_URL = "https://wavalidator.com/api/v1/validate/";

export async function checkWhatsAppNumber(phoneNumber, apiKey, retries = 3) {
  for (let attempt = 0; attempt < retries; attempt++) {
    const res = await fetch(API_URL, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${apiKey}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify({ phone_number: phoneNumber }),
    });

    if (res.status === 429) {
      const body = await res.json();
      await new Promise((r) => setTimeout(r, (body.retry_after ?? 2 ** attempt) * 1000));
      continue;
    }
    if (!res.ok) throw new Error(`HTTP ${res.status}: ${await res.text()}`);
    return res.json();
  }
  throw new Error("Rate limit still reached after retries");
}

const number = process.argv[2] ?? "14155551234";
const result = await checkWhatsAppNumber(number, process.env.WAVALIDATOR_API_KEY);

const label = { valid: "has WhatsApp", invalid: "no WhatsApp" }[result.status] ?? "check failed, try again later";
console.log(`${result.phone_number}: ${label}`);
console.log(`Credits remaining: ${result.credits_remaining}`);
