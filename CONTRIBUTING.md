## Contribution guidelines

If you want to contribute something, then submit a pull request, and the maintainers will look at it. Then stuff will happen.

*Notes:*
* Please keep code concise, and be sure to include any change summaries in the PR documentation.
* Looks for previous issues/PRs before starting to develop to avoid overlap with features/PRs. I don't want to waste your time.
* I am open to new ideas/features that could be introduced to make this little daemon better/more useful. If you have any that are big/introduce breaking changes, send a message before dedicating a lot of time to it so it can be verified that this is where we want to head and it won't introduce any breaking changes. Again, I  don't want to waste your time.
* If you want to work on current development, look at the below section.

### Features requiring more work
This daemon is still a bit of a work in progress, so if you want to contribute to its development, here are some sections that I am working on implementing.

* Adding more notifiers (look in `src/backend`).
* Dynamic configuration for notifiers.
*Potentially to be removed/unsupported*
* Updating DNS entries based on an IP change.
* Implementing more DNS updaters rather than just Cloudflare. 