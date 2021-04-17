[ What is this? ](#What is this?)

[ How to build? ](#How to build)

[ How to run ](#How to run)

[ Sample output ](#Sample output)


## What is this?
Strip api.github.com` metadata to retreive repository information
by user name`\
Now it can `Create a new repo` or `Create a lot of repos from a file`\
Now it can `Inject a webhook in a repo` or `Concurrently inject lots of webhooks into lots of repos`\

## How to build
```
go build main.go
```

## How to run
#### List all the repositories under user
```
./main f <githubusername>
viz: ./main f mainak90
```

#### Create a gist inside your account
```
./main create <name> <description> <text-file>
viz: ./main create testdoc "test document" file.txt
```

#### Create a new repo inside your authenticated account
```
./main newrepo <reponame>
viz: ./main newrepo test-repo
```

#### Create a list of repos(empty) in your autheticated github account
This repo file should be a text document with newline seperated list of repos.
```
./main newrepos <repo-file-path>
./main newrepos tests/repos.txt
```

#### Setup a webhook inside a github repo(Authenticated)
```
./main pushwebhook <owner> <repo> <hookurl>
./main pushwebhook mainak90 GitGists https://www.example.com/hook
```

#### Setup webhooks inside a lot of repos via a json file
This one needs a properly formatted json document that has a the structure defined on `tests/hooklist.json`
Please refer to struct `models.WebhookList` to get the gist of things.
```
./main pushwebhooks <json-file-path>
./main pushwebhooks tests/hooklist.json
```

#### Get all repos from a GitHub organization
WIP...


## Sample output
```
2020/11/18 00:32:38 [{95495190 Angular-ShoppingCart mainak90/Angular-ShoppingCart 0 false} {236586823 ansible-kafka mainak90/ansible-kafka 0 false} {84765270 Ansible-Trials mainak90/Ansible-Trials 0 false} {234415237 ansible.role.jenkins mainak90/ansible.role.jenkins 0 false} {244761352 AWS-CI-CD mainak90/AWS-CI-CD 0 false} {95494568 chat-cat mainak90/chat-cat 0 false} {142584392 chef-fluency-badge mainak90/chef-fluency-badge 0 false} {230978751 cloudformation-pipelines mainak90/cloudformation-pipelines 0 false} {256621936 content-kubernetes-security-ac mainak90/content-kubernetes-security-ac 0 false} {220959059 crud-mongo mainak90/crud-mongo 0 false} {98127189 customer-search-api mainak90/customer-search-api 0 false} {123824885 docker-webapp mainak90/docker-webapp 0 false} {249078099 flask-auth-jwt mainak90/flask-auth-jwt 0 false} {220850422 go-crud-jwt mainak90/go-crud-jwt 0 false} {219611918 go-mux-postgres mainak90/go-mux-postgres 0 false} {236590586 groovy-goodness mainak90/groovy-goodness 0 false} {225657230 grpc-mongo mainak90/grpc-mongo 0 false} {98129214 hello-couchbase mainak90/hello-couchbase 0 false} {234418974 helm-nginx mainak90/helm-nginx 0 false} {275845900 helmer mainak90/helmer 0 false} {232927203 hvac-openshift-feeder mainak90/hvac-openshift-feeder 0 false} {268463203 katacoda-scenarios mainak90/katacoda-scenarios 0 false} {239371675 kube-test-ext mainak90/kube-test-ext 0 false} {259121732 kube-vault-loader mainak90/kube-vault-loader 0 false} {124057273 lamp-docker mainak90/lamp-docker 0 false} {95494797 MyQRCode mainak90/MyQRCode 0 false} {105584335 nodejs-ex mainak90/nodejs-ex 0 false} {95494284 noteapp mainak90/noteapp 0 false} {303825242 openshift-secret-manager mainak90/openshift-secret-manager 0 false} {244753172 orderbook-parser mainak90/orderbook-parser 0 false}]
```