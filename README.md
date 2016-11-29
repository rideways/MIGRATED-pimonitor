# pimonitor

We use this at [rideways](http://www.rideways.com) to display our alerting status

It runs on a raspberry pi, and uses a [blink1](https://blink1.thingm.com/) which changes colour based
on the alerting status.

Our monitoring is based on [sensu](https://sensuapp.org/) and pimonitor uses it's API to find 
out our alerting status.

It uses [blink1-tiny-server](https://github.com/todbot/blink1/tree/master/commandline) to control the blink1

## Usage

```
pimonitor [options]

 -blink-url string
    	blink(1) server url (default "http://localhost:4567")
  -critical string
    	critical colour (default "ff0000")
  -ok string
    	ok colour (default "00ff00")
  -password string
    	sensu api password (default "password")
  -sensu-api-url string
    	sensu api url (default "http://localhost:8080")
  -username string
    	sensu api username (default "username")
  -warning string
    	warning colour (default "ffbe00")
```