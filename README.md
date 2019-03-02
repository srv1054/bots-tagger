# bots-tagger

Tagger will tag things people say in a slack channel that contain specific words with a pre-determined Emoji.

Such as hearing the word "Synergy" in a slack message tagger could stick a Business Cat emoji on that message.

This is only limited to your imagination and available slack emojis.  You can create all of these via the tags.json file, as many as you'd like.  Read on...

## To get it running
You need a slack webhook and a slack bot token to make it run.

Here's where you can build an app (the truly right way to implement): https://api.slack.com/apps?new_app=1  
You can configure the Webhook and BOT features in the "Add Features and Functionality" Section.

***However*** - for *free slacks* with limited integrations, or the quick and dirty way to do this you can simply add a standard slack webhook (emoji directory) and wrap the tagger avatar and name on it and use that URL.   Then in the Custom Integratons section add a stand-alone bot wrapping the tagger avatar (emoji directory) and tagger name around that.  This will get you the two required slack tokens you need above.  All of this can be found in the URL for your slack instance that https://&lt;MY SLACK INSTANCE&gt;.slack.com/apps/manage/custom-integrations

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

`tags.json` must be in the running directory **OR** if not, you must specify the path to it using the `-json` parameter.  The file **must** be called tags.json no matter what path you specify it in.

## Creating the Config
You will need to create or pull from source and edit tagger.json using standard JSON file format fill in the few required items.
- LogChannel - this should be set to a dedicated slack channel so the bot can dump logs about what its doing and errors it runs into directly into Slack.  This channel can be public or private and the Bot should be invited to it.  Put the name in here including the `#` symbol used by slack.   If this is wrong or left blank slack will default these messages to whatever channel the Tokens are tied to.  This can be noisey so i suggest creating a channel for it or using a non popular channel!
- Debug - setting this to true will cause the BOT to dump errors and messages to the console it is running in.  In some cases this can provide additional info if needed for troubleshooting.  Normally this should just be set to False

`config.json` must be in the running directory **OR** if not, you must specify the path to it using the `-conf` parameter.  The file **must** be called config.json no matter what path you specify it in.

## Images
Added emoji directory with both 512x512 and 128x128 tagger graphics for slack.   The 128 can be uploaded as an emoji as well as used for slackhook avatar.   The 512x512 is required if you setup an app bot, use this larger file for the avatar.

## Tagger Slack commands 
Here's a list of messages you can send tagger inside slack to get information:
- `@tagger show all tags` - tagger will Direct Message you inside Slack with all of tags in `tags.json` and their associated keywords.
- `@tagger reload tags` - tagger will re-read and load the `tags.json` file.  If you make edits to it this will make them effective without restarting tagger.  If you screw up the json formatting tagger will most likely crash after this command.
- `@tagger show keywords for <tag name>` - Will show all keywords tied to the specific tag (emoji) listed. For example: `@tagger show keywords for :businesscat:` (colons are optional)

## TODO
- More graceful RTM disconnects
- ~Ability to send slack command to reload `tags.json` without having to restart the bot~
- ~Clean up config.json for realz~
- ~Ability to request all words assigned to a tag~
- ~Ability to request all available tags~
- ~In Slack help request to display available commands~
- ~Ability to specify path to tags.json on command line~
- ~Ability to specify path to config.json on command line~
