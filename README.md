# 目的
以下の手順{A,B,C}で、どのような挙動をするか、また違いがあるか？を調べる。
## 手順A
- golang の処理の中で shell script を実行
- その shell script の中で nohup で別の shell script を実行
- golangのプロセスをkill
## 手順B
- golang の処理の中で nohup で shell script を実行
- その shell script の中で nohup で別の shell script を実行
- golangのプロセスをkill
## 手順C
- golang の処理の中で nohup で shell script を実行
- その shell script の中で別の shell script を実行
- golangのプロセスをkill

# 背景
- ラズパイ4の環境で手順Aでコードが実行された時、nohupで呼び出された shell script のプロセスも死んでしまった。
- nohup は非同期での実行なので、親プロセスが死んでも挙動に影響がないと思ったので、違和感を感じた。
- 調べてみることにした。

# 結果



## go run ./go_exec/main.go false false

- 実行した後、goのプロセスをkillする → hup.sh も sleep.sh も動き続ける

```
$ps aux | grep -e sleep -e hup -e go_exec
tokinaga         75265   0.1  0.0 410068608    960 s020  S+    8:54AM   0:00.00 sleep 1
tokinaga         75267   0.0  0.0 410732240   1392 s018  S+    8:54AM   0:00.00 grep -e sleep -e hup -e go_exec
tokinaga         74893   0.0  0.0 410206416   1760 s020  S+    8:54AM   0:00.03 /bin/sh ./scripts/sleep.sh
tokinaga         74892   0.0  0.0 410200272   1568 s020  S+    8:53AM   0:00.00 /bin/sh ./scripts/hup.sh
tokinaga         74874   0.0  0.1 411381360  21840 s020  S+    8:53AM   0:00.10 go run go_exec/main.go false false

$kill -9 74874

$ps aux | grep -e sleep -e hup -e go_exec
tokinaga         75392   0.5  0.0 410068608    960 s020  S     8:54AM   0:00.00 sleep 1
tokinaga         74893   0.1  0.0 410206416   1760 s020  S     8:54AM   0:00.04 /bin/sh ./scripts/sleep.sh
tokinaga         75394   0.0  0.0 410724048   1296 s018  R+    8:54AM   0:00.01 grep -e sleep -e hup -e go_exec
tokinaga         74892   0.0  0.0 410200272   1568 s020  S     8:53AM   0:00.00 /bin/sh ./scripts/hup.sh
```

## go run ./go_exec/main.go false true

```
[tokinaga:/Users/tokinaga/develop/nohup_in_go]$ps aux | grep -e sleep -e hup -e go_exec
tokinaga         21565   0.1  0.0 410068608    960 s029  S+   10:34AM   0:00.00 sleep 1
tokinaga         21567   0.0  0.0 410733264   1456 s018  S+   10:34AM   0:00.00 grep -e sleep -e hup -e go_exec
tokinaga         21559   0.0  0.0 410203344   1680 s029  S+   10:34AM   0:00.00 /bin/sh ./scripts/sleep.sh
tokinaga         21543   0.0  0.1 411381360  21568 s029  S+   10:34AM   0:00.10 go run go_exec/main.go false true
[tokinaga:/Users/tokinaga/develop/nohup_in_go]$kill -9 21543
[tokinaga:/Users/tokinaga/develop/nohup_in_go]$ps aux | grep -e sleep -e hup -e go_exec
tokinaga         21690   0.4  0.0 410068608    960 s029  S    10:34AM   0:00.00 sleep 1
tokinaga         21559   0.1  0.0 410204368   1712 s029  S    10:34AM   0:00.01 /bin/sh ./scripts/sleep.sh
tokinaga         21692   0.0  0.0 410742480   1616 s018  S+   10:34AM   0:00.00 grep -e sleep -e hup -e go_exec
[tokinaga:/Users/tokinaga/develop/nohup_in_go]$
```

go のプロセス殺しても、stdoutに”main.go runnnin...” が流れ続ける・・・