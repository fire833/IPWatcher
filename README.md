# IPWatcher

For users who want to know where they fit into the internet.

## What is IPWatcher
IPWatcher is a daemon to track your public IP address and report changes to a backend notification API. It supports multiple notification methods, including notification through MS Teams, Discord, and Slack webhooks, pushover notifications, or even basic OS notifications. 

_Note: currently, only webhooks and pushover are the only actively maintained notification backends that are actually enabled/fully supported by the daemon to push local change messages to remote hosts. If you want to add new notification backends, please look at the [contributing guideline](CONTRIBUTING.md)._

## Installation from source

Simple:

```
git clone https://github.com/fire833/IPWatcher

cd ipwatcher

make

sudo make install
```