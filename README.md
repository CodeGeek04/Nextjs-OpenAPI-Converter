# Nextjs-OpenAPI-Converter

A command-line tool that automatically converts your Next.js App Router API routes into a Postman Collection. The generated collection can be imported into Postman, Insomnia, Hoppscotch, and other API testing tools.

## Features

- 🗺️ Automatically scans Next.js App Router API routes (`app/api/*`)
- 📝 Generates a Postman Collection file
- 🔄 Preserves your API route hierarchy
- 📁 Handles nested routes and folders
- 🚀 Direct import into Postman and similar tools

## Requirements

- Next.js project using App Router (not compatible with Pages Router)
- Go installed on your machine
- API routes defined in the `app/api` directory

## Installation

1. Clone the repository:
```bash
git clone https://github.com/CodeGeek04/Nextjs-OpenAPI-Converter.git
cd Nextjs-OpenAPI-Converter
```

2. Install dependencies:
```bash
go mod tidy
```

## Usage

1. Open `main.go` and modify the following paths in the main function:
```go
item := handler.HandleFolder("NextJS API Path", "NextJS Root Path")
```
Replace:
- `"NextJS API Path"` with the path to your Next.js API directory (e.g., `/app/api`)
- `"NextJS Root Path"` with the path to your Next.js project root directory

Example:
```go
item := handler.HandleFolder("./my-nextjs-project/app/api", "./my-nextjs-project")
```

2. Run the converter:
```bash
go run main.go
```

3. The tool will generate a `postman_collection.json` file in the current directory.

4. Import the generated file into Postman:
   - Open Postman
   - Click "Import"
   - Choose the generated `postman_collection.json` file

## Current Limitations

- Only works with Next.js App Router (not compatible with Pages Router)
- Request names are currently set to "route.ts"
- Host is hardcoded to "localhost"
- Port is hardcoded to 3000
- Basic request body structure

## Upcoming Features

- [ ] Pages Router support
- [ ] Environment variables for host and port
- [ ] Dynamic route handling
- [ ] Sample request body generation
- [ ] Custom request naming
- [ ] Better route metadata parsing
- [ ] Support for API authentication
- [ ] Dynamic route parameters
- [ ] Request body examples
- [ ] Response examples
- [ ] OpenAPI/Swagger output option

## Contributing

Contributions are welcome! Feel free to:
1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Next.js team for their excellent framework and App Router
- Postman for their API platform and collection format
- All contributors and users of this tool

## Support

If you encounter any issues or have questions, please:
1. Check existing GitHub issues
2. Create a new issue if needed
3. Provide detailed information about:
   - Your Next.js project structure
   - The error message (if any)
   - Your App Router API route structure
