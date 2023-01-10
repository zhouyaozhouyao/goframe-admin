package upload

import (
	"context"

	"github.com/gogf/gf/v2/os/gfile"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"

	"github.com/gogf/gf/v2/frame/g"
)

type ossUpload struct {
}

var Oss = ossUpload{}

func (o *ossUpload) Upload(ctx context.Context, path string, filename string) (url string, err error) {
	aliCfg := g.Cfg().MustGet(ctx, "ali").MapStrStr()
	endpoint := "https://oss-cn-hangzhou.aliyuncs.com"
	accessKeyId := aliCfg["accessKey"]
	accessKeySecret := aliCfg["secret"]
	bucketName := "ybc-go"
	objectName := filename
	localFileName := gfile.Join(path, filename)
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return url, err
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		return url, err
	}

	err = bucket.SetObjectACL(filename, oss.ACLPublicReadWrite)
	if err != nil {
		return url, err
	}
	_, err = bucket.GetObjectACL(filename)
	if err != nil {
		return url, err
	}
	return "https://ybc-go.oss-cn-hangzhou.aliyuncs.com/" + filename, nil
}
