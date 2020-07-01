1.python中加载openvino
把openvino安装目录 C:\Intel\openvino\python\python3.6 下的 openvino文件夹   
粘贴到python安装目录下   C:\Python36\Lib\site-packages

2.找代码demo
C:\Intel\openvino_2019.1.087\deployment_tools\inference_engine\samples\python_samples\object_detection_demo_ssd_async目录下
object_detection_demo_ssd_async.py文件代码拿来使用

3.编译
C:\Intel\openvino\deployment_tools\inference_engine\samples 路径下 执行
build_samples_msvc2017.bat

执行完后在C:\Users\kang\Documents\Intel\OpenVINO 目录生成
inference_engine_samples_build_2017 文件夹

cpu_extension = "C:\Users\kang\Documents\Intel\OpenVINO\inference_engine_samples_build_2017\intel64\Release\cpu_extension.dll"


4.下载模型 road-segmentation-adas-0001 记录xml地址


5.填写参数
plugin_dir = ""    
model_xml = ""   
model_bin = "" 
cap = cv2.VideoCapture("")


