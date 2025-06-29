package zaplogger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var Logger *zap.Logger
var SugaredLogger *zap.SugaredLogger

func init() {
	InitLogger()
}

func InitLogger() {
	//获取编码器
	encoder := getEncoder()

	//日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})

	//info文件WriteSyncer
	infoFileWriteSyncer := getInfoWriterSyncer()
	//error文件WriteSyncer
	errorFileWriteSyncer := getErrorWriterSyncer()

	//生成core
	//multiWriteSyncer := zapcore.NewMultiWriteSyncer(writerSyncer, zapcore.AddSync(os.Stdout)) //AddSync将io.Writer转换成WriteSyncer的类型
	//同时输出到控制台 和 指定的日志文件中
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority)
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority)

	//将infocore 和 errcore 加入core切片
	var coreArr []zapcore.Core
	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)

	//生成Logger
	Logger = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller()) //zap.AddCaller() 显示文件名 和 行号
	SugaredLogger = Logger.Sugar()
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func levelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	var level string
	switch l {
	case zapcore.DebugLevel:
		level = "[DEBUG]"
	case zapcore.InfoLevel:
		level = "[INFO]"
	case zapcore.WarnLevel:
		level = "[WARN]"
	case zapcore.ErrorLevel:
		level = "[ERROR]"
	case zapcore.DPanicLevel:
		level = "[DPANIC]"
	case zapcore.PanicLevel:
		level = "[PANIC]"
	case zapcore.FatalLevel:
		level = "[FATAL]"
	default:
		level = fmt.Sprintf("[LEVEL(%d)]", l)
	}
	enc.AppendString(level)
}

func shortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(fmt.Sprintf("[%s]", caller.TrimmedPath()))
}

func NewEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    levelEncoder, //zapcore.CapitalLevelEncoder,
		EncodeTime:     timeEncoder,  //指定时间格式
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   shortCallerEncoder, //zapcore.ShortCallerEncoder,
	}
}

// core 三个参数之  Encoder 获取编码器
func getEncoder() zapcore.Encoder {
	//自定义编码配置,下方NewJSONEncoder输出如下的日志格式
	//{"L":"[INFO]","T":"2022-09-16 14:24:59.552","C":"[prototest/main.go:113]","M":"name = xiaoli, age = 18"}
	return zapcore.NewJSONEncoder(NewEncoderConfig())

	//下方NewConsoleEncoder输出如下的日志格式
	//2022-09-16 14:26:02.933 [INFO]  [prototest/main.go:113] name = xiaoli, age = 18
	//return zapcore.NewConsoleEncoder(NewEncoderConfig())
}

// core 三个参数之  日志输出路径
func getInfoWriterSyncer() zapcore.WriteSyncer {
	//file, _ := os.Create("./server/zaplog/log.log")
	//或者将上面的NewMultiWriteSyncer放到这里来，进行返回
	//return zapcore.AddSync(file)

	//引入第三方库 Lumberjack 加入日志切割功能
	infoLumberIO := &lumberjack.Logger{
		Filename:   "./server/zaplog/info.log",
		MaxSize:    10, // megabytes
		MaxBackups: 100,
		MaxAge:     28,    // days
		Compress:   false, //Compress确定是否应该使用gzip压缩已旋转的日志文件。默认值是不执行压缩。
	}
	return zapcore.AddSync(infoLumberIO)
}

func getErrorWriterSyncer() zapcore.WriteSyncer {
	//引入第三方库 Lumberjack 加入日志切割功能
	lumberWriteSyncer := &lumberjack.Logger{
		Filename:   "./server/zaplog/error.log",
		MaxSize:    10, // megabytes
		MaxBackups: 100,
		MaxAge:     28,    // days
		Compress:   false, //Compress确定是否应该使用gzip压缩已旋转的日志文件。默认值是不执行压缩。
	}
	return zapcore.AddSync(lumberWriteSyncer)
}

// Debugf 不再封装使用 - 在显示调用者文件名的时候，会全部显示调用者为logger/zaplogger.go - 所以如果要显示调用者文件名和行号，这里的封装就不合适了
//直接使用 logger.Logger.Info(xxx)
//或者   logger.SugaredLogger.Infof("xxx%s", name)
// logs.Debug(...) 再封装
func Debugf(format string, v ...interface{}) {
	Logger.Sugar().Debugf(format, v...)
}

func Infof(format string, v ...interface{}) {
	Logger.Sugar().Infof(format, v...)
}

func Warnf(format string, v ...interface{}) {
	Logger.Sugar().Warnf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	Logger.Sugar().Errorf(format, v...)
}

func Panicf(format string, v ...interface{}) {
	Logger.Sugar().Panicf(format, v...)
}

// logs.Debug(...) 再封装
func Debug(format string, fileds ...zapcore.Field) {
	Logger.Debug(format, fileds...)
}

func Info(format string, fileds ...zapcore.Field) {
	Logger.Info(format, fileds...)
}

func Warn(format string, fileds ...zapcore.Field) {
	Logger.Warn(format, fileds...)
}

func Error(format string, fileds ...zapcore.Field) {
	Logger.Error(format, fileds...)
}

func Panic(format string, fileds ...zapcore.Field) {
	Logger.Panic(format, fileds...)
}
