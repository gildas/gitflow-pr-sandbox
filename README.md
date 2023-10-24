# Git Flow and Pull Request Sandbox

A place to understand how to use Pull Requests with [git-flow](https://github.com/petervanderdoes/gitflow-avh).

We assume the following:

- You have [git flow (AVH Edition)](https://github.com/petervanderdoes/gitflow-avh) installed,
- You have [Github's CLI](https://cli.github.com) installed.

The example app is written in Go, but it does not really matter as we will not even compile it.

## Clone the repository

Using [gh](https://cli.github.com), let's clone this sandbox:

```
gh repo clone gildas/gitflow-pr-sandbox
```

## Initialize git flow

Once inside the repository, initialize [git flow](https://github.com/petervanderdoes/gitflow-avh):

```
git flow init
```

Make sure to:

- name the `develop` branch `dev` (I never liked that long label)
- set the version tag prefix to `v`

## TODO

- Use [gildas/gitflow-hooks](gildas/gitflow-hooks) to initalize and configure git flow.
- Implement all this with [bitbucket](https://bitbucket.org).
