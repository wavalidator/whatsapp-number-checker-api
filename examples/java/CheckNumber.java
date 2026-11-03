// Check if a phone number has WhatsApp with the WA Validator API.
// Java 11+, no dependencies. Usage:
//   export WAVALIDATOR_API_KEY=your_key
//   java CheckNumber.java 14155551234

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.time.Duration;

public class CheckNumber {
    private static final String API_URL = "https://wavalidator.com/api/v1/validate/";

    public static String checkWhatsAppNumber(String phoneNumber, String apiKey) throws Exception {
        String json = "{\"phone_number\":\"" + phoneNumber.replaceAll("[^0-9+]", "") + "\"}";

        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(API_URL))
                .timeout(Duration.ofSeconds(30))
                .header("Authorization", "Bearer " + apiKey)
                .header("Content-Type", "application/json")
                .POST(HttpRequest.BodyPublishers.ofString(json))
                .build();

        HttpResponse<String> response = HttpClient.newHttpClient()
                .send(request, HttpResponse.BodyHandlers.ofString());

        if (response.statusCode() != 200) {
            throw new RuntimeException("HTTP " + response.statusCode() + ": " + response.body());
        }
        return response.body();
    }

    public static void main(String[] args) throws Exception {
        String number = args.length > 0 ? args[0] : "14155551234";
        String body = checkWhatsAppNumber(number, System.getenv("WAVALIDATOR_API_KEY"));

        // Use Jackson or Gson in production; a plain check keeps this example dependency-free.
        if (body.contains("\"status\": \"valid\"") || body.contains("\"status\":\"valid\"")) {
            System.out.println(number + ": has WhatsApp");
        } else if (body.contains("invalid")) {
            System.out.println(number + ": no WhatsApp");
        } else {
            System.out.println(number + ": check failed, try again later");
        }
        System.out.println(body);
    }
}
