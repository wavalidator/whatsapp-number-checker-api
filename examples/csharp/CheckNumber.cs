// Check if a phone number has WhatsApp with the WA Validator API.
// .NET 6+. Usage:
//   dotnet new console -n CheckNumber && replace Program.cs with this file
//   export WAVALIDATOR_API_KEY=your_key
//   dotnet run 14155551234

using System.Net.Http.Headers;
using System.Net.Http.Json;
using System.Text.Json.Serialization;

var number = args.Length > 0 ? args[0] : "14155551234";
var apiKey = Environment.GetEnvironmentVariable("WAVALIDATOR_API_KEY");

using var http = new HttpClient { Timeout = TimeSpan.FromSeconds(30) };
http.DefaultRequestHeaders.Authorization = new AuthenticationHeaderValue("Bearer", apiKey);

var response = await http.PostAsJsonAsync(
    "https://wavalidator.com/api/v1/validate/",
    new { phone_number = number });

if (!response.IsSuccessStatusCode)
{
    Console.Error.WriteLine($"HTTP {(int)response.StatusCode}: {await response.Content.ReadAsStringAsync()}");
    return 1;
}

var result = await response.Content.ReadFromJsonAsync<ValidateResponse>();
var label = result!.Status switch
{
    "valid" => "has WhatsApp",
    "invalid" => "no WhatsApp",
    _ => "check failed, try again later",
};
Console.WriteLine($"{result.PhoneNumber}: {label}");
Console.WriteLine($"Credits remaining: {result.CreditsRemaining}");
return 0;

record ValidateResponse(
    [property: JsonPropertyName("phone_number")] string PhoneNumber,
    [property: JsonPropertyName("status")] string Status,
    [property: JsonPropertyName("credits_remaining")] int CreditsRemaining);
