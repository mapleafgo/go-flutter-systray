package go_flutter_systray

import (
	"encoding/json"
	"log"

	"github.com/getlantern/systray"
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const channelName = "go_flutter_systray"

// MenuItemEntry menuitem
type MenuItemEntry struct {
	Key        string          `json:"key"`
	Icon       []byte          `json:"icon"`
	Title      string          `json:"title"`
	Tooltip    string          `json:"tooltip"`
	IsCheckbox bool            `json:"isCheckbox"`
	Child      []MenuItemEntry `json:"child"`
}

const (
	quitCallMethod string = "systray_quit_call"
	mainMenuKey    string = "main"
)

// isSureExit 是否退出
var isSureExit bool

// isInit 是否初始化过
var isInit bool

// GoFlutterSystrayPlugin implements flutter.Plugin and handles method.
type GoFlutterSystrayPlugin struct {
	filterExit bool
	channel    *plugin.MethodChannel
	window     *glfw.Window
	menuList   map[string]*systray.MenuItem
}

// Default 默认接管窗口关闭
var Default = &GoFlutterSystrayPlugin{
	filterExit: true,
}

var _ flutter.Plugin = &GoFlutterSystrayPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *GoFlutterSystrayPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	p.channel = plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	p.channel.HandleFunc("hideWindow", p.hideWindow)
	p.channel.HandleFunc("showWindow", p.showWindow)
	p.channel.HandleFunc("exitWindow", p.exitWindow)
	p.channel.HandleFunc("runSystray", p.runSystray)
	p.channel.HandleFunc("quitSystray", p.quitSystray)

	p.channel.HandleFunc("setIcon", p.setIcon)
	p.channel.HandleFunc("setTitle", p.setTitle)
	p.channel.HandleFunc("setTooltip", p.setTooltip)
	p.channel.HandleFunc("itemCheck", p.itemCheck)
	p.channel.HandleFunc("itemUncheck", p.itemUncheck)
	p.channel.HandleFunc("itemChecked", p.itemChecked)
	p.channel.HandleFunc("itemDisable", p.itemDisable)
	p.channel.HandleFunc("itemEnable", p.itemEnable)
	p.channel.HandleFunc("itemDisabled", p.itemDisabled)
	p.channel.HandleFunc("itemHide", p.itemHide)
	p.channel.HandleFunc("itemShow", p.itemShow)
	return nil
}

// InitPluginGLFW is called after the call to InitPlugin. When an error is
// returned it is printend the application is stopped.
func (p *GoFlutterSystrayPlugin) InitPluginGLFW(window *glfw.Window) error {
	p.window = window
	p.window.SetCloseCallback(func(w *glfw.Window) {
		if p.filterExit && !isSureExit {
			w.SetShouldClose(false)
			w.Hide()
		}
	})
	return nil
}

func (p *GoFlutterSystrayPlugin) callHandler(methodName string, arguments interface{}) error {
	return p.channel.InvokeMethod(methodName, arguments)
}

