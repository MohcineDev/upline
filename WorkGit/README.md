### Initialized empty Git repository in /home/mohdev/Desktop/github/WorkGit/.git/
git init

### check the status
```
git status
```
### use "git add <file>..." to include in what will be committed
```
git add .
```

### commit the changes, the working tree should be clean.
``
git commit -m "change gile content"
``
### stage the changed file
> git add hello.sh

# History

### Show the history of the working directory.
> git log

### Show One-Line History for a condensed view showing only commit hashes and messages.
git log --oneline

## Controlled Entries: 
```
git log --oneline -n 2 --since="5 minutes ago"
```
## Personalized Format: 
git log --oneline -n 1 --format="* %h %as | %s %d [%an]"


# Check it out
### Restore First Snapshot: 
git checkout f719360
cat hello.sh

### Restore Second Recent Snapshot: 
git checkout 4f82212
cat hello.sh

### Return to Latest Version: 
git checkout master


# TAG me
### Referencing Current Version: 
git tag v1

### Tagging Previous Version: 
git tag v1-beta HEAD~1

