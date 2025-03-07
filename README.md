# File Metadata Service project

## Overview:
Create a gRPC service that provides metadata about files stored in a system. The service should allow clients to query information about a file, such as its size, type, last modified date, and whether it exists.

## Operations:
* GetFileMetadata - A client can send the file path to the server, and the server will respond with metadata about the file, such as:
    * File size (in bytes)
    * File type (e.g., "text", "image", "pdf")
    * Last modified date
* CheckFileExists - A client can send a file path to the server, and the server will simply return whether the file exists on the system.
    * Whether the file exists

## Requirements:
**Proto File:**

### Define a FileService service.
Define messages for requesting file metadata (including file path), and responses containing the metadata.
#### Server:

[server](./cmd/)

Implement the server in Go, which will:
* Accept file paths from clients.
* Use the Go os and time libraries to get the file size, type, and last modified time.
* Return the results in a gRPC response message.

#### Client:

[client](./client/)

* Implement the client in Go to interact with the service:
* Send a file path to the service and get back the metadata.
* Optionally, handle errors (e.g., file not found, invalid file path).
* Extensions (Optional for extra challenge):
    * Implement file type detection based on file content (e.g., use a library to detect MIME types or extension-based checks).
    * Add logging or error handling to track when file paths are invalid or missing.
    * Add support for multiple files in a single request (e.g., the client could send an array of file paths to get metadata for multiple files at once).
