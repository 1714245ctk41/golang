package main

func main(){
	client := &client{}
	mac := &mac{}

	client.insertLightningConnectorIntoComputer(mac)

	windowMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowMachine : windowMachine
	}

	client.insertLightningConnectorIntoComputer(windowsMachineAdapter)

}