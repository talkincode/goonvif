# Goonvif
Easy management of IP devices, including cameras. Goonvif is an ONVIF implementation for managing IP devices. The purpose of this library is to manage IP cameras and other devices that support the ONVIF standard easily and conveniently.

## Installation
To install the library, use the go get utility:
```
go get github.com/talkincode/goonvif
```
## Supported services
The following services are fully implemented:
- Device
- Media
- PTZ
- Imaging

## Usage

### General concept
1) Connection to the device
2) Authentication (if required)
3) Definition of data types
4) Execution of the required method

#### Connect to device
If there is a device at *192.168.13.42* in the network, and its ONVIF services use port *1234*, then you can connect to the device using the following method:
```
dev, err := goonvif.NewDevice("192.168.13.42:1234")
```

*ONVIF port may be different depending on the device and to find out which port to use, you can go to the web interface of the device. **Usually it's port 80.

#### Authentication
If any function of one of the ONVIF services requires authentication, the `Authenticate` method must be used.
```
device := onvif.NewDevice("192.168.13.42:1234")
device.Authenticate("username", "password")
```

#### Data Type Definition
Each ONVIF service in this library has its own package in which all data types of this service are defined, and the name of the package is identical to the service name and starts with a capital letter.
Goonvif has defined structures for each function of each ONVIF service supported by this library.
Let's define the data type of the `GetCapabilities` function of the `Device` service. This is done as follows:
```
capabilities := Device.GetCapabilities{Category: "All"}
```

Why does the GetCapabilities structure have a Category field and why is the value of this field All?

The figure below shows the documentation of the [GetCapabilities](https://www.onvif.org/ver10/device/wsdl/devicemgmt.wsdl) function. You can see that the function takes one parameter Category and its value should be one of the following:  `'All', 'Analytics', 'Device', 'Events', 'Imaging', 'Media' or 'PTZ'`.

![Device GetCapabilities](img/exmp_GetCapabilities.png)

An example of the GetServiceCapabilities  function data type definition [PTZ](https://www.onvif.org/ver20/ptz/wsdl/ptz.wsdl):
```
ptzCapabilities := PTZ.GetServiceCapabilities{}
```
You can see in the figure below that GetServiceCapabilities takes no arguments.

![PTZ GetServiceCapabilities](img/GetServiceCapabilities.png)

*General data types are in the xsd/onvif package. Data types (structures) that can be common to all services are defined in the onvif package.

An example of the data type of the CreateUsers function of the [Device](https://www.onvif.org/ver10/device/wsdl/devicemgmt.wsdl) service:
```
CreateUsers := Device.CreateUsers{User: onvif.User{Username: "admin", Password: "qwerty", UserLevel: "User"}}
```

You can see from the figure below that in this example the CreateUsers structure field should be User, the data type of which is a User structure containing Username, Password, UserLevel fields and an optional Extension. The User structure is located in the onvif package.

![Device CreateUsers](img/exmp_CreateUsers.png)

#### Executing the required method
To execute any function of one of the ONVIF services whose structure has been defined, the ``CallMethod'' of the device object must be used.
```
createUsers := Device.CreateUsers{User: onvif.User{Username: "admin", Password: "qwerty", UserLevel: "User"}
device := = onvif.NewDevice("192.168.13.42:1234")
device.Authenticate("username", "password")
resp, err := dev.CallMethod(createUsers)
```
