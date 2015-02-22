# terminal-watcher
Monitor your shell scripts scripts from your phone. Terminal watcher helps to monitor long running shell commands like compilation or batch jobs. It works by sending you a push notification on your phone when the command completes.

## Installation
See the relevant [release](https://github.com/rcoh/terminal-watcher/releases). Currently:
```
curl -L "https://raw.githubusercontent.com/rcoh/terminal-watcher/master/installer.sh" | bash
```
Then install the [Android App](https://play.google.com/store/apps/details?id=me.rcoh.terminalwatcher&hl=en)
The Android app source is at: https://github.com/rcoh/tw-android

## Usage
Once installed, run a command as you normally would be preface it with `tw`. The start, end, and exit status code will be sent to your phone. Eg:
```
bash $ tw sleep 10
```
In 10 seconds, your phone will buzz with a notification!
