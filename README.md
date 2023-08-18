## go-geo-resolver

Just started learning Go by creating a very simple georesolver API.

- Start the server

```js
go run .
```

- Resolve a city

```js
http://localhost:4000/geocode/?city=Vienna
```

If the location is found, you will get back the geocoordinates.
Works with most major cities in the dach-region.
