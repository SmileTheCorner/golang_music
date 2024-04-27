package conf

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
	"go_study/src/enum"
	"go_study/src/global"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"
)

var (
	ctx      = context.Background()
	location = viper.GetString("upload.minio.location")
)

// InitMinio 初始化minio连接
func InitMinio() {
	endpoint := viper.GetString("upload.minio.endpoint")
	accessKeyID := viper.GetString("upload.minio.accessKeyId")
	secretAccessKey := viper.GetString("upload.minio.secretAccessKey")
	useSSL := viper.GetBool("upload.minio.secure")
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		global.Logger.Error(fmt.Sprintf("Connect Minio Fail : %s", err.Error()))
		panic(fmt.Sprintf("Connect Minio Fail : %s", err.Error()))
	}
	global.MinioClient = client
}

// BucketNameExists 判断桶存不存在
func BucketNameExists(bucketName string) (bool, error) {
	exists, errBucketExists := global.MinioClient.BucketExists(ctx, bucketName)
	if errBucketExists == nil && exists {
		global.Logger.Info(fmt.Sprintf("Bucket  Does Exist  : %s", bucketName))
		return exists, errBucketExists
	} else {
		global.Logger.Error(fmt.Sprintf("Cat Bucket Fail : %s", errBucketExists.Error()))
		return exists, errBucketExists
	}
}

// CreateBucket 创建桶
func CreateBucket(bucketName string) (bool, error) {
	err := global.MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		return false, err
	}
	return true, err
}

// RemoveObject 删除对象
func RemoveObject(bucketName, objectName string) error {
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	err := global.MinioClient.RemoveObject(ctx, bucketName, objectName, opts)
	if err != nil {
		return err
	}
	return nil
}

// UploadFile 上传文件
func UploadFile(file *multipart.FileHeader, bucketName string, classifyName string) (string, error) {
	//文件上传前判断当前桶存不存在
	exists, _ := BucketNameExists(bucketName)
	if !exists {
		//当前桶不存在则创建桶
		_, err := CreateBucket(bucketName)
		if err != nil {
			return "", err
		}
	}

	var contentType string
	fileName := file.Filename
	//获取文件的扩展名
	ext := filepath.Ext(fileName)
	if ext != "" {
		//获取文件后缀名
		lower := strings.ToLower(ext[1:])
		contentType = "application/" + lower
	} else {
		contentType = "application/octet-stream"
	}
	fileSize := file.Size
	fileAllName := "/" + classifyName + "/" + time.Now().Format("2006-01-02 15:03:04") + " " + fileName

	fileInfo, err := file.Open()
	defer fileInfo.Close()
	if err != nil {
		return "上传文件有误", err
	}

	uploadInfo, uploadErr := global.MinioClient.PutObject(
		ctx, bucketName, fileAllName, fileInfo, fileSize,
		minio.PutObjectOptions{ContentType: contentType})
	if uploadErr != nil {
		return "文件上传错误", uploadErr
	}

	url := enum.Url + viper.GetString("upload.minio.endpoint") + "/" + bucketName + fileAllName
	fmt.Println("上传文件的信息", uploadInfo)
	return url, uploadErr
}
