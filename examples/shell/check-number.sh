#!/usr/bin/env sh
# Check if a phone number has WhatsApp with the WA Validator API.
# Usage:
#   export WAVALIDATOR_API_KEY=your_key
#   ./check-number.sh 14155551234

NUMBER="${1:-14155551234}"

curl -sS -X POST "https://wavalidator.com/api/v1/validate/" \
  -H "Authorization: Bearer ${WAVALIDATOR_API_KEY}" \
  -H "Content-Type: application/json" \
  -d "{\"phone_number\": \"${NUMBER}\"}"
echo
