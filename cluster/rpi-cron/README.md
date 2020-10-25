# Crontab Demo

## cron-app

Simple GO application that creates a random file at `/home/pi/pixel`, with random RGB colors.
After 10 second the file will be deleted, and the program returns.

**cron-app.yaml**

Contains a CronJob running the cron-app with the set interval.
