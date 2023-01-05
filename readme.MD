<p align="center">
    🌟FiveM Resource Loader - A fast way to install FiveM scripts from Github  🌟  
</p>
<p align="center">
    🚨Work in progress developing an UI and enhancing the command-line tool🚨
</p>

## 🖥️ Download
[Go to releases](https://github.com/happsie/fivem-loader/releases)

## ✨ Features
* Download FiveM scripts from Github
* Places the script in selected resource folder
* Adds `ensure {script_name}` to `server.cfg`

## 🔭Usage
### ⚡ Command-line commands
* **help**
```
help
```
* **First time setup**  
(replace `C:\Users\fivem\server-data` with your actual path on server)
```
setup --server-data-path C:\Users\fivem\server-data
``` 
* **Load a new script**  
This will download the **master** branch as a zip and extract them into selected resource destination. Make sure the repository is set to public and is following the [supported project structure](#Compatible-repositories)  
Use `--skip-cfg` to skip making changes to your `server.cfg`
```
load --github https://github.com/happsie/fivem-hello-world --script-name hello_world
```
* **See loaded scripts by FiveM-loader**  
```
loaded
```
* **Unload script**  
Replace `hello_world` with your script name. You can see script names by using the command `loaded`.  
**Important**, fivem-loader will at this version not automatically remove the config for the project
```
unload --script-name hello_world
```
* **Update single script installed by FiveM Loader**
```
update --script-name hello_world
```
* **Update all installed scripts by FiveM Loader**
```
update
```

## 🖖Compatible repositories
The supported structure for repositories using this tool will be updated to support different kind of repository structures

**only one FiveM-script per repository is currently supported to 100%**

## 🚀Upcoming features
* Web GUI for installing script
* List of compatible scripts (easy click-install)
* Automatic update of scripts (setting per script)
* Allow private repositories
* Custom addition to `server.cfg`

## 🆘Need help?
[Create an issue](https://github.com/happsie/fivem-loader/issues) 


