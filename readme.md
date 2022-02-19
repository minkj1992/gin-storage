# Full stack upload files with gcs

> This is a full stack project focused on serving static files

## features

### client side

- post an article web page
  - upload images
  - add text form
  - upload audio(mp4) file
- retrieve an article
  - text
  - images
  - audio file

### backend side

- upload files to gcs
- delete files from gcs
- streaming audio(HLS / HTML5) file on front side

## setup

### frontend

> [refs](https://www.bezkoder.com/react-file-upload-axios/)

```bash
npm uninstall -g create-react-app
npm add create-react-app
npx create-react-app react-file-upload
npm install bootstrap axios

```

### backend

```bash
go get -u github.com/gin-gonic/gin
go get -u cloud.google.com/go/storage
go run main.go
```
