# Patch testing

This repository represents an open-source repository.

# Creating the mirror

Create a new git repository with the following branch settings:

* Lock branch Branch is read-only. Users cannot push to the branch.
* Do not allow bypassing the above settings The above settings will apply to administrators and custom roles with the "bypass branch protections" permission.

**Create a new repository. Do not create a fork**

# Synchronizing the mirror

From the mirror repository, simply run:

```bash
git remote add upstream {upstream URL}
git fetch upstream main
git push origin --force 'refs/remotes/upstream/main:refs/heads/main'

# Update the local ref
git reset --hard origin/main
```

# Creating a patch

```
git clone {mirror}
git checkout -b "branch"
git add {changes}
git commit -m "commit message"
# git push will fail because {mirror} is read-only
# -1 will create a single patch from the previous commit.
git format-patch -1
```

# Submitting the patch

Open a pull request to the `patches` repository.
