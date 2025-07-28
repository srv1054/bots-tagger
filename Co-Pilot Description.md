# bots-tagger Application Summary

## Overview

**bots-tagger** is a Slack bot written in Go that automatically tags Slack messages with emoji reactions based on the presence of specific keywords. The bot uses Slack's RTM (Real Time Messaging) API to monitor messages in real time and applies emoji reactions according to rules defined in a configurable JSON file. The project is now obsolete due to Slack's RTM deprecation; see [tagger-v2](https://github.com/srv1054/tagger-v2) for the replacement.

---

## Key Features

- **Keyword-Based Emoji Tagging:**  
  Monitors Slack messages and adds emoji reactions when messages contain predefined keywords.

- **Configurable Tag Rules:**  
  Uses a `tags.json` file to define which keywords trigger which emoji reactions. Multiple tags and keywords are supported.

- **Slack Integration:**  
  Requires a Slack webhook and bot token for operation. Can send logs and notifications to a specified Slack channel.

- **Live Tag Reloading:**  
  Supports reloading the `tags.json` file at runtime via a Slack command, allowing updates without restarting the bot.

- **Direct Message Commands:**  
  Responds to DM commands for help, listing all tags, and showing keywords for specific tags.

- **Logging:**  
  Logs actions and errors to a dedicated Slack channel and optionally to the console if debug mode is enabled.

- **Custom Emoji Support:**  
  Includes sample emoji images for use as Slack emoji or bot avatars.

---

## Typical Workflow

1. **Startup:**  
   Loads configuration from `config.json` and tag rules from `tags.json`. Announces startup in the log channel.

2. **Message Monitoring:**  
   Listens to all Slack messages in channels where the bot is present.

3. **Tagging:**  
   When a message contains a keyword from `tags.json`, the bot reacts with the corresponding emoji.

4. **Slack Commands:**  
   Users can interact with the bot using commands like `show all tags`, `reload tags`, and `show keywords for <tag>`.

5. **Logging:**  
   All actions and errors are logged to the specified Slack channel.

---

## Configuration Highlights

- **config.json:**  
  - `logchannel`: Slack channel for logs (e.g., `#scott-channel-of-test`)
  - `debug`: Enables verbose console logging

- **tags.json:**  
  - Defines emoji tags and associated keywords (see [tags.json](x:/src/bots-tagger/tags.json) for examples)

- **Command Line Parameters:**  
  - `-slackhook`: Slack webhook URL (required)
  - `-slacktoken`: Slack bot token (required)
  - `-conf`: Path to `config.json` (optional)
  - `-json`: Path to `tags.json` (optional)
  - `-v`: Print version and exit

---

## Slack Commands

- `@tagger show all tags`  
  DM with all tags and their keywords

- `@tagger reload tags`  
  Reloads `tags.json` without restarting

- `@tagger show keywords for <tag name>`  
  DM with keywords for a specific tag

- `@tagger help`  
  DM with available commands and usage

---

## References

- [README.md](x:/src/bots-tagger/README.md)
- [README-original.md](x:/src/bots-tagger/README-original.md)
- [tags.json](x:/src/bots-tagger/tags.json)
- [config.json](x:/src/bots-tagger/config.json)
- [tagbot.go](x:/src/bots-tagger/tagbot.go)
- [tagger.go](x:/src/bots-tagger/tagger/tagger.go)
- [slack.go](x:/src/bots-tagger/tagger/slack.go)
- [config.go](x:/src/bots-tagger/tagger/config.go)

---

**Note:**  
This bot is deprecated due to Slack RTM API changes. For the latest version, see [tagger-v2](https://github.com/srv1054/tagger-v2).