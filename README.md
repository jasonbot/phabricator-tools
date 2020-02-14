# phabricator-tools
A collection of command line tools for manipulating stuff on Phabricator

Phabricator is nice in how powerful it is, but much like other general purpose tools like Jira, the UI leaves much to be desired for focused, simple workflows.

I wish to solve that for myself.


## First problem

I can't directly use the conduit endpoint via conduit; every time I hit it, it returns HTML trying to redirect my client to a SSO Google auth page.

So! Let's call the `arc call-conduit` command ourselves, it takes JSON in and out. And we can still use all the types provided in gonduit, _aaaand_ it automatically pulls and handles the `.arcrc` so if `arc` is set up then this just works too.

## Second problortunity

Gonduit doesn't support all the modern methods (like `user.search`), and some of tghe entities don 't use the right named fields for the installation at work (looking at you, `User`) so I can just drop gonduit altogether at this point.

# Commands

## `phabwhoami`

Returns your user's logged in PHID to stdout

## `phabwhois`

Returns some information about a user by PHID

## `phabstatuses`

Lists all possible maniphest/differential task statuses

## `anchormanagement`

Fast text-based UI for doing daily maniphest grooming
