package main

import (
	"fmt"
    "net"
    "bufio"
    "strconv"
	"github.com/getlantern/systray"
	"WizLightTray/icon"
)

func main() {
	onExit := func() {
	}

	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTitle("Wiz Light Tray")
	systray.SetTooltip("Control Wiz Light")

	go func() {
		turnOn := systray.AddMenuItem("Turn On", "Turn Light On")
		turnOff := systray.AddMenuItem("Turn Off", "Turn Light On")
		systray.AddSeparator()
		turnBright := systray.AddMenuItem("Bright", "Turn Light On")
		turnNight := systray.AddMenuItem("Night", "Turn Light On")
		turnFocus := systray.AddMenuItem("Focus", "Turn Light On")
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Exit", "Quit the whole app")


		for {
			select {
			case <-turnOn.ClickedCh:
				set_scene_wrapper(1)
			case <-turnOff.ClickedCh:
				set_scene_wrapper(0)
			case <-turnBright.ClickedCh:
				set_scene_wrapper(12)
			case <-turnNight.ClickedCh:
				set_scene_wrapper(14)
			case <-turnFocus.ClickedCh:
				set_scene_wrapper(15)
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
}


func set_scene_wrapper(sceneId int){
    for i := 0; i <= 255; i++ {   
        go set_scene(i, sceneId)
    }
}


func set_scene(i int, sceneId int){
    ip_add := "192.168.1." +strconv.Itoa(i)+ ":38899"
    conn, err := net.Dial("udp", ip_add)
    if err != nil {
        fmt.Printf("Some error %v", err)
        return
    }
	if sceneId == 0{
		fmt.Fprintf(conn, "{ \"method\": \"setState\", \"params\": { \"state\": false  } }")
	}else if sceneId == 1{
		fmt.Fprintf(conn, "{ \"method\": \"setState\", \"params\": { \"state\": true  } }")
	}else {
		fmt.Fprintf(conn, "{ \"method\": \"setState\", \"params\": { \"sceneId\": " +strconv.Itoa(sceneId) +"  } }")

	}
    
    p :=  make([]byte, 2048)
    _, err = bufio.NewReader(conn).Read(p)
    if err == nil {
        fmt.Printf("%s\n", p)
    } else {
        fmt.Printf("Some error %v\n", err)
    }
    conn.Close()
}