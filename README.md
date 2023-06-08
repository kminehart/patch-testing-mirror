# Patch testing

This repository represents an open-source repository.

# Creating the mirror

Create a new git repository with the following branch settings:

* Lock branch Branch is read-only. Users cannot push to the branch.
* Do not allow bypassing the above settings The above settings will apply to administrators and custom roles with the "bypass branch protections" permission.

# Synchronizing the mirror

From the mirror repository, simply run:

```bash
git remote add upstream {upstream URL}
git fetch upstream main
git push origin --force 'refs/remotes/upstream/main:refs/heads/main'
```
