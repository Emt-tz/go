package main

import (
	"fmt"
	"github.com/andlabs/ui"
)

func main() {
	login()
}

//page 125 at now

func createuser() {
	err := ui.Main(func(){
		window := ui.NewWindow("New user",300,300,false)
		window.SetMargined(true)
		window.OnClosing(func(*ui.Window) bool{
			ui.Quit()
			return true
		})


		//Entry Fields
		username := ui.NewEntry()
		password := ui.NewPasswordEntry()
		pwdconfirm := ui.NewPasswordEntry()

		//Buttons 
		Createuser := ui.NewButton("Create")
		Createuser.OnClicked(func(*ui.Button){
			fmt.Println("User Created")
		})

		Backtologin := ui.NewButton("Login")
		Backtologin.OnClicked(func(*ui.Button){
			ui.Quit()
		})


		//Set up our box
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel(""),false)
		box.Append(username,false)
		box.Append(ui.NewLabel(""),false)
		box.Append(password,false)
		box.Append(ui.NewLabel(""),false)
		box.Append(pwdconfirm,false)
		box.Append(ui.NewLabel(""),false)
		box.Append(Createuser,false)
		box.Append(ui.NewLabel(""),false)
		box.Append(Backtologin,false)

		//display

		window.SetChild(box)
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

func login() {
	err := ui.Main(func(){
		window := ui.NewWindow("login",300,300,false)
		window.SetMargined(true)
		window.OnClosing(func(*ui.Window) bool{
			ui.Quit()
			return true
		})

		//display results label
		res := ui.NewLabel("")
		res.SetText("")
		//our top title is place here
		toptitle := ui.NewLabel("Login to Continue")
		//ui.NewLabel("") empty space for spacing

		//Entru box to get user values
		username := ui.NewEntry()
		username.SetText("Enter username")


		password := ui.NewPasswordEntry()
		password.SetText("Enter Password")
		//the login button
		loginbutton := ui.NewButton("login")
		loginbutton.OnClicked(func(*ui.Button){
			un := username.Text()
			//pwd := password.Text()

			res.SetText(un)
		})

		//the Createuserbutton

		newuserbutton := ui.NewButton("Create User")
		newuserbutton.OnClicked(func(*ui.Button) {
			ui.Quit()

			createuser()
		})

		//this is the first box
		box := ui.NewVerticalBox()
		box.Append(ui.NewVerticalSeparator(),false)
		box.Append(toptitle, false)
		box.Append(ui.NewLabel(""),false)
		box.Append(username,false)
		box.Append(ui.NewLabel(""),false)
		box.Append(password,false)
		box.Append(ui.NewLabel(""),false)
		box.Append(loginbutton, false)
		box.Append(ui.NewLabel(""),false)
		box.Append(newuserbutton,false)
		box.Append(ui.NewLabel(""),false)
		box.Append(ui.NewLabel(""),false)
		box.Append(res,false)
	


		window.SetChild(box)
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

