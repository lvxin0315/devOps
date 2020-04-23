package devOps

import (
	"errors"
	"github.com/sirupsen/logrus"
	"os"
)

const DockerffmpegToM3U8Cmd = "docker run -v {{.Mp4FileDir}}:/temp/ jrottenberg/ffmpeg " +
	"-i /temp/{{.FileName}} " +
	"-f segment -segment_time 10 -segment_format mpegts " +
	"-segment_list /temp/{{.OutPutDir}}/{{.OnlyName}}.m3u8 " +
	"-c copy -bsf:v h264_mp4toannexb -map 0 /temp/{{.OutPutDir}}/{{.OnlyName}}-%08d.ts"

const DockerffmpegCutJPEGCmd = "docker run -v {{.Mp4FileDir}}:/temp/ jrottenberg/ffmpeg " +
	"-ss {{.CutTime}} " +
	"-i /temp/{{.FileName}} " +
	"-r 1 " +
	"-vframes 1 " +
	"-y /temp/{{.OutPutDir}}/{{.OnlyName}}.jpeg"

type DockerFFmpegMp4ToM3u8 struct {
	mp4FilePath  string
	mp4FileDir   string
	fileName     string
	onlyName     string
	outputDir    string
	jpegFileName string
}

func (f *DockerFFmpegMp4ToM3u8) setMp4FilePath(mp4FilePath string) error {
	if f.mp4FilePath != "" {
		return nil
	}

	logrus.Info(mp4FilePath)

	_, err := os.Stat(mp4FilePath)
	if os.IsNotExist(err) {
		err := errors.New("mp4文件不存在")
		logrus.Error("file.Exist：", err)
		return err
	}
	f.mp4FilePath = mp4FilePath
	return nil
}

func (f *DockerFFmpegMp4ToM3u8) cutMp4FileName(outputDir string) error {
	if f.mp4FileDir != "" && f.fileName != "" {
		return nil
	}
	//根据文件地址截取文件名称
	mp4FileDir, fileName, onlyName, _ := cutFileName(f.mp4FilePath)
	//处理输出目录
	err := dirStat(mp4FileDir + outputDir)
	if err != nil {
		return err
	}
	f.onlyName = onlyName
	f.outputDir = outputDir
	f.mp4FileDir = mp4FileDir
	f.fileName = fileName
	return nil
}

/**
使用docker 进行mp4文件转换m3u8
@param string mp4FilePath mp4文件地址
@param string outputDir 相对mp4文件地址的目录
*/
func (f *DockerFFmpegMp4ToM3u8) DockerMp4ToM3u8(mp4FilePath string, outputDir string) error {
	err := f.setMp4FilePath(mp4FilePath)
	if err != nil {
		return err
	}
	err = f.cutMp4FileName(outputDir)
	if err != nil {
		return err
	}
	//执行转换命令
	params := make(map[string]interface{})
	params["Mp4FileDir"] = f.mp4FileDir
	params["FileName"] = f.fileName
	params["OnlyName"] = f.onlyName
	params["OutPutDir"] = outputDir
	err = cmd(DockerffmpegToM3U8Cmd, params)
	if err != nil {
		return err
	}
	return nil
}

/**
使用docker 进行mp4文件截图
@param string mp4FilePath mp4文件地址
@param string outputDir 相对mp4文件地址的目录
@param string ssTime 时间格式, 例如00:01:22
*/
func (f *DockerFFmpegMp4ToM3u8) DockerMp4CutJpeg(mp4FilePath string, outputDir string, ssTime string) error {
	err := f.setMp4FilePath(mp4FilePath)
	if err != nil {
		return err
	}
	err = f.cutMp4FileName(outputDir)
	if err != nil {
		return err
	}
	//执行转换命令
	params := make(map[string]interface{})
	params["Mp4FileDir"] = f.mp4FileDir
	params["FileName"] = f.fileName
	params["OnlyName"] = f.onlyName
	params["OutPutDir"] = f.outputDir
	params["CutTime"] = ssTime

	err = cmd(DockerffmpegCutJPEGCmd, params)
	if err != nil {
		return err
	}
	//jpeg文件名
	f.jpegFileName = f.onlyName + ".jpeg"
	return nil
}

func (f *DockerFFmpegMp4ToM3u8) getJpegFileName() string {
	return f.jpegFileName
}
