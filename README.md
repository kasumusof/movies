# movies

 [link go file](gomovyies.herokuapp.com/documentation/)

 I was not able to finish the features
 However, I did justice to a few.

### Task

Create a small set of rest API endpoints using Golang to do the following

- List an array of movies containing the name, opening crawl and comment count (not completed)
- Add a new comment for a movie (implemented)
- List the comments for a movie (implemented)
- Get list of characters for a movie (implemented)
 ## General requirements

- The application should have basic documentation that lists available endpoints and methods along with their request and response signatures (done however, documentation is not completed)
- The exact api design in terms of total number of endpoints and http verbs is up to you
- Keep your application source code on a public Github repository (done)
- Deploy the API endpoints and provide a live demo URL of the API documentation. Heroku is a good option. (done)
- Bonus, but not mandatory, if you can dockerize the development environment. (tried and failed)

## Data requirements
 - The movie data should be fetched online from `**[https://swapi.dev](https://swapi.dev)**` (implemented partially)
- Movie names in the movie list endpoint should be sorted by release date from earliest to newest and each movie should be listed along with opening crawls and count of comments. (not implemented)
- Data fetched from `**[https://swapi.dev](https://swapi.dev)`** should be cached with Redis and then accessed from the cache for subsequent requests. (implemented partially)
- Comments should be stored in a Postgres database. (implemented)
- Error responses should be returned in case of errors. (barely implemented)

## Character list requirements

- Endpoint should accept sort parameters to sort by one of name, gender or height in ascending or descending order.
- Endpoint should also accept a filter parameter to filter by gender. (not implemented)
- The response should also return metadata that contains the total number of characters that match the criteria along with the total height of the characters that match the criteria. (not implemented)
- The total height should be provided both in cm and in feet/inches. For instance, 170cm makes 5ft and 6.93 inches.(not implemented)

## Comment requirements

- Comment list should be retrieved in reverse chronological order. (implemented)
- Comments should be retrieved along with the public IP address of the commenter and UTC date&time they were stored (implemented)
- Comment length should be limited to 500 characters (implemented)
 
