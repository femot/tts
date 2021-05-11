# TTS
This is a simple TTS service meant to be used together with a Twitch bot (e.g. Streamlabs Chatbot) to play back TTS.

# Disclaimer
This project was written for my own personal use.

- ⚠ No support
- ⚠ No guarantee it works/keeps working

# Prerequisites
- VLC installed under `C:\Program Files\VideoLAN\VLC\vlc.exe` (can be changed in code before running/compiling, if installed elsewhere)

# Preparing VLC
The TTS will play at the volume VLC last ran as, so you need to start VLC, set the desired volume and close it again.

You can run other VLC instances, while this TTS bot is running, but changing the volume, or muting other VLC windows might affect the TTS too.

# Streamlabs Chatbot Integration
You will need to add two commands:
- `!tts` to allow anyone to trigger a TTS message
- `!skiptts` in case you ever need to stop the spam :)

You can adjust the usage permissions on the command level within the Streamlabs Chatbot.

## !tts
In the Actions > Response field, add the following response:
```
$readapi(http://localhost:7777/?text=$msg)
```

## !skiptts
In the Actions > Response field, add the following response:
```
$readapi(http://localhost:7777/skip)
```