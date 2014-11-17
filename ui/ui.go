package ui

import(
	"github.com/conformal/gotk3/gtk"
	"log"

 )

type MainWindow struct {
 	Window *gtk.Window
 	headerbar *gtk.HeaderBar
 	NewTaskButton *gtk.Button
 	TaskList *gtk.ListBox
 	MainBox *gtk.Box


 	aboutItem *gtk.MenuItem
 	err error


}
/* Constructor */
func MainWindowNew(windowtype gtk.WindowType) *MainWindow {
 	win := new(MainWindow)
 	win.Window, win.err = gtk.WindowNew(windowtype)

 	if win.err != nil {
 		log.Fatal("unable to create window:", win.err)
 	}
 	win.InitializeHeaderBar()

 	win.Window.Connect("destroy", func(){
 		gtk.MainQuit()
 	})
 	win.TaskList, win.err = gtk.ListBoxNew()
	win.MainBox, win.err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2)

	win.aboutItem, win.err = gtk.MenuItemNewWithLabel("About")


	win.MainBox.Add(win.TaskList)

	win.Window.Add(win.MainBox)
 	win.Window.SetSizeRequest(800,600)



 	return win
}

/* Methods */

func ( win *MainWindow ) InitializeHeaderBar() {
	win.headerbar, win.err = gtk.HeaderBarNew()

	win.NewTaskButton, win.err = gtk.ButtonNewWithLabel("Add Task")
	win.NewTaskButton.Connect("clicked",  func() { win.NewTaskClick() })

	win.headerbar.PackStart(win.NewTaskButton)
	win.headerbar.SetShowCloseButton(true);
	win.headerbar.SetTitle("Tasks");
	win.Window.SetTitlebar(win.headerbar)
}

func (win *MainWindow) NewTaskClick() {
	var dialog *gtk.Dialog
	var taskEntry *gtk.Entry
	var dialogContent *gtk.Box
	dialog, win.err = gtk.DialogNew()
	dialog.AddButton("Add",gtk.RESPONSE_ACCEPT)



	taskEntry, win.err = gtk.EntryNew()

	dialogContent, win.err = dialog.GetContentArea()
	dialog.SetTitle("Add Task")
	dialogContent.Add(taskEntry)
	dialog.SetModal(true)
	taskEntry.Show()
	response := dialog.Run()
	if  response == int(gtk.RESPONSE_ACCEPT) {
		var content *gtk.Entry
		var delete *gtk.Button
		var rowBox *gtk.Box

		text,notext:= taskEntry.GetText()

		if notext != nil {
			dialog.Destroy()
		}
		delete, win.err = gtk.ButtonNewWithLabel("Delete")
		rowBox, win.err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 1)
		content, win.err = gtk.EntryNew()
		delete.SetRelief(gtk.RELIEF_NONE)
		rowBox.Add(content)
		rowBox.Add(delete)

		delete.Connect("clicked", func(){ win.DeleteTask(rowBox)} )
		rowBox.SetChildPacking(content, true, true, 0, gtk.PACK_START)
		rowBox.SetHomogeneous(false)
		content.SetText(text)
		log.Println(text)

		win.TaskList.Insert(rowBox, -1)
		win.Window.ShowAll()
		dialog.Destroy()
	}


}

func (win *MainWindow) DeleteTask (row *gtk.Box){
	row.Destroy()
}
