<p align="center">
    🌟FiveM Resource Loader - A fast way to install FiveM scripts from github  🌟  
</p>
<p align="center">
    🚨Work in progress developing an UI and enhancing the command-line tool🚨
</p>

## ✨ Features
* Download FiveM scripts from Github
* Places the script in the `[local]` folder
* Adds `ensure {script_name}` to `server.cfg`

## 🔭Usage
### ⚡ Command-line tool
* **help**
```
fivem-loader help
```
* **First time setup** 
(replace `C:\Users\fivem\server-data` with your actual path on server)
```
fivem-loader setup --server-data-path C:\Users\fivem\server-data
``` 
* **Installing a script**
This will download the **master** branch as a zip and extract them into `[local]`. Make sure the repository is set to public and is following the [supported structure](#Compatible-repositories)
```
fivem-loader load --github https://github.com/happsie/fivem-hello-world --script-name hello_world
```

## 🖖Compatible repositories
The supported structure for repositories using this tool will be updated to support different kind of repository structures

**only *one* FiveM script per repository is currently supported**

## 🚀Upcoming features
* Web GUI for installing script
* List of compatible scripts (easy click-install)
* Automatic update of scripts (setting per script)
* Making it more dynamic to match advanced servers

## 🆘Need help?
[Create an issue](https://github.com/happsie/fivem-loader/issues) 


