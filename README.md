# DataDog Object Manager (DDOM)

**DDOM** is a simple tool built by SRE team for easelly manage DataDog dashboards, monitors or screenboards. 
## Configuration
First of all you need to specify in **ddom.yaml** (ddom.yaml should be placed in the same directory whith the ddom binary file) a list of DataDog organisations. The format for this file is:

```yaml
---
organisation1:
  APIKey: api-key-here
  AppKey: application-key-here
  URL: custom-url-here
organisation2:
  APIKey: api-key-here
  AppKey: application-key-here
  URL: custom-url-here
```

## Methods:
To see a list with all the methods, just run **help** in the DDOM's interactive console.
### checkOrg
Is used to check the configuration for a specific organisation.
### setObject
Is used to set a part of the app context. It will let you choose what kind of object want to manage.
### setMethod
Will let you chose what kind of operation want to apply to the selected objects.
**Available methods:**
* ***Backup*** - it will ask for a organisation that will be used as data source. After that will list you all the objects (monitors or dashboards or screenbords) considering what you choosed on the **setObject** step. After you select the IDS of the objects (a comma separated list) and a pathon the disk where to store the files, the app will download and save for you the JSON serialized form of the selected objects.
* ***Transfer*** - this method will ask you for a source organisation name and a destination organisation name. Next, it will list for you the objects considering your choice on the **setObject** step. After you provide the comma separated list of IDS it will try to download the selected objects from the source organisation and to push it to the destination organisation.
* ***LoadFromFile*** - This method will ask for a destination organisation name and a path where it can find some JSON files. It will try to read JSON files contents and push it into the destination organisation.

### run
This is the trigger for the selected method. ***setObject*** and ***setMethod*** will ensure the context for the run but this method will trigger the execution.

## Usage example:
```bash
>>> setObject
Use 1 for Monitors or 2 for Dashboards or 3 for Timeboards 1
Context set to Monitors
>>> setMethod
1 for Backup
2 for Transfer
3 for LoadFromFile
>>> 1
Method set to Backup
>>> run
Enter a DataDog organisation: mcc1
Enter a path where to store: ./bkp
ID          Name
7252418     Some monitor name
7418288     Another monitor name
7445084     Yet Another Monitor Name

Enter comma separated ids to be stored: 7252418,7418288,7445084
>>> exit
```

