<a href="https://github.com/Irmine/GoMine">
    <img src="https://github.com/Irmine/GoMine/blob/master/GoMineBanner.jpg" width="600" height="200" align="left">
</a> <br> <br> <br> <br> <br> <br> <br> <br> <hr>

#### GoMine is a Minecraft Bedrock Edition server software written in Go.

### Information
GoMine is a fast multi-threaded Minecraft server software. It aims to provide a highly customizable API for plugin developers to use. GoMine aims to make the setup of a server very easy, (through an executable) with low compile times, and aims to make GoMine usable for other purposes than just a vanilla server.

### Current State
GoMine is currently under heavy development and is not usable for production servers yet. It lacks many features which are yet to be implemented, and has (yet unknown) bugs that should be resolved.

### Releases and Development Builds
Development builds of GoMine might be unstable and should be used with care. It is always recommended to run officially released versions of GoMine for production where possible, to ensure no nasty bugs appear. If you do decide to run a development version, be aware that bugs may occur. Don't hesitate to report those bugs.

### Setup
GoMine aims to make the setup of a server very easily. The setup of GoMine can be explained in a couple steps.
If you want to use an official release:
1. Download the executable for your operating system from `Releases` and move it to your setup directory.
2. Execute the executable to run the server.

If you would like to use a development version:
1. Install Go > 1.9 from the official release page.
2. To clone the repository, execute `go get github.com/irmine/gomine`.
3. Compile GoMine by navigating into the `irmine/gomine` folder and executing `go install`.
4. Navigate to the folder at `GOBIN`, and grab the executable.
5. Move it to your setup folder and execute the executable.

### Issues
Issues can be reported in the `Issues` tab. Please provide enough information for us to solve the problem. The more information you provide, the easier it makes it for us to fix your issue.

### License
GoMine is licensed under the GNU General Public License.
