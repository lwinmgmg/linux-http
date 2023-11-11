# LINUX HTTP

This software is design to run linux cmd in the hosted machine through http request.

## APIs

```
[GIN-debug] POST   /api/v1/func/linux/execute
```

Example Input Json
```
{
    "cmd": "sudo apt upgrade",
    "input": "Y"
}
or 
{
    "cmd": "sudo apt -y upgrade",
    "input": ""
}
```

Success Output Json
```
{
    "output": "api\ndb\nenv\ngo.mod\ngo.sum\nmain.go\nmiddleware\nmodels\nReadMe.md\nservices\nutils\n",
    "exit_code": 0
}
```

## OAuth Authentication
Claim Format
```
{
  "iss": "github",
  "exp": 1699208521,
  "iat": 1699204921
}
```

Key is the hash value of the combination of $LH_SECRET+$iat.

## Command Limit

We can limit cmd through $LH_ALLOW_CMDS settings.

Eg. if we want to allow 'ls' and 'ls -ahl' cmd only
```
LH_ALLOW_CMDS="ls -ahl|||ls"
```
As you see, '|||' delimiter is required for multiple cmds.
All cmds will be allowed if $LH_ALLOW_CMDS settings is not set.
