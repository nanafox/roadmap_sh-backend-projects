# GitHub User Activity - A roadmap.sh Backend Project

This project fetches user activity and displays it in the terminal using
GitHub API.

Find the [GitHub User Activity](https://roadmap.sh/projects/github-user-activity) on [roadmap.sh](https://roadmap.sh)

## Usage

```bash
github-activity <username> [number of events: Default is 10],
```

## Description

`github-activity` is a simple CLI tool to fetch the latest activities a user
is performing with their GitHub account. It returns the results in a simple
listing format.

For efficiency, responses are cached for later in the `/tmp/github-activity-cache/`
directory. If you need to get fresh updates for the same user, feel free to
delete the cached files for that user or all users if you prefer.

The cache files are saved in this format, `<username>_<page_size>`. So as an
example running `github-activity nanafox 10` will return the first 10 events
of the `nanafox` user. The resultant cache file will be `nanafox_10.json`

## Dependencies

This project uses an [HTTP Client](https://github.com/nanafox/simple-http-client)
I developed. It simply makes the request much more straightforward and it was
just what I needed for this project.

You can [visit the repo](https://github.com/nanafox/simple-http-client) to learn
more about it.
