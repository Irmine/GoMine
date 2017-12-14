$host.ui.RawUI.WindowTitle = "GoMine : A Minecraft Bedrock Edition server software in Go"
$Loop = "false"

$env:GOPATH = "${PSScriptRoot}\"
ECHO($GOPATH)

function StartServer{
    $command = "go build -i -o ./GoMine.exe ./src/main.go"
    $command2 = "./GoMine.exe"
    iex $command
    iex $command2
}

if (!(Get-Command "go" -ErrorAction SilentlyContinue)){ 
    echo('You require Go / Golang to run this program!');
    echo('Download it from https://golang.org/ and try again!');
    exit 1
} else {
    $loops = 1
    StartServer
    while($Loop -eq "true") {
       	if($loops -ne 0){
		echo ("Restarted " + $loops + " time(s)")
	}
	$loops++
	echo "To escape the loop, press CTRL+C now. Otherwise, wait 5 seconds for the server to restart."
	echo ""
	Start-Sleep 5
	StartServer        
    }
    cmd /c pause | Out-Null
}