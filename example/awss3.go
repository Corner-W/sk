package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
)

func main() {
	bucket := aws.String("")        //bucket名称
	key := aws.String("testobject") //object keyname
	access_key := ""
	secret_key := ""
	end_point := "http://s3.nie.netease.com"       //endpoint设置，不要动
	myContentType := aws.String("application/zip") //content-type设置
	myACL := aws.String("public-read")             //acl 设置
	metadata_key := "udf-metadata"                 //自定义Metadata key
	metadata_value := "abc"                        //自定义Metadata value
	myMetadata := map[string]*string{
		metadata_key: &metadata_value,
	}
	// Configure to use S3 Server
	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(access_key, secret_key, ""),
		Endpoint:         aws.String(end_point),
		Region:           aws.String("us-east-1"),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(false), //virtual-host style方式，不要修改
	}
	newSession := session.New(s3Config)
	s3Client := s3.New(newSession)
	cparams := &s3.HeadBucketInput{
		Bucket: bucket, // Required
	}
	_, err := s3Client.HeadBucket(cparams)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	uploader := s3manager.NewUploader(newSession)
	filename := "/tmp/demo.pdf" //上传文件路径
	f, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("failed to open file %q, %v", filename, err)
		return
	}
	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      bucket,
		Key:         key,
		Body:        f,
		ContentType: myContentType,
		ACL:         myACL,
		Metadata:    myMetadata,
	}, func(u *s3manager.Uploader) {
		u.PartSize = 10 * 1024 * 1024 // 分块大小,当文件体积超过10M开始进行分块上传
		u.LeavePartsOnError = true
		u.Concurrency = 3
	}) //并发数
	if err != nil {
		fmt.Printf("Failed to upload data to %s/%s, %s\n", *bucket, *key, err.Error())
		return
	}
	fmt.Printf("file uploaded to, %s\n", result.Location)
	down_file := "/tmp/down_file.pdf" //下载路径
	file, err := os.Create(down_file)
	if err != nil {
		fmt.Println("Failed to create file", err)
		return
	}
	defer file.Close()
	downloader := s3manager.NewDownloader(newSession)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: bucket,
			Key:    key,
		})
	if err != nil {
		fmt.Println("Failed to download file", err)
		return
	}
	fmt.Println("Downloaded file", file.Name(), numBytes, "bytes")
}
