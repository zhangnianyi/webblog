package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	//"github.com/rifflock/lfshook"
	"math"
	"os"
	"time"
)

func Logger()gin.HandlerFunc{
	//把日志写进文件
	filePath :="log/log"
	LinkName :="latest_log.log"
	src,err :=os.OpenFile(filePath,os.O_RDWR|os.O_CREATE,0755)
	if err !=nil{

		fmt.Println(err)
	}
	logger :=logrus.New()
	logger.Out =src
	logger.SetLevel(logrus.DebugLevel)
	logWriter,_ :=retalog.New(
		filePath+"%Y%m%d.log",
		retalog.WithMaxAge(7*24*time.Hour),  //日志保留七天
		retalog.WithRotationTime(24*time.Hour),  //日期24小时切割一次
		retalog.WithLinkName(LinkName),
		)
	Writrmap :=lfshook.WriterMap{
		logrus.InfoLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook :=lfshook.NewHook(Writrmap,&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.AddHook(Hook)

	//下面是你日志里需要收集的东西
	return func(c *gin.Context) {
		starttime :=time.Now()
		c.Next()
		stoptime :=time.Since(starttime)
		//speendTIme :=int(math.Ceil(float64(stoptime.Nanoseconds())/1000000.0))  //得到程序运行的时候 并且四舍五入
		//得到程序运行的时候 并且四舍五入
		speendTIme :=fmt.Sprintf("%d ms",int(math.Ceil(float64(stoptime.Nanoseconds())/1000000.0)))
		//得到程序运行的主机
		Hostname,err:=os.Hostname()
		if err !=nil{
			Hostname="UnKnow"
		}
		////得到程序运行的请求码
		statusCode :=c.Writer.Status()
		//得到 客户端的ipclientIP
		clientIP :=c.ClientIP()
		//得到客户端请求的类型
		userAgent :=c.Request.UserAgent()
		//得到请求长度
		dataSize :=c.Writer.Size()
		if dataSize <0{
			dataSize = 0
		}
		//得到请求方法
		method :=c.Request.Method
		//得到请求的url路径
		path :=c.Request.RequestURI

		//把上面你要获取的东西塞进你的logger里面
		entry :=logger.WithFields(logrus.Fields{
			"Hostname" :Hostname,
			"Status":statusCode,
			"Speedtime":speendTIme,
			"IP":clientIP,
			"method":method,
			"Datasize":dataSize,
			"Useragent":userAgent,
			"path":path,
		})
		if len(c.Errors) >0{
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >=500{
			entry.Error()
		}else if statusCode >=400{
			entry.Warn()
		}else {
			entry.Info()
		}

	}

}