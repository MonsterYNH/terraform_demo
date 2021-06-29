### How to use
## run mockapi server
```
$ export MONSTER_ACCOUNT=account
$ export MONSTER_PASSWORD=password
$ cd mockapi/bin
$ go run main.go
```
## test with terraform plugin
```
$ go build -o ~/.terraform.d/plugins/terraform.ynhmonster/monsterworlds/monster/1.0.0/linux_amd64/terraform-provider-monster
$ cd example/vdc
$ rm -rf .terraform
$ rm .terraform.lock.hcl
$ rm terraform.tfstate
$ rm terraform.tfstate.backup
$ rm vdc.json
$ terraform init
$ terraform plan
$ terraform apply
$ terraform destroy
```
## test with unit test
```
$ cd monster
$ go test -v
```