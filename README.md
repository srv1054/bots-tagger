# bots-tagger

Tagger will tag things people say in a slack channel that contain specific words with a pre-determined Emoji.

Such as hearing the word "Synergy" in a slack message tagger will stick a Business Cat emoji on that message.

You need a slack webhook and a slack bot token to make it run.

Here's where: https://api.slack.com/apps?new_app=1  
You can configure the Webhook and BOT features in the "Add Features and Functionality" Section.

To make this bot go you will need two things once this is all complete:
- Your Incoming Webhook URL. It will look kinda like this: `https://hooks.slack.com/services/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QkaNlP` )
- Your Slack App BOT Token (Look in the OAUTH section).  It will look kinda like this: 
`xoxb-76d11452182-7611398727-217937067376-46e6edasdff9054daf3000baf1aa36b8836a`

You will need to pass these to your bot on the command line when you launch it:

`tagger -slacktoken YOURTOKEN -slackhook WEBURL`  (no quotes needed)

## Creating the Config
You will need to edit tagger.json using standard JSON file format fill in the few required items.
- LogChannel - this should be set to a dedicated slack channel so the bot can dump logs about what its doing and errors it runs into directly into Slack.  This channel can be public or private and the Bot should be invited to it.  Put the name in here including the `#` symbol used by slack.
- Debug - setting this to true will cause the BOT to dump errors and messages to the console it is running in.  In some cases this can provide additional info if needed for troubleshooting.  Normally this should just be set to False


