#!/usr/bin/osascript
tell application "iTerm2"
    tell current session of current tab of current window
        write text "cd ~/go/src/github.com/kyleu/rituals.dev"
        write text "rituals.dev"
        split vertically with default profile
    end tell
    tell second session of current tab of current window
        write text "cd ~/go/src/github.com/kyleu/rituals.dev"
        write text "scripts/build-client-watch.sh"
        split horizontally with default profile
    end tell
    tell third session of current tab of current window
        write text "cd ~/go/src/github.com/kyleu/rituals.dev"
        write text "scripts/build-css-watch.sh"
    end tell
end tell
