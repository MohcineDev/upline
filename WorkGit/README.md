### Initialized empty Git repository in /home/mohdev/Desktop/github/WorkGit/.git/
git init

### check the status
git status

### use "git add <file>..." to include in what will be committed
git add .


### commit the changes, the working tree should be clean.
git commit -m "change gile content"

### stage the changed file
git add hello.sh

# History

### Show the history of the working directory.
git log

### Show One-Line History for a condensed view showing only commit hashes and messages.
git log --oneline

## Controlled Entries: 
git log --oneline -n 2 --since="5 minutes ago"

## Personalized Format: 