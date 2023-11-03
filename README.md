# Git Flow and Pull Request Sandbox

A place to understand how to use Pull Requests with [git-flow](https://github.com/petervanderdoes/gitflow-avh).

We assume the following:

- You have [git flow (AVH Edition)](https://github.com/petervanderdoes/gitflow-avh) installed,
- You have [Github's CLI](https://cli.github.com) installed.

The example app is written in Go, but it does not really matter as we will not even compile it.

## Clone the repository

Using [gh](https://cli.github.com), let's clone this sandbox:

```sh
gh repo clone gildas/gitflow-pr-sandbox
```

## Initialize git flow

Once inside the repository, initialize [git flow](https://github.com/petervanderdoes/gitflow-avh):

```sh
git flow init
```

Make sure to:

- name the `develop` branch `dev` (I never liked that long label)
- set the version tag prefix to `v`

## Developing a feature PR style

Let's start our feature:

```sh
git flow feature start myfeature
```

Add and commit some changes...

You can share the feature with fellow developers by pushing the branch to the repository (do not use `git flow feature publish` since we will use it to create the Pull Request).

```sh
git push --set-upstream origin feature/myfeature
```

or

```sh
git push -u origin feature/myfeature
```

Once the feature is finished, publish the feature:

```sh
git flow feature publish myfeature
```

If you use my [git flow hooks](https://github.com/gildas/gitflow-hooks), the Pull Request will be automatically created.

Otherwise, create the Pull Request:

```sh
gh pr create \
  --title "Merge feature myfeature" \
  --body  "Feature myfeature" \
  --base  dev
```

Now, as the Pull Request reviewer, log on github and merge the Pull Request **without** deleting the feature branch. Or use the CLI:

```sh
gh pr merge \
  --subject 'Merged feature myfeature'
```

Back to the feature owner, grab the merge:

```sh
git checkout dev
git pull
```

And finish the feature:

```sh
git flow feature finish myfeature
```

## Baking a release PR style

Since preparing a release usually involves the master branch, typically, the people that take care of that task are usually the maintainers of the project.

Let's start a new release:

```sh
git flow release start 1.0.0
```

If you use my [git flow hooks](https://github.com/gildas/gitflow-hooks), you will get the auto-versioning for [node.js](https://nodejs.org) and [Go](https://go.dev) based projects.

If not, you should modify your project's version manually, in the current case:

```sh
sed -Ei '/VERSION\s+=/s/[0-9]+\.[0-9]+\.[0-9]+/1.0.0/' version.go
git add version.go
git commit -m "Bumped to version 1.0.0" version.go
```

You can also share the release branch with others by pushing it to the repository (like features):

```sh
git push -u origin release/1.0.0
```

When all bugs have been fixed and the QA process is done, the release should be published to created a Pull Request:

```sh
git flow release publish
```

If you use my [git flow hooks](https://github.com/gildas/gitflow-hooks), the Pull Request will be automatically created.

Otherwise, create the Pull Request:

```sh
gh pr create \
  --title "Merge release 1.0.0" \
  --body  "Release 1.0.0." \
  --base  master
```

Now, as the Pull Request reviewer, log on github and merge the Pull Request **without** deleting the release branch. Or use the CLI:

```sh
gh pr merge \
  --subject 'Merged release 1.0.0'
```

Back to the release _baker_, grab the merge:
```sh
git checkout master
git pull
```

And finish the release:
```sh
git flow release finish 1.0.0
```

## Writing a hotfix PR style

Let's start a new hotfix:

```sh
git flow hotfix start 1.0.1
```

If you use my [git flow hooks](https://github.com/gildas/gitflow-hooks), you will get the auto-versioning for [node.js](https://nodejs.org) and [Go](https://go.dev) based projects.

If not, you should modify your project's version manually, in the current case:

```sh
sed -Ei '/VERSION\s+=/s/[0-9]+\.[0-9]+\.[0-9]+/1.0.1/' version.go
git add version.go
git commit -m "Bumped to version 1.0.1" version.go
```

You can also share the hotfix branch with others by pushing it to the repository (like features and releases):

```sh
git push -u origin hotfix/1.0.1
```

When all bugs have been fixed and the QA process is done, the hotfix should be published to created a Pull Request:

```sh
git flow hotfix publish
```

If you use my [git flow hooks](https://github.com/gildas/gitflow-hooks), the Pull Request will be automatically created.

Otherwise, create the Pull Request:

```sh
gh pr create \
  --title "Merge hotfix 1.0.1" \
  --body  "Hotfix 1.0.1" \
  --base  master
```

Now, as the Pull Request reviewer, log on github and merge the Pull Request **without** deleting the hotfix branch. Or use the CLI:

```sh
gh pr merge \
  --subject 'Merged hotfix 1.0.1'
```

Back to the hotfix maker, grab the merge:
```sh
git checkout master
git pull
```

And finish the hotfix:
```sh
git flow hotfix finish 1.0.1
```
