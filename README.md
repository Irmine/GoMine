<a href="https://github.com/Irmine/GoMine">
    <img src="https://github.com/Irmine/GoMine/blob/master/GoMineBanner.jpg" width="600" height="200" align="left">
</a> <br> <br> <br> <br> <br> <br> <br> <br> <hr>

#### GoMine is a Minecraft Bedrock Edition server software written in Go.

### Information
GoMine is a fast multi-threaded server software. It aims to provide a highly customizable API for plugin developers to use. GoMine aims to make the setup of a server very easy, (through an executable) with low compile times.

### Current State
GoMine is currently under heavy development and is not usable for production servers yet. It lacks many features which are yet to be implemented, and has (yet unknown) bugs that should be resolved.

### Releases and Development Builds
Development builds of GoMine might be unstable and should be used with care. It is always recommended to run officially released versions of GoMine for production, to ensure no nasty bugs appear.

### Setup
GoMine aims to make the setup of a server very easily. The setup of GoMine can be explained in a couple steps.
If you want to use an official release:
1. Install Go 1.9.2 from the official release page.
2. Download the executable for your operating system from `Releases` and move it to your setup directory.
3. Execute the executable to run the server.

If you want to use a development version:
1. Install Go 1.9.2 from the official release page.
2. Use: `git clone --recursive https://github.com/Irmine/GoMine` to download the full repository.
3. Navigate into the GoMine/src folder and execute: `go build`. It will now start compiling GoMine.
4. An executable should have appeared. Execute it to run the server.

### Issues
Issues can be reported in the `Issues` tab. Please provide enough information for us to solve your problem. We're no magicians, we can't read what's in your mind. The more information you give us, the bigger the chance your issue gets fixed quickly so we can both move on.

### License
GoMine is licensed under the GNU General Public License.
