# apiRequestScheduler
Trigger API requests based on scheduling with ICS (Google Calendar, etc)

# What does it do?
1. Pulls a calendar ICS (Google Calender etc)
2. Check the current event summary (Subject of the event)
3. Matches this against defined values
4. Performs a POST request to an API with appropriate payload

## My usecase
To schedule the power of a server which I have in my home. I can create an event in Google Calendar, which triggers the server to power on or off at certain times in the day. The powering ON allows me to not have to wait for it to boot. The powering OFF allows it to be a graceful shutdown. You can buy smart plugs, which pull the power on the server, but this is not graceful and has potential risks to the OS, data and hardware of the server. 

The intension to be be flexible, by allowing user configurable keywords to trigger user configurable API requests, so if I wanted to do something more in the future, I can just update the config file and not the code. It is possible to point this at any API that supports basic authentication and trigger requests on a schedule. 

I personally run this on a RaspberryPi that is on the same LAN as the server, so do not have to worry about exposing the server API to the internet for a service to trigger the API. This also means I have a self signed certificate which I have enabled the option to ignore warnings for. This runs as a cron job every X minutes and outputs StdOut to a log file.

## How to use this?
1. Pull the code and compile your own binary for your platform (I have not provided binaries, though it may be an option in the future)
2. Copy the sampleConfig.json to config.json
3. Update the config.json with your API detail (URL, user, pass, payload, etc)
4. Pass the config file name as an CLI argument

`$ ./apiRequestScheduler config.json`
`$ ./apiRequestScheduler /home/pi/apiRequestScheduler/config.json`

If you want to use this in cron, this is what I'm using on Raspbian on the RaspberryPi

`$ crontab -e`

Run every 1 minute
```
# m h  dom mon dow   command
*/1 * * * * /home/pi/apiRequestScheduler/apiRequestScheduler.arm.bin /home/pi/apiRequestScheduler/config.json
```

Run every 1 minute and save StdOut to a logfile
```
# m h  dom mon dow   command
*/1 * * * * /home/pi/apiRequestScheduler/apiRequestScheduler.arm.bin /home/pi/apiRequestScheduler/config.json 1>> /home/pi/apiRequestScheduler/apiRequestScheduler.log
```