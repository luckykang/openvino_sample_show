## 视频处理

### 一.视频分析推理流程

解码 -- 预处理 -- 推理 -- 后处理 -- 编码

![0831143415](https://cdn.jsdelivr.net/gh/luckykang/picture_bed/blogs_images/0831143415.png)

### 二.执行视频解码的几种方法：
- Media-SDK 

    Intel推荐使用的解码以及图像处理。如果系统中有Intel集成显卡并且安装了Media-SDK，将使用专用硬件执行编解码操作，称之为`Intel Quick Sync`，也可以在CPU上运行，并且使用CPU软件来完成加速。使用起来有点难度，不容易上手。
- OpenCV  

    通过OpenCV调用Media-SDK执行硬件加速，这可能是最简单的视频处理方法。

- GStreamer/ffmpeg
  
    OpenCV也可以调用这两个软件的功能。

### 三.预处理

解码完成后我们需要获取每一帧的图像，以匹配神经网络的输入格式。

- 图像处理以提高质量
- 将多个流聚合为一批
- 缩放，调整大小，切出感兴趣的区域
- 仅选择选定的帧

### 四.推理

- 使用深度学习的模型推理
- 可以是并行、串行的多个模型

### 五.后处理

推理阶段完成后需要的处理，比如在屏幕显示结果或者打印标签等等。

### 六.编码

在获得视频流处理结果时，要重新对视频流进行压缩，这样可以对视频进行存储，或者是上传到云端


python3 multi_camera_multi_person_tracking.py   --m_detector models/person-detection-retail-0013.xml --m_reid models/person-reidentification-retail-0300.xml --config config.py  -i Videos/video1.avi Videos/video2.avi Videos/video3.avi --no_show    --output out.avi