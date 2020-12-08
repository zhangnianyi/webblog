package modles

import (
	"context"
	"ginblog/utils"
	"ginblog/utils/errormessage"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"mime/multipart"
)

var Access  =utils.AccessKey
var ScretKey = utils.SecretKey
var Bucket =utils.Bucket
var Imgurl=utils.QiniuServer

func UploadFile(file multipart.File,fileszie int64)(string,int){
	putpolicy :=storage.PutPolicy{
		Scope: Bucket,
	}
	mac :=qbox.NewMac(Access,ScretKey)
	upToken :=putpolicy.UploadToken(mac)
	cfg :=storage.Config{
		Zone: &storage.ZoneHuanan,
		UseCdnDomains: false,
		UseHTTPS: false,
	}
	putExtra :=storage.PutExtra{}
	fromUploader :=storage.NewFormUploader(&cfg)
	ret :=storage.PutRet{}
	err :=fromUploader.PutWithoutKey(context.Background(),&ret,upToken,file,fileszie,&putExtra)
	if err !=nil{
		return  "",errormessage.ERROR
	}
	url :=Imgurl +ret.Key
	return url,errormessage.SUCCESS

}