package main

import(
	"github.com/conformal/gotk3/gtk"
	"tasks/ui"
 )
 
 func main(){
 	gtk.Init(nil)
 	
 	
 	window := ui.MainWindowNew(gtk.WINDOW_TOPLEVEL)

 	window.Window.ShowAll()
 
 	
 
 	
 	gtk.Main()
 
 }
