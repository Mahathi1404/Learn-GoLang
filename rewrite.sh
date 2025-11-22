# Define the bad email (the one currently in your commit messages)
OLD_EMAIL="mahathi14041@outlook.com"
# Define the correct protected email
CORRECT_EMAIL="61062537+Mahathi1404@users.noreply.github.com"
# Define your GitHub username
CORRECT_NAME="Mahathi1404"


# This command rewrites the email and name on every commit in the current branch

git filter-branch --env-filter '
    export GIT_COMMITTER_NAME="'$CORRECT_NAME'"
    export GIT_AUTHOR_NAME="'$CORRECT_NAME'"
    export GIT_COMMITTER_EMAIL="'$CORRECT_EMAIL'"
    export GIT_AUTHOR_EMAIL="'$CORRECT_EMAIL'"
' --msg-filter 'sed "s/'$OLD_EMAIL'/'$CORRECT_EMAIL'/g"' --tag-name-filter cat -- --branches
