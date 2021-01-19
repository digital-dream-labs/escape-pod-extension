# battery

Ask Vector how his battery is doing!

## Usage

This is meant to interface with other packages in this repository.  Feel free to use it elsewhere if you find it useful.

## Environment variables

|Variable| Description |
|--|--|
| ESCAPE_POD_ROBOT_TARGET | The IP address of your vector |
| ESCAPE_POD_ROBOT_TOKEN | The token for your robot |

## Intent configuration

You'll need to write a separate intent in the escape pod.  Here's an example:

```json
{
  "name": "battery",
  "description": "",
  "intent": "intent_custom_hack_battery",
  "utterance_list": "how is your battery",
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