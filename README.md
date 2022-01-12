# escape-pod-extension

Want to add additional functionality to your [escape pod?](https://www.digitaldreamlabs.com/products/escape-pod)  You're in the right place!

Do you want your Vector robot to dance for every response on the third Tuesday of every month?  Do you want them to tell you the weather?  Write an app, turn on the "external_parser" option in your intent, and have at it!

## Package structure

|Directory| Description |
|--|--|
| proto | The proto definitions for inter-process communication, along with generated libraries in C++, C#, Go, Javascript, and Python |
| examples | Example applications, separated by language |

## How do I configure the escape pod to use these?

Note:  This will only function on beta 0.4 or above.  

At the bare minimum, you'll need to set the following variables in /etc/escape-pod.conf

|Variable| Description |
|--|--|
| ENABLE_EXTENSIONS | Set this to true to tell the escape pod to look for further configuration |
| ESCAPEPOD_EXTENDER_TARGET | The hostname and port the service you've written is running on |
| ESCAPEPOD_EXTENDER_DISABLE_TLS | This will allow clear-text communication with your serivce (but please use TLS!) |

## How do I get my token?
Obtaining the token is relatively easy in OSKR robots: When you are logged in via SSH into your robot, enter the command:  
`cat /data/data/com.anki.victor/persistent/token/token.jwt`  
DDL is working on a way to effectively and easily obtain the robot's token through alternative methods for production robots, as they cannot be accessed via SSH.

For a full list of additional variables, please see the [hugh grpc client library viper loader](https://github.com/digital-dream-labs/hugh/blob/main/grpc/client/viper.go)
