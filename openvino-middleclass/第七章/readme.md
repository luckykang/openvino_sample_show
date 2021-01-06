## Gstreamer流水线视频处理

### 1.如何模拟部署一个完整AI应用的流程
测量性能的最佳方法是，快速构建一个真实流水线，这个流水线应用要看起来尽可能接近我们构建的实际应用，使用的输入内容最好是我们在产品中使用的视频。视频中实际操作量的大小会影响进行解码所需要的算力，如果是camera或者RTSP流，请模拟真实对象。输出内容也尽量要符合这样的要求。如果用户需要实时的查看输出视频，或者可能存储或者上传内容时，视频实际操作量的大小同时会影响编码所需要的算力。如果不需要输出，可以省去成本高昂的编码流程，或者编码较低的分辨率、较低的帧速率都会造成实际的影响。现在，让我们可以模拟整个流程。

### 2.DL-Streamer
Deep-Learning Streamer是一个框架，可以构建计算图和流水线，它被包含在OpenVINO中。可以把开发的视频分析元素添加到Gstreamer elements列表中，这些新元素用openvino来检测、分类、识别、跟踪、可视化等操作。DL-Streamer有c++、python和纯Gstreamer脚本实例。



实验一:Gstreamer构建视频处理流水线

age-gender-recognition-retail-0013.json

bash face_detection_and_classification-file.sh

实验二：运行车辆追踪处理流水线

bash vehicle_pedestrian_tracking-file.sh road.mp4

实验三：通道性能评估实验

bash face_detection_and_classification.sh road.mp4

# 1.输入视频 2.解码设备 2. 推理设备 3.通道数
bash benchmark.sh road.mp4 CPU CPU 1 

其实执行的是下面的指令

gst-launch-1.0 -v  filesrc location=road.mp4 ! decodebin ! videoscale ! video/x-raw ! gvadetect model-instance-id=inf0 model=/home/dc2-user/CSDN200/07/exercise-3/models/face-detection-adas-0001/FP32/face-detection-adas-0001.xml device=CPU pre-process-backend=ie ! queue ! gvafpscounter ! fakesink 

不过请注意：
"gvadetect"   用于 detection models
(like the face detection or the person-vehicle-bike-detection-crossroad-0078)
"gvaclassify"  用于classification models.