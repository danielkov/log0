# How to contribute

Thank you for considering contributing to `log0`. It is a small project and so the tiniest contributions can benefit it a lot, which is why I invite you to read the following guidelines on how to contribute to this project.

## Getting Started

 - Submit a ticket for your issue, assuming one does not already exist. Use the [Issue Template](./ISSUE_TEMPLATE.md) and follow the checklist.
 - Clearly describe the issue including steps to reproduce when it is a bug.
 - Make sure you fill in the earliest version that you know has the issue.
 - Fork the repository on GitHub.
 - Make the changes you deem necessary
 - Submit a pull request, using the [Pull Request Template](./PULL_REQUEST_TEMPLATE.md) on the appropriate branch.

## Version control

The main branch of this project is `master`, which is also the branch that has the current latest version of the project that was released. Release branches are named like `v1.2`, where **1** is the major version and **2** is the minor version. Minor releases arew used for bug fixes, while major releases are for breaking changes. This ensures that using services like [gopkg.in](https://gopkg.in) yield the best results.

The current tip of developments is the `development` branch, which is usually what gets converted into a new release, when enough changes get merged. This is the branch you should submit pull requests to, unless there is a special feature branch that is not yet to be merged onto development, in which case use that branch for PR's.

Only target release branches if you are certain your fix must be on that branch. Please avoid working directly on the master branch.

Make regular commits of logical units and make sure commit history stays nice and clean to help your future-self and others.

Check for unnecessary whitespace with git diff --check before committing.

Make sure your commit messages are in a clear and understandable format and in English language.

Make sure you have added the necessary tests for your changes.

Make sure that your changes do not decrease test coverage and that the project still works with all previous tests with your changes, unless otherwise needed (regression).

Make sure your changes also pass the linters, like `golint` and `gometalint`. If unsure, use `gometalint ./...` and fix until you see no warnings.

## Documenting changes

When making changes that also change the general public API of the project, (e.g.: any public functions or structs) make sure you also reflect these changes in the documentation so that users can make use of your changes. Due to this being a small project, there is no separate WIKI folder, however the [Readme](./README.md) should be kept clean and structured logically.

## Submitting Changes

**By submitting your work you release it under the project's [MIT license](./LICENSE.MD).**

Push your changes to a topic branch in your fork of the repository.

Submit your pull request from your topic branch to the appropriate branch of the project (usually `development`).

Make sure you link the issue number in your pull request.

Include a link to the pull request in the issue ticket, if there is one (there should be).

Optionally, @tag someone in the PR if you want multiple people to review it.

After feedback has been given, follow it if changes are necessary, if not, once all pipelines success, your PR should be accepted and you will be notified if your changes make it to a release.