func (p *GoFlutterSystrayPlugin) exitWindow(interface{}) (reply interface{}, err error) {
	isSureExit = true
	p.window.SetShouldClose(true)
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) hideWindow(interface{}) (reply interface{}, err error) {
	p.window.Hide()
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) showWindow(interface{}) (reply interface{}, err error) {
	p.window.Show()
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) runSystray(arguments interface{}) (reply interface{}, err error) {
	if isInit {
		log.Print("do not repeat initialization")
		return
	}
	mainMenu := &MenuItemEntry{}
	if err := json.Unmarshal([]byte(arguments.(string)), &mainMenu); err != nil {
		return nil, err
	}

	onReady := func() {
		systray.SetIcon(mainMenu.Icon)
		systray.SetTitle(mainMenu.Title)
		systray.SetTooltip(mainMenu.Tooltip)
		if len(mainMenu.Child) > 0 {
			p.menuList = make(map[string]*systray.MenuItem)
			for _, item := range mainMenu.Child {
				p.putMenuItem(nil, item)
			}
		}
		// 初始化完成
		isInit = true
	}

	systray.Register(onReady, func() {
		_ = p.callHandler(quitCallMethod, nil)
	})
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) putMenuItem(menuItem *systray.MenuItem, entry MenuItemEntry) {
	var menu *systray.MenuItem
	if len(entry.Key) == 0 {
		if menuItem == nil {
			systray.AddSeparator()
		}
		return
	}

	if menuItem == nil {
		if entry.IsCheckbox {
			menu = systray.AddMenuItemCheckbox(entry.Title, entry.Tooltip, false)
		} else {
			menu = systray.AddMenuItem(entry.Title, entry.Tooltip)
		}
	} else {
		if entry.IsCheckbox {
			menu = menuItem.AddSubMenuItemCheckbox(entry.Title, entry.Tooltip, false)
		} else {
			menu = menuItem.AddSubMenuItem(entry.Title, entry.Tooltip)
		}
	}
	p.menuList[entry.Key] = menu
	if entry.Icon != nil {
		menu.SetIcon(entry.Icon)
	}

	go p.startChan(entry.Key, menu)
	for _, item := range entry.Child {
		p.putMenuItem(menu, item)
	}
}

func (p *GoFlutterSystrayPlugin) startChan(key string, menu *systray.MenuItem) {
	for {
		<-menu.ClickedCh
		if err := p.callHandler(key, nil); err != nil {
			log.Panicln(err)
		}
	}
}

func (p *GoFlutterSystrayPlugin) quitSystray(interface{}) (reply interface{}, err error) {
	systray.Quit()
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) setIcon(arguments interface{}) (reply interface{}, err error) {
	params := arguments.([]interface{})
	key, iconBytes := params[0].(string), params[1].([]byte)
	if key == mainMenuKey {
		systray.SetIcon(iconBytes)
	} else {
		p.menuList[key].SetIcon(iconBytes)
	}
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) setTitle(arguments interface{}) (reply interface{}, err error) {
	params := arguments.([]interface{})
	key, title := params[0].(string), params[1].(string)
	if key == mainMenuKey {
		systray.SetTitle(title)
	} else {
		p.menuList[key].SetTitle(title)
	}
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) setTooltip(arguments interface{}) (reply interface{}, err error) {
	params := arguments.([]interface{})
	key, tooltip := params[0].(string), params[1].(string)
	if key == mainMenuKey {
		systray.SetTooltip(tooltip)
	} else {
		p.menuList[key].SetTooltip(tooltip)
	}
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) itemCheck(arguments interface{}) (reply interface{}, err error) {
	key := arguments.(string)
	p.menuList[key].Check()
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) itemUncheck(arguments interface{}) (reply interface{}, err error) {
	key := arguments.(string)
	p.menuList[key].Uncheck()
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) itemChecked(arguments interface{}) (reply interface{}, err error) {
	key := arguments.(string)
	return p.menuList[key].Checked(), nil
}

func (p *GoFlutterSystrayPlugin) itemDisable(arguments interface{}) (reply interface{}, err error) {
	key := arguments.(string)
	p.menuList[key].Disable()
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) itemEnable(arguments interface{}) (reply interface{}, err error) {
	key := arguments.(string)
	p.menuList[key].Enable()
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) itemDisabled(arguments interface{}) (reply interface{}, err error) {
	key := arguments.(string)
	return p.menuList[key].Disabled(), nil
}

func (p *GoFlutterSystrayPlugin) itemHide(arguments interface{}) (reply interface{}, err error) {
	key := arguments.(string)
	p.menuList[key].Hide()
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) itemShow(arguments interface{}) (reply interface{}, err error) {
	key := arguments.(string)
	p.menuList[key].Show()
	return nil, nil
}
