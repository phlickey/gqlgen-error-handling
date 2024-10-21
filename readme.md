# gqlgen Error Handling

This server errors. As a result, some queries will serve less data than we would expect.


Take this query for example:


```gql
query MyCatsAndDogsQuery {
  cats{
    id
    name
    meow
  }
  dogs {
    id
    name
    breed
  }
}
```

Executed against the server it will return the following result:

```json
{
  "errors": [
    {
      "message": "meow error. cat id: 1 cannot meow.",
      "path": [
        "cats",
        0,
        "meow"
      ]
    }
  ],
  "data": null
}
```

However, if we remove part of our query, we see there was much more data available than what we returned. 
```gql
query MyCatsAndDogsQuery {
  cats{
    id
    name
    # meow
  }
  dogs {
    id
    name
    breed
  }
}
```


returns:

```json
{
  "data": {
    "cats": [
      {
        "id": "1",
        "name": "Cat 1"
      },
      {
        "id": "2",
        "name": "Cat 2"
      }
    ],
    "dogs": [
      {
        "id": "1",
        "name": "Dog 1",
        "breed": "Breed 1"
      },
      {
        "id": "2",
        "name": "Dog 2",
        "breed": "Breed 2"
      }
    ]
  }
}
```


PRs to this repo show some potential ways of both surfacing errors and returning data that is not affected by unrelated erroring resolvers.
