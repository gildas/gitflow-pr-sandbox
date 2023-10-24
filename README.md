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

## Developing a feature PR style

Let's start our feature:

```
git flow feature start myfeature
```

Add and commit some changes...

You can share the feature with fellow developers by pushing the branch to the repository (do not use `git flow feature publish` since we will use it to create the PR).

```
git push --set-upstream origin feature/myfeature
```

or

```
git push -u origin feature/myfeature
```

Once the feature is finished, publish the feature:

```
git flow feature publish myfeature
```

And create the PR:

```
gh pr create \
  --title "Merge feature myfeature" \
  --body  "Feature myfeature" \
  --base  dev \
  --repo  gildas/gitflow-pr-sandbox
```

Now, as the PR reviewer, log on github and merge the PR **without** deleting the feature branch. Or use the CLI:

```
gh pr merge \
  --subject 'Merged feature myfeature'
```

Back to the feature owner, grab the merge:

```
git checkout dev
git pull
```

And finish the feature:

```
git flow feature finish myfeature
```

## TODO

- Use [gildas/gitflow-hooks](https://github.com/gildas/gitflow-hooks) to initalize and configure git flow.
- Implement all this with [bitbucket](https://bitbucket.org).
