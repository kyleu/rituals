# rituals.dev

[rituals.dev](https://rituals.dev) helps you manage your team's meetings. 
Join estimation sessions to help plan your work, log daily standups to keep up with your progress, and participate in retrospectives to help improve.
Create sprints to collect your team's efforts in one place, and join any number of teams so your work remains private. 

https://github.com/KyleU/rituals.dev

## Documentation

- [features.md](doc/features.md)
- [installing.md](doc/installing.md)
- [scripts.md](doc/scripts.md)

## Building

For CI servers and go-only changes, simply `make build`. For full stack development, you'll need some tools installed

- Run `bin/bootstrap.sh` to install required Go utilities
- For macOS, you can install all dependencies with `brew install md5sha1sum closure-compiler typescript sass/sass/sass`
- After editing stylesheets, use `bin/build-css.sh`; you'll need `sass` installed
- For TypeScript changes, use `bin/build-client.sh`; you'll need `tsc` and `closure-compiler` installed
- For a developer environment, run `bin/workspace.sh`, which will watch all files and hot-reload
