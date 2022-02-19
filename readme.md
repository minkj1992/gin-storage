# Full stack upload files with gcs

> This is a full stack project focused on serving static files

![](static/1.png)
![](static/2.png)
![](static/3.png)
![](static/4.png)
![](static/5.png)

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

> [refs](https://github.com/satansdeer/ultimate-react-hook-form-form.git)

```bash
npm install
npm run start
```

### backend

```bash
go get -u github.com/gin-gonic/gin
go get -u cloud.google.com/go/storage
go run main.go
```
