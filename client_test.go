package marionette_client

import (
	"fmt"
	"testing"
)

var client *Client

func init() {
	client = NewClient()
	client.Transport(&MarionetteTransport{})
}

func TestNewSession(t *testing.T) {
	err := client.Connect("", 0)
	if err != nil {
		t.Error(err)
	}

	r, err := client.NewSession("", nil)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(r.Value)
	//client.Close()
}

// working
func TestGetSessionID(t *testing.T) {
	if client.SessionId != client.GetSessionID() {
		fmt.Println("SessionId differ...")
		t.FailNow()
	}

	fmt.Println("session is : ", client.SessionId)
}

func TestGetPage(t *testing.T) {
	r, err := client.Get("http://www.abola.pt/")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r.Value)
}

//func TestConnectWithActiveConnection(t *testing.T) {
//	err := client.Connect("", 0)
//	if err == nil {
//		fmt.Println(err)
//	}
//
//	fmt.Println("No Error..")
//}

// working
func TestGetSessionCapabilities(t *testing.T) {
	r, err := client.GetSessionCapabilities()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r.BrowserName)
}


func TestScreenshot(t *testing.T) {
	r, err := client.Screenshot()
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	println(r.MessageID)
	println(r.DriverError)
	println(r.Size)
	//this print ise a problem for travis builds, since it can surpass the 4 MB of log size.
	// don't print the base64 encoded image.
	//println(r.Value)
}


// working
func TestLog(t *testing.T) {
	r, err := client.Log("message testing", "warning")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r.Value)
}

// working
func TestGetLogs(t *testing.T) {
	r, err := client.GetLogs()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r.Value)
}

func TestSetContext(t *testing.T) {
	r, err := client.SetContext(CONTEXT_CHROME)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r.Value)

	r, err = client.SetContext(CONTEXT_CONTENT)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r.Value)
}

func TestGetContext(t *testing.T) {
	r, err := client.GetContext()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r.Value)
}

func TestSetScriptTimout(t *testing.T) {
	r, err := client.SetScriptTimeout(1000)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r.Value)
}

func TestSetPageTimout(t *testing.T) {
	r, err := client.SetPageTimeout(1000)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r.Value)
}

func TestSetSearchTimout(t *testing.T) {
	r, err := client.SetSearchTimeout(1000)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r.Value)
}

func TestExecuteScript(t *testing.T) {
	script := "function testMyGoMarionetteClient() { alert('yes'); } testMyGoMarionetteClient();"
	args := []interface{}{}
	r, err := client.ExecuteScript(script, args, 1000, false)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r.Value)

	client.DismissDialog()
}

func TestExecuteScriptWithArgs(t *testing.T) {
	script := "function testMyGoMarionetteClientArgs(a, b) { alert(a + b); }; testMyGoMarionetteClientArgs(arguments[0], arguments[1]);"
	args := []interface{}{1, 3}
	r, err := client.ExecuteScript(script, args, 1000, false)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r.Value)

	client.DismissDialog()
}

func TestGetTitle(t *testing.T) {
	title, err := client.GetTitle()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(title)

}
func TestFindElement(t *testing.T) {
	element, err := client.FindElement("id", "clubes-hp", nil)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(element.Id())
	fmt.Println(element.IsEnabled())
	fmt.Println(element.IsSelected())
	fmt.Println(element.IsDisplayed())
	fmt.Println(element.TagName())
	fmt.Println(element.Text())
	fmt.Println(element.Attribute("id"))
	fmt.Println(element.CssValue("text-decoration"))
	fmt.Println(element.Rect())

	collection, err := element.FindElements("css selector", "li")
	if 18 != len(collection) {
		t.FailNow()
	}
}

func TestSendKeys(t *testing.T) {
	e, err := client.FindElement("id", "topo_txtPesquisa", nil)
	if err != nil {
		fmt.Println(err)
	}

	e.SendKeys("teste")
}

func TestFindElements(t *testing.T) {
	elements, err := client.FindElements("css selector", "li", nil)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(len(elements))
}

func TestCurrentWindowHandle(t *testing.T) {
	r, err := client.GetCurrentWindowHandle()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r.Value)
}

func TestCurrentChromeWindowHandle(t *testing.T) {
	r, err := client.GetCurrentChromeWindowHandle()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r.Value)
}

func TestWindowHandles(t *testing.T) {
	r, err := client.WindowHandles()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(r)
}

// working
//func TestQuitApplication(t *testing.T) {
//	r, err := client.QuitApplication()
//	if err != nil {
//		fmt.Println(err)
//		t.FailNow()
//	}
//
//	fmt.Println(r.ResponseError)
//}
