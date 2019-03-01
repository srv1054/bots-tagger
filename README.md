# bots-tagger

Tagger will tag things people say in a slack channel that contain specific words with a pre-determined Emoji.

Such as hearing the word "Synergy" in a slack message tagger could stick a Business Cat emoji on that message.

You need a slack webhook and a slack bot token to make it run.

Here's where: https://api.slack.com/apps?new_app=1  
You can configure the Webhook and BOT features in the "Add Features and Functionality" Section.

To make this bot go you will need two things once this is all complete:
- Your Incoming Webhook URL. It will look kinda like this: `https://hooks.slack.com/services/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QkaNlP` )
- Your Slack App BOT Token (Look in the OAUTH section).  It will look kinda like this (it must start with xoxb and not anything else): 
`xoxb-76d11452182-7611398727-217937067376-46e6edasdff9054daf3000baf1aa36b8836a`

You will need to pass these to your bot on the command line when you launch it:

`tagger -slacktoken YOURTOKEN -slackhook WEBURL`  (no quotes needed)

## Notes on running
At the moment tag bot just holds the console open, and will dump messages to it if DEBUG is on as well as to slack logging channel specified in the config.json.  Either run this inside a "screen" session or dedicated window, or run it as a service so that if it crashes it will restart itself.    If the Slack RTM disconnects you (it can happen) you will have to relaunch the bot.  I'm working on a more graceful recovery around that, so stay tuned.

## tags.json - Required Data File
This file tells tagger what to do.   See the sample in the source code, feel free to steal it and edit it.  You must properly format this json file or tagger will puke on you when you run it.  Each section contains the emoji to spray tag something with and the words that trigger it.   Be careful what words you use, or you will have tagger hitting everything all the time....quite annoying really.  

Taggers searching is wildcarded so that the slack statement "contains" the words you list, anywhere in the statement.

As of release 1.3 you can now update the `tags.json` file while tagger is running and then from slack issue a command to have him re-read/load the tags to get your updates.  Command:  `@tagger reload tags`

## Creating the Config
You will need to create or pull from source and edit tagger.json using standard JSON file format fill in the few required items.
- LogChannel - this should be set to a dedicated slack channel so the bot can dump logs about what its doing and errors it runs into directly into Slack.  This channel can be public or private and the Bot should be invited to it.  Put the name in here including the `#` symbol used by slack.   If this is wrong or left blank slack will default these messages to whatever channel the Tokens are tied to.  This can be noisey so i suggest creating a channel for it or using a non popular channel!
- Debug - setting this to true will cause the BOT to dump errors and messages to the console it is running in.  In some cases this can provide additional info if needed for troubleshooting.  Normally this should just be set to False
- Currently ALL other values in the config.json are being ignored.

## TODO
- More graceful RTM disconnects
- ~Ability to send slack command to reload `tags.json` without having to restart the bot~
- Clean up config.json for realz

