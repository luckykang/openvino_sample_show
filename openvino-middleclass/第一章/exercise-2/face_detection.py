import cv2 as cv

'''
input: 1x3x384x672 BGR
output: 1,1,N,7
'''



# 加载模型 
net = cv.dnn.readNet('face-detection-adas-0001.xml','face-detection-adas-0001.bin')
# 指定运行的硬件设备为CPU
net.setPreferableTarget(cv.dnn.DNN_TARGET_CPU)
# 读取图像
frame = cv.imread('faces.jpg')
# 设置图像大小先宽后高，并进行推理准备
blob = cv.dnn.blobFromImage(Frame,size=(672,384), ddepth=cv.CV_8U)
# 运行推理
net.setInput(blob)
out = net.forward()
# 获得结果，并在检测到的面部上画框 Nx7,N是面部数量
for detection in out.reshape(-1,7):
    # 置信度是列表中的第三项
    confidence = float(detection[2])
    xmin = int(detection[3] * frame.shape[1])
    ymin = int(detection[4] * frame.shape[0])
    xmax = int(detection[5] * frame.shape[1])
    ymax = int(detection[6] * frame.shape[0])
    # 为了识别更多的人脸，可以调节检测的置信度，使得程序对人脸的判定没那么严格，这里下调为0.2(原来是0.5)
    if confidence > 0.2:
        cv.rectangle(frame, (xmin,ymin), (xmax,ymax), color=(0,255,0))
    
# 保存图片
cv.imwrite('out.png',frame)



