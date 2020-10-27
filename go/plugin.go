package go_flutter_systray

import (
	"encoding/json"

	"github.com/getlantern/systray"
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const channelName = "go_flutter_systray"

// MenuItemEntry menuitem
type MenuItemEntry struct {
	Key     string          `json:"key"`
	Icon    []byte          `json:"icon"`
	Title   string          `json:"title"`
	Tooltip string          `json:"tooltip"`
	Child   []MenuItemEntry `json:"child"`
}

// GoFlutterSystrayPlugin implements flutter.Plugin and handles method.
type GoFlutterSystrayPlugin struct {
	channel  *plugin.MethodChannel
	window   *glfw.Window
	menuList map[string]*systray.MenuItem
}

var _ flutter.Plugin = &GoFlutterSystrayPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *GoFlutterSystrayPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	p.channel = plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	p.channel.HandleFunc("hideWindow", p.hideWindow)
	p.channel.HandleFunc("showWindow", p.showWindow)
	p.channel.HandleFunc("runSystray", p.runSystray)

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
	return nil
}

func (p *GoFlutterSystrayPlugin) callHandler(methodName string, arguments interface{}) error {
	return p.channel.InvokeMethod(methodName, arguments)
}

func (p *GoFlutterSystrayPlugin) hideWindow(arguments interface{}) (reply interface{}, err error) {
	p.window.Hide()
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) showWindow(arguments interface{}) (reply interface{}, err error) {
	p.window.Show()
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) runSystray(arguments interface{}) (reply interface{}, err error) {
	params, _ := arguments.([]interface{})
	var mainMenu MenuItemEntry
	exitMethod := params[1].(string)
	if err := json.Unmarshal([]byte(params[0].(string)), &mainMenu); err != nil {
		return nil, err
	}
	systray.Run(func() {
		systray.SetIcon(mainMenu.Icon)
		systray.SetTitle(mainMenu.Title)
		systray.SetTooltip(mainMenu.Tooltip)
		for _, item := range mainMenu.Child {
			p.menuList[item.Key] = p.putMenuItem(nil, item)
		}
	}, func() {
		p.callHandler(exitMethod, nil)
	})
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) putMenuItem(menuItem *systray.MenuItem, entry MenuItemEntry) *systray.MenuItem {
	var menu *systray.MenuItem
	if menuItem == nil {
		menu = systray.AddMenuItem(entry.Title, entry.Tooltip)
	} else {
		menu = menuItem.AddSubMenuItem(entry.Title, entry.Tooltip)
	}
	menu.SetIcon(entry.Icon)
	for _, item := range entry.Child {
		if item.Key == "" {
			systray.AddSeparator()
		} else {
			p.menuList[item.Key] = p.putMenuItem(menu, item)
		}
	}
	return menu
}

func (p *GoFlutterSystrayPlugin) setIcon(arguments interface{}) (reply interface{}, err error) {
	params, _ := arguments.([]interface{})
	iconBytes := params[0].([]byte)
	systray.SetIcon(iconBytes)
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) setTitle(arguments interface{}) (reply interface{}, err error) {
	params, _ := arguments.([]interface{})
	var title string
	title = params[0].(string)
	systray.SetTitle(title)
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) setTooltip(arguments interface{}) (reply interface{}, err error) {
	params, _ := arguments.([]interface{})
	var tooltip string
	tooltip = params[0].(string)
	systray.SetTooltip(tooltip)
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) itemCheck(arguments interface{}) (reply interface{}, err error) {
	params, _ := arguments.([]interface{})
	var tooltip string
	tooltip = params[0].(string)
	systray.SetTooltip(tooltip)
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) itemUncheck(arguments interface{}) (reply interface{}, err error) {
	params, _ := arguments.([]interface{})
	var tooltip string
	tooltip = params[0].(string)
	systray.SetTooltip(tooltip)
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) itemChecked(arguments interface{}) (reply interface{}, err error) {
	params, _ := arguments.([]interface{})
	var tooltip string
	tooltip = params[0].(string)
	systray.SetTooltip(tooltip)
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) itemDisable(arguments interface{}) (reply interface{}, err error) {
	params, _ := arguments.([]interface{})
	var tooltip string
	tooltip = params[0].(string)
	systray.SetTooltip(tooltip)
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) itemEnable(arguments interface{}) (reply interface{}, err error) {
	params, _ := arguments.([]interface{})
	var tooltip string
	tooltip = params[0].(string)
	systray.SetTooltip(tooltip)
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) itemDisabled(arguments interface{}) (reply interface{}, err error) {
	params, _ := arguments.([]interface{})
	var tooltip string
	tooltip = params[0].(string)
	systray.SetTooltip(tooltip)
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) itemHide(arguments interface{}) (reply interface{}, err error) {
	params, _ := arguments.([]interface{})
	var tooltip string
	tooltip = params[0].(string)
	systray.SetTooltip(tooltip)
	return nil, nil
}

func (p *GoFlutterSystrayPlugin) itemShow(arguments interface{}) (reply interface{}, err error) {
	params, _ := arguments.([]interface{})
	var tooltip string
	tooltip = params[0].(string)
	systray.SetTooltip(tooltip)
	return nil, nil
}
