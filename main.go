package main

import (
	"fmt"
    "net"
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
		turnOff := systray.AddMenuItem("Turn Off", "Turn Light Off")
		systray.AddSeparator()
		turnBright := systray.AddMenuItem("Bright", "Bright")
		turnNight := systray.AddMenuItem("Night", "Night")
		turnFocus := systray.AddMenuItem("Focus", "Focus")
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
		ip_add := "192.168.1." +strconv.Itoa(i)+ ":38899"
		conn, err := net.Dial("udp", ip_add)
		if err != nil {
			return
		}

		if sceneId == 0{
			fmt.Fprintf(conn, "{ \"method\": \"setState\", \"params\": { \"state\": false  } }")
		}else if sceneId == 1{
			fmt.Fprintf(conn, "{ \"method\": \"setState\", \"params\": { \"state\": true  } }")
		}else {
			fmt.Fprintf(conn, "{ \"method\": \"setState\", \"params\": { \"sceneId\": " +strconv.Itoa(sceneId) +"  } }")
		}

		conn.Close()
	}
}