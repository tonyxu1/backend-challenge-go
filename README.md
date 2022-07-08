# RareCircles Coding Challenge

## Requirements

Design an API endpoint that provides token information based on token title.

- the endpoint is exposed at `/tokens`
- the partial (or complete) token title is passed as a query string parameter `q`
- the endpoint returns a JSON response with an array suggested matches
  - each suggestion has a name
  - each suggestion has a symbol
  - each suggestion has a address
  - each suggestion has a decimals
  - each suggestion has a total supply
- at-least 2 go-test needs to be implemented
- concurrency should be applied where appropriate  
- feel free to add more features if you like!

#### Notes
- you have a list of tokens accounts defined inthe `addresses.jsonl` file that should be used as a seed to your application
- we have included a `eth` lirbary that should help with any decoding and rpc calls

#### Sample responses

These responses are meant to provide guidance. The exact values can vary based on the data source and scoring algorithm.

**Near match**

    GET /tokens?q=rare

```json
{
  "tokesn": [
    {
      "name": "RareCircles",
      "symbol": "RCI",
      "address": "0xdbf1344a0ff21bc098eb9ad4eef7de0f9722c02b",
      "decimals": 18,
      "totalSupply": 1000000000,
    },
    {
      "name": "Rareible",
      "symbol": "RRI",
      "address": "e9c8934ebd00bf73b0e961d1ad0794fb22837206",
      "decimals": 9,
      "totalSupply": 100,
    },
  ]
}
```

**No match**

    GET /tokens?q=ThereIsNoTokenHere

```json
{
  "tokens": []
}
```


### Non-functional

- All code should be written in Goland, Typescript or PHP.
- Mitigations to handle high levels of traffic should be implemented.
- Challenge is submitted as pull request against this repo ([fork it](https://help.github.com/articles/fork-a-repo/) and [create a pull request](https://help.github.com/articles/creating-a-pull-request-from-a-fork/)).
- Documentation and maintainability is a plus.

## Dataset

You can find the necessary dataset along with its description and documentation in the [`data`](data/) directory.

## Evaluation

We will use the following criteria to evaluate your solution:

- Capacity to follow instructions
- Developer Experience (how easy it is to run your solution locally, how clear your documentation is, etc)
- Solution correctness
- Performance
- Tests (quality and coverage)
- Attention to detail
- Ability to make sensible assumptions

It is ok to ask us questions!

We know that the time for this project is limited and it is hard to create a "perfect" solution, so we will consider that along with your experience when evaluating the submission.

## Getting Started

### Prerequisites

You are going to need:

- `Git`
- `go`

### Starting the application

To start a local server run:

```
go run ./cmd/challenger
```

## Solution
### Design
1. An HTTP Server to serve `GET /tokens?q=<query parameters>` by using [echo](https://echo.labstack.com/)
2. A cron job to periodically get updated NFT information from ETH mainnet, push the contract metadata to Golang sync map.
3. Get proper contract metadata from webserver memory once client requests come in.
4. Break contract address list into smaller chunks and send them to a separate goroutine to get contract metadata.
5. Extend the query parameter to support comma separated values.
6. Extend the query parameter to serch metadata by contract address.

### Benefits of this approach:
1. Reduce the denpendency on external services. (alchemy services, etc)
2. No extra work for horizontal scaling.
3. Best performance.
4. No need to worry about the underlying technology, even alchemy is down, we can still serve the requests, with the existing data.
5. No extra work need to be done to run the server behind a load balancer.

### Running the server
To run the application from local machine:  

```
    git clone <the repo>
    PORT=2304 ADDRESS_FILE_DIR=./data go run ./cmd/challenge
```

### Notes:
1. The `PORT` and `ADDRESS_FILE_DIR` are optional, will use default values if not provided.
2. The `ADDRESS_FILE_DIR` is the directory that contains the `addresses.jsonl` file.
3. The `addresses.jsonl` file is a list of addresses that we will use to get the contract metadata.
4. Decimal places are set to be 0 for all tokens, which is correct value for NFTs
5. Some of hardcoded values should be configurable.
6. New query parameter should be added to return all tokens e.g. `/tokens?q=all_tokens`
7. Since the metadata is loaded asyncronously, we may not return any values for a query in first few seconds.