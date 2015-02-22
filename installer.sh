#!/bin/bash
VERSION=0.3-beta
BASE_URL="https://github.com/rcoh/terminal-watcher/releases/download/$VERSION"
LINUX_BINARY_URL="$BASE_URL/client-linux"
OSX_BINARY_URL="$BASE_URL/client-osx"
TW_SOURCE="$BASE_URL/tw.source"
if [ "$(uname)" == "Darwin" ]; then
    # Do something under Mac OS X platform        
    echo "Detected Mac OS X (Darwin)"
  	BINARY_URL=$OSX_BINARY_URL

elif [ "$(uname -s)" == "Linux" ]; then
    echo "Detected Linux"
    BINARY_URL=$LINUX_BINARY_URL
else
	echo "Unsupported platform (Mac and Linux are supported only)"
    exit 1
fi

echo "Downloading binary from $BINARY_URL"
sudo curl -L "$BINARY_URL" -o /usr/local/bin/twclient
curl -L "$TW_SOURCE" -o ~/.tw-wrapper
sudo chmod +x /usr/local/bin/twclient
 
read -p "Please specify your shellrc: (eg. ~/.bashrc, ~/.zshrc: " RC_FILE < /dev/tty || {
	echo "Couldn't read input from /dev/tty"
	exit 1
}
# Perform ~ expansion
eval RC_FILE=$RC_FILE

echo "Removing any previous tw references ($RC_FILE backed up to $RC_FILE.backup)"
sed -i.backup '/source ~\/\.tw\-wrapper/d' $RC_FILE 
echo "source ~/.tw-wrapper" >> $RC_FILE
echo "Installation complete! Running test command"
source ~/.tw-wrapper
tw echo "Hello world!"
echo ""
echo ""
echo "Installation complete. Run: \"source $RC_FILE\" to complete installation"
