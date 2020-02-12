# phabricator-tools
A collection of command line tools for manipulating stuff on Phabricator

Phabricator is nice in how powerful it is, but much like other general purpose tools like Jira, the UI leaves much to be desired for focused, simple workflows.

I wish to solve that for myself.


## First problem

I can't directly use the gonduit endpoint; every time I hit it, it returns HTML trying to redirect my client to a SSO Google auth page.

So! Let's call the `arc call-conduit` command ourselves, it takes JSON in and out. And we can still use all the types provided in gonduit, _aaaand_ it automatically pulls and handles the `.arcrc` so if `arc` is set up then this just works too.

# Commands

## `whoamiphab`

Returns your user's logged in PHID to stdout

## `whoisphab`

Returns some information about a user by PHID

## `statusesphab`

Lists all possible maniphest/differential task statuses
