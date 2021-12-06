curl -s https://ytrack.learn.ynov.com/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"GROUX\"}}){id}}"}'| sed 's/[^0-9]*//g'
