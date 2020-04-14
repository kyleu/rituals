# Scripts

There's scripts in the `./bin` directory to help you build, run, test, and publish rituals.dev

They're designed for macOS, but should work on Linux or Windows (via WSL).

- `build-css`: Uses `scss` to compile the stylesheets in `web/stylesheets`
- `build-css-watch`: Builds the css resources using `build-css`, then watches for changes in `stylesheets`
- `build-docker`: Makes a release build, builds a docker image, then exports and zips the output
- `check`: Runs code statistics, checks for outdated dependencies, then runs various linters
- `dev`: Watches the project directories, and runs the main application, restarting when changes are detected
- `doc`: Runs godoc for all projects, linking between projects and using custom logos and styling
- `format`: Runs `gofmt` on all projects
- `publish`: Runs the code formatter, checks all the projects, then builds binaries for Linux, macOS, and Windows
- `run-docker`: Runs the Docker image produced by `build-docker`, exposing an HTTP port
- `run-release`: Builds the project in release mode and runs it
