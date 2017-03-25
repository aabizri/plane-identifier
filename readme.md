# plane-identifier is a web service accessible via a REST API for identifying a plane based upon measurements

Current version: v0.0

This project was developped for the [2017 Prague Hackaton](https://praguehackaton.com).

Built by [Alexandre A. Bizri](https://github.com/aabizri), [Marin Godechot](https://github.com/houguiram) & Camille Marchetti (Team K-1000).

Licensed under UNLICENSE

## Building the server
In the directory:
`go build`

## Launching it
`./plane-identifier`

Flags:
- `-p` allows you to select the port to listen to (ex: `-p 8080`)
- `-logpath` is the directory in which to log requests (ex: `-logpath /tmp/plane`)
- `-path` is the path to listen to (ex: `-path /task2/input`)

Calling without flags is equivalent to using `-p 8080 -logpath "/tmp/" -path "/"`

## Sample input
```json
{
  "measurements" : [
    {
      "type" : "A380",
      "noise-level": 103,
      "brake-distance": 2130,
      "vibrations": 0.81
    },{
      "type" : "A380",
      "noise-level": 101,
      "brake-distance": 2070,
      "vibrations": 0.88
    },{
      "type" : "737",
      "noise-level": 94,
      "brake-distance": 1730,
      "vibrations": 0.82
    },{
      "type" : "737",
      "noise-level": 96,
      "brake-distance": 1820,
      "vibrations": 0.79
    }
  ],
  "samples" : [
    {
      "id" : 1,
      "noise-level": 102,
      "brake-distance": 2105,
      "vibrations": 0.80
    },{
      "id" : 2,
      "noise-level": 97,
      "brake-distance": 1830,
      "vibrations": 0.80
    }
  ]
}
```

## Sample output
```json
{
    "result" : [
        {
            "id" : 1,
            "type" : "A380"
        },{
            "id" : 2,
            "type" : "737"
        }
    ]
}
```
