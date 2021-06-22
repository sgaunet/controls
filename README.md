# controls

Little CLI to launch series of controls:

* SSH
* Get zabbix Problems
* Get connections number of postgresql Database
* Check HTTP requests

The CLI prints controls on stdout and save it in a markdown report.

**This tool is under development.**


# Example

Configuration file :

```
db:
  - dbhost: db1.pg.local
    dbuser: postgres
    dbpassword: ...
    dbport: 5432
    dbname: mydb
    sizelimit: 400    # Go
  - dbhost: db2.pg.local
    dbuser: postgres
    dbpassword: 
    dbport: 5432
    dbname: mydb
    sizelimit: 400

sshAsserts:
  GROUP1:
    - cmd: "systemctl status ntpd > /dev/null 2>&1; echo $?"
      expected: "0"
  GROUP2:
    - cmd: script.sh | awk '($3=="running"){ print $0 }' | wc -l
      expected: "18"
    - cmd: echo foo
      expected: "foo"

sshServers:
  GROUP1:
    - host: host1.mysociety.local
      user: ubuntu
      password: 
      sshkey:                         # prefered
    - host: host2.mysociety.local
      user: ec2-user
      password: 
      sshkey: 
  GROUP2:
    - host: host3.mysociety.local
      user: myuser
      password: 
      sshkey: 

assertsHTTP:
  - host: http://internal-e2734db1-foo-bar-66vy-345675678.eu-west-3.elb.amazonaws.com
    hostheader: real.dns.fr

zbxCtl:
  apiEndpoint: "http://zbx.mysociety.local/api_jsonrpc.php"
  user: "LOGIN"
  password: ""
  since: 172800   # 2 days
  severityThreshold: 4
```

