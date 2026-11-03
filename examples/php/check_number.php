<?php
// Check if a phone number has WhatsApp with the WA Validator API.
// Usage:
//   export WAVALIDATOR_API_KEY=your_key
//   php check_number.php 14155551234

function checkWhatsAppNumber(string $phoneNumber, string $apiKey): array
{
    $ch = curl_init('https://wavalidator.com/api/v1/validate/');
    curl_setopt_array($ch, [
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_POST => true,
        CURLOPT_TIMEOUT => 30,
        CURLOPT_HTTPHEADER => [
            'Authorization: Bearer ' . $apiKey,
            'Content-Type: application/json',
        ],
        CURLOPT_POSTFIELDS => json_encode(['phone_number' => $phoneNumber]),
    ]);
    $body = curl_exec($ch);
    $status = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    curl_close($ch);

    $result = json_decode($body, true);
    if ($status !== 200) {
        throw new RuntimeException("HTTP $status: " . ($result['error'] ?? $body));
    }
    return $result;
}

$number = $argv[1] ?? '14155551234';
$result = checkWhatsAppNumber($number, getenv('WAVALIDATOR_API_KEY'));

$labels = ['valid' => 'has WhatsApp', 'invalid' => 'no WhatsApp'];
echo $result['phone_number'] . ': ' . ($labels[$result['status']] ?? 'check failed, try again later') . PHP_EOL;
echo 'Credits remaining: ' . $result['credits_remaining'] . PHP_EOL;
