# generic response

Easily identify a generic response!

## Usage

This is meant to interface with other packages in this repository.  Feel free to use it elsewhere if you find it useful.

## Environment variables

|Variable| Description |
|--|--|
| ESCAPE_POD_ROBOT_TARGET | The IP address of your vector |
| ESCAPE_POD_ROBOT_TOKEN | The token for your robot |

## Initialization

When initializing this, you'll want to set the actual response you want.  Here's an example!

```go
	home, err := genericresponse.New(
		genericresponse.WithViper(),
		genericresponse.WithResponse(
			"I'm talking with the escape pod.  There's loads more to do!",
		),
	)
	if err != nil {
		return nil, err
	}
```

## Intent configuration

You'll need to write a separate intent in the escape pod.  Here's an example:

```json
{
  "name": "escapepod",
  "description": "",
  "intent": "intent_custom_hack_escapepod",
  "utterance_list": "where are you pointed",
  "extended_options": {
    "block_list": [],
    "parser": null,
    "external_parser": true
  },
  "response_parameters": {
    "final_intent": "intent_play_cantdo"
  }
}
```

The response_parameters.final_intent field will not send if this transaction succeeds, as the SDK temporarily takes over all cloud response interactions.  Anything you'll want to do will need to be handled in your custom module