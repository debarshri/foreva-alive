#### Foreva alive
A go based app that keep other apps alive even after they die due to exception

Service management utility

go build faserve.go


App comes back to live when it dies due to any exception

Execute

```
faserve add test "ls -ltr"
faserve list
faserve start test
faserve delete test
faserve delete-all

```




