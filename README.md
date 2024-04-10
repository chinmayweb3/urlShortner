# urlshortner

## API Link
You can access the URL Shortener API at the following link: https://urlshortner-njjo.onrender.com/api/v1/shortener or https://urlshortner-njjo.onrender.com/tester

## Endpoints
- `/:shortUrl`

- shortUrl (required): The short URL to be redirected.
- A 301 HTTP status code is returned, redirecting the client to the original URL.
    
- `/api/v1/shortener`
- This endpoint creates a new short URL for a given original URL.

Request Body:
- url (required): The original URL to be shortened.
  
Response:
- A JSON object containing the short URL is returned.

## Rate Limiting
To prevent abuse and ensure fair usage, the API enforces a rate limit of 10 URLs per day per IP address. Exceeding this limit will result in a 429 Too Many Requests HTTP status code.

Creating a Short URL
``` bash
curl -X POST https://urlshortner-njjo.onrender.com/shortener \
-H "Content-Type: application/json" \
-d '{"url": "https://www.example.com"}'
```

Redirecting to the Original URL
```bash
curl -X GET https://urlshortner-njjo.onrender.com/shortUrl
```
