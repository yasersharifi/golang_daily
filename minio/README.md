# MinIO Go 

A **single-file Go program** to interact with **MinIO** (S3-compatible storage).

---

## Features

- Connect to MinIO  
- Create bucket  
- Upload image & video  
- Generate presigned download URL  
- List & delete objects  
- Optional: public-read policy  

---

## Project

```
minio-demo/
├── main.go
├── sample.jpg
├── sample.mp4
├── go.mod
└── README.md
```

---

## Run MinIO

```bash
docker run -d -p 9000:9000 -p 9001:9001 \
  -e "MINIO_ROOT_USER=minioadmin" \
  -e "MINIO_ROOT_PASSWORD=minioadmin" \
  quay.io/minio/minio server /data --console-address ":9001"
```

Console: http://localhost:9001

---

## Setup & Run

```bash
mkdir -p ~/go/minio-demo && cd ~/go/minio-demo

# Save main.go (from code)
go mod init minio-demo
go get github.com/minio/minio-go/v7

# Sample files
convert -size 100x100 xc:blue sample.jpg
ffmpeg -f lavfi -i testsrc=duration=1:size=320x240:rate=30 -c:v libx264 sample.mp4

go run main.go
```

---

## Config (`main.go`)

| Variable     | Default           |
|--------------|-------------------|
| `endpoint`   | `localhost:9000`  |
| `accessKey`  | `minioadmin`      |
| `secretKey`  | `minioadmin`      |
| `bucketName` | `demo-bucket`     |

> Use environment variables in production.

---

## Output

```
Connected to MinIO
Bucket "demo-bucket" created
Uploaded photos/sample.jpg ...
Download link: http://localhost:9000/demo-bucket/photos/sample.jpg?...

Objects (before delete):
  • photos/sample.jpg
  • videos/sample.mp4

Deleted object: photos/sample.jpg

Objects (after delete):
  • videos/sample.mp4
```

---

## Customize

- Use real files  
- Enable SSL: `useSSL = true`  
- Add versioning: `SetBucketVersioning()`  
- Bulk delete: `RemoveObjects()`  

---

## Production

- Use context timeouts  
- Retry on errors  
- Secure credentials  
- Enable bucket versioning  

---

## What’s Next?

Here’s the **natural progression** for this project:

| Next Step | Description |
|---------|-----------|
| **1. Add Environment Variables** | Replace hardcoded credentials with `os.Getenv()` |
| **2. Enable Bucket Versioning** | `client.SetBucketVersioning()` + delete markers |
| **3. Bulk Upload/Delete** | Use `ListObjects` + `RemoveObjects` channel |
| **4. Add Lifecycle Rules** | Auto-delete old objects via `SetBucketLifecycle` |
| **5. HTTPS & Custom Domain** | Use `useSSL = true` + real certs |
| **6. CLI Tool** | Turn into `minioctl` with flags |
| **7. Dockerize the App** | `Dockerfile` + `docker-compose.yml` |

Want me to generate **any of these next steps**? Just say the number or name.