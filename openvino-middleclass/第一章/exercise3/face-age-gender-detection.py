import cv2 as cv

# Load face-detection model
face_net = cv.dnn.readNet('/home/kang/open_model_zoo-2019/model_downloader/Transportation/object_detection/face/pruned_mobilenet_reduced_ssd_shared_weights/dldt/face-detection-adas-0001.xml', '/home/kang/open_model_zoo-2019/model_downloader/Transportation/object_detection/face/pruned_mobilenet_reduced_ssd_shared_weights/dldt/face-detection-adas-0001.bin')

# load age-gender detection model
# --> Add your code here

# Specify target device for the faced detection(CPU)
face_net.setPreferableTarget(cv.dnn.DNN_TARGET_CPU)

# Specify target device for the age-gender model (CPU)
#--> Add your code here

#labels for gender     
GENDER_LABELS = ['Female', 'Male']

# Read an image
frame = cv.imread('/home/kang/Desktop/0903171727.png')

# Prepare input blob
blob = cv.dnn.blobFromImage(frame, size=(672, 384), ddepth=cv.CV_8U)

#perform inference (face detection)
face_net.setInput(blob)
out = face_net.forward()


# for each detected face 
for detection in out.reshape(-1, 7):
    confidence = float(detection[2])

    if confidence > 0.5:

        xmin = int(detection[3] * frame.shape[1])
        ymin = int(detection[4] * frame.shape[0])
        xmax = int(detection[5] * frame.shape[1])
        ymax = int(detection[6] * frame.shape[0])
        
        #Draw rectangle over face 
        cv.rectangle(frame, (xmin, ymin), (xmax, ymax), color=(0, 255, 0))

        #perform age-gender detection
        #--> your code here..


        #process results (age, gender)
        #--> your code here..

        #print results to image
        #--> your code here..

# Save the frame to an image file
cv.imwrite('age_gender.png', frame)