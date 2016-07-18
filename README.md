# Wand

Wand is a simple tool for serving a single file over HTTP.

## Usage

Pass a file of any type to <code>wand</code> to start an HTTP server with one route, `/`. Wand will respond to `GET` requests made to this endpoint with the contents of the file. The following command line arguments may also be provided: 

<table>
    <tr>
        <td><code>-p, -port</code></td>
        <td>Set the port to bind Wand to. Default: <code>11002</code></td>
    </tr>
    <tr>
        <td><code>-s, -script</code></td>
        <td>Execute the file when it is requested. Default: <code>false</code></td>
    </tr>
</table>

## Script Mode
When `-s` is passed, the file will be run as a shell script and Wand will produce a JSON response containing the output of the script as `value`.

## License
[Apache License 2.0](LICENSE.md)
