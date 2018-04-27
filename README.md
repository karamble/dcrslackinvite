# dcrslackinvite

Web portal for slack invitations.

## Installation

Posting the email addresses that request an invite into the slack channel we need to install the slack-poster dependency:

```bash
go get -u github.com/m0t0k1ch1/go-slack-poster/cmd/slackpost
```
Build the package:
```bash
go build
```
and then run the file:
```bash
./dcrslackinvite
```

This will launch the invitation web portal for you. The web server is listening on port 8000 on localhost.

### Configuration

Before running the dcrslackinvite portal you need to edit the included sample-config.json file and rename it to config.json

In order to get the Slack API Token visit: https://api.slack.com/custom-integrations/legacy-tokens
The token will look something like this `xoxo-2100000415-0000000000-0000000000-ab1ab1`.

## Contact

If you have any further questions you can find us at:

- irc.freenode.net (channel #decred)
- [webchat](https://webchat.freenode.net/?channels=decred)
- forum.decred.org
- decred.slack.com

## Issue Tracker

The
[integrated github issue tracker](https://github.com/karamble/dcrslackinvite/issues)
is used for this project.

## License

dcrslackinvite is licensed under the [copyfree](http://copyfree.org) ISC License.
