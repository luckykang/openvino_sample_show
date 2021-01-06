import cv2 as cv

# Load face-detection model
face_net = cv.dnn.readNet('/home/kang/open_model_zoo-2019/model_downloader/Transportation/object_detection/face/pruned_mobilenet_reduced_ssd_shared_weights/dldt/face-detection-adas-0001.xml', '/home/kang/open_model_zoo-2019/model_downloader/Transportation/object_detection/face/pruned_mobilenet_reduced_ssd_shared_weights/dldt/face-detection-adas-0001.bin')

# load age-gender detection model
age_gender_net = cv.dnn.readNet('/home/kang/open_model_zoo-2019/model_downloader/Retail/object_attributes/age_gender/dldt/age-gender-recognition-retail-0013.xml','/home/kang/open_model_zoo-2019/model_downloader/Retail/object_attributes/age_gender/dldt/age-gender-recognition-retail-0013.bin')

# Specify target device (CPU)
face_net.setPreferableTarget(cv.dnn.DNN_TARGET_CPU)
age_gender_net.setPreferableTarget(cv.dnn.DNN_TARGET_CPU)

GENDER_LABELS = ['Female', 'Male']

# Read an image
frame = cv.imread('/home/kang/Desktop/human-20180303-01.jpg')

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
        blob2=cv.dnn.blobFromImage(frame[ymin:ymax, xmin:xmax], size=(62,62), ddepth=cv.CV_8U)
        age_gender_net.setInput(blob2)
        detections = age_gender_net.forwardAndRetrieve(['prob','age_conv3'])

        #process results (age, gender)
        gender = GENDER_LABELS[detections[0][0][0].argmax()]
        age = int(detections[1][0][0][0][0][0] * 100)
        text=gender + " : " + str(age)

        #print results to image
        cv.putText(frame,text,(xmin,ymin),cv.FONT_HERSHEY_COMPLEX,0.5,(0,0,255),1)

# Save the frame to an image file
cv.imwrite('age_gender.png', frame)
