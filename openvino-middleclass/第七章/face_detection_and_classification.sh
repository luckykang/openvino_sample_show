#!/bin/bash
# ==============================================================================
# Copyright (C) 2018-2020 Intel Corporation
#
# SPDX-License-Identifier: MIT
# ==============================================================================
set -e

INPUT=${1:"faces.mp4"}

PRECISION=${2:-"FP32"}

MODEL1=face-detection-adas-0001
MODEL2=age-gender-recognition-retail-0013
MODEL3=emotions-recognition-retail-0003
MODEL4=landmarks-regression-retail-0009

DEVICE=CPU

if [[ $INPUT == "/dev/video"* ]]; then
  SOURCE_ELEMENT="v4l2src device=${INPUT}"
elif [[ $INPUT == *"://"* ]]; then
  SOURCE_ELEMENT="urisourcebin buffer-size=4096 uri=${INPUT}"
else
  SOURCE_ELEMENT="filesrc location= faces.mp4"
fi

GET_MODEL_PATH() {
    model_name=$1
    for models_dir in ${MODELS_PATH//:/ }; do
        paths=$(find $models_dir -type f -name "*$model_name.xml" -print)
        if [ ! -z "$paths" ];
        then
            considered_precision_paths=$(echo "$paths" | grep "/$PRECISION/")
           if [ ! -z "$considered_precision_paths" ];
            then
                echo $(echo "$considered_precision_paths" | head -n 1)
                exit 0
dc2-user@10-255-0-77:~/CSDN200/07/exercise-1$ ls
07-exercise1.how  face_detection_and_classification-file.sh  model_proc  op.mp4
answer.mp4        faces.mp4                                  models      output.mp4
dc2-user@10-255-0-77:~/CSDN200/07/exercise-1$ show faces.mp4 
dc2-user@10-255-0-77:~/CSDN200/07/exercise-1$ bash face_detection_and_classification-file.sh
gst-launch-1.0 -e filesrc location= faces.mp4 ! decodebin ! videoconvert ! video/x-raw,format=BGRx ! gvadetect model=/home/dc2-user/CSDN200/07/exercise-1/models/face-detection-adas-0001/FP32/face-detection-adas-0001.xml device=CPU ! queue ! gvaclassify model=/home/dc2-user/CSDN200/07/exercise-1/models/age-gender-recognition-retail-0013/FP32/age-gender-recognition-retail-0013.xml model-proc=./model_proc/age-gender-recognition-retail-0013.json device=CPU ! queue ! gvaclassify model=/home/dc2-user/CSDN200/07/exercise-1/models/emotions-recognition-retail-0003/FP32/emotions-
#!/bin/bash
# ==============================================================================
# Copyright (C) 2020 Intel Corporation
#
# SPDX-License-Identifier: MIT
# ==============================================================================

set -e

# input parameters
FILE=${1:-https://github.com/intel-iot-devkit/sample-videos/raw/master/person-bicycle-car-detection.mp4}
DETECTION_INTERVAL=${2:-10}
INFERENCE_PRECISION=${3:-"FP32"}
INFERENCE_DEVICE=CPU

MODEL_1=person-vehicle-bike-detection-crossroad-0078
MODEL_2=person-attributes-recognition-crossroad-0230
MODEL_3=vehicle-attributes-recognition-barrier-0039

TRACKING_TYPE="short-term"

RECLASSIFY_INTERVAL=10

if [[ $FILE == "/dev/video"* ]]; then
  SOURCE_ELEMENT="v4l2src device=${FILE}"
elif [[ $FILE == *"://"* ]]; then
  SOURCE_ELEMENT="urisourcebin buffer-size=4096 uri=${FILE}"
else
  SOURCE_ELEMENT="filesrc location=${FILE}"
fi

GET_MODEL_PATH() {
    model_name=$1
    precision=${INFERENCE_PRECISION}
    for models_dir in ${MODELS_PATH//:/ }; do
        paths=$(find $models_dir -type f -name "*$model_name.xml" -print)
        if [ ! -z "$paths" ];
        then
            considered_precision_paths=$(echo "$paths" | grep "/$precision/")
           if [ ! -z "$considered_precision_paths" ];
            then
                echo $(echo "$considered_precision_paths" | head -n 1)
                exit 0
            else
                echo $(echo "$paths" | head -n 1)
                exit 0
            fi
        fi
    done

    echo -e "\e[31mModel $model_name file was not found. Please set MODELS_PATH\e[0m" 1>&2
    exit 1
}

PROC_PATH() {
    echo $(dirname "$0")/model_proc/$1.json
}

DETECTION_MODEL=$(GET_MODEL_PATH $MODEL_1)
PERSON_CLASSIFICATION_MODEL=$(GET_MODEL_PATH $MODEL_2)
VEHICLE_CLASSIFICATION_MODEL=$(GET_MODEL_PATH $MODEL_3)

DETECTION_MODEL_PROC=$(PROC_PATH $MODEL_1)
PERSON_CLASSIFICATION_MODEL_PROC=$(PROC_PATH $MODEL_2)
VEHICLE_CLASSIFICATION_MODEL_PROC=$(PROC_PATH $MODEL_3)

PIPELINE="gst-launch-1.0 \
  ${SOURCE_ELEMENT} ! decodebin ! videoconvert ! video/x-raw,format=BGRx ! \
  gvadetect model=$DETECTION_MODEL \
            model-proc=$DETECTION_MODEL_PROC \
            inference-interval=${DETECTION_INTERVAL} \
            threshold=0.6 \
            device=${INFERENCE_DEVICE} ! \
  queue ! \
  gvatrack tracking-type=${TRACKING_TYPE} ! \
  queue ! \
  gvaclassify model=$PERSON_CLASSIFICATION_MODEL \
              model-proc=$PERSON_CLASSIFICATION_MODEL_PROC \
#!/bin/bash
# ==============================================================================
# Copyright (C) 2018-2020 Intel Corporation
#
# SPDX-License-Identifier: MIT
# ==============================================================================
set -e

INPUT=${1:-https://github.com/intel-iot-devkit/sample-videos/raw/master/head-pose-face-detection-female-and-male.mp4}

PRECISION=${2:-"FP32"}

MODEL1=face-detection-adas-0001
MODEL2=age-gender-recognition-retail-0013
MODEL3=emotions-recognition-retail-0003
MODEL4=landmarks-regression-retail-0009

DEVICE=CPU

if [[ $INPUT == "/dev/video"* ]]; then
  SOURCE_ELEMENT="v4l2src device=${INPUT}"
elif [[ $INPUT == *"://"* ]]; then
  SOURCE_ELEMENT="urisourcebin buffer-size=4096 uri=${INPUT}"
else
  SOURCE_ELEMENT="filesrc location=${INPUT}"
fi

GET_MODEL_PATH() {
    model_name=$1
    for models_dir in ${MODELS_PATH//:/ }; do
        paths=$(find $models_dir -type f -name "*$model_name.xml" -print)
        if [ ! -z "$paths" ];
        then
            considered_precision_paths=$(echo "$paths" | grep "/$PRECISION/")
           if [ ! -z "$considered_precision_paths" ];
            then
                echo $(echo "$considered_precision_paths" | head -n 1)
                exit 0
            else
                echo $(echo "$paths" | head -n 1)
                exit 0
            fi
        fi
    done

    echo -e "\e[31mModel $model_name file was not found. Please set MODELS_PATH\e[0m" 1>&2
    exit 1
}

PROC_PATH() {
    echo $(dirname "$0")/model_proc/$1.json
}

DETECT_MODEL_PATH=$(GET_MODEL_PATH $MODEL1)
CLASS_MODEL_PATH=$(GET_MODEL_PATH $MODEL2)
CLASS_MODEL_PATH1=$(GET_MODEL_PATH $MODEL3)
CLASS_MODEL_PATH2=$(GET_MODEL_PATH $MODEL4)

MODEL2_PROC=$(PROC_PATH $MODEL2)
MODEL3_PROC=$(PROC_PATH $MODEL3)
MODEL4_PROC=$(PROC_PATH $MODEL4)

PIPELINE="gst-launch-1.0 $SOURCE_ELEMENT ! \
decodebin ! videoconvert ! video/x-raw,format=BGRx ! \
gvadetect model=$DETECT_MODEL_PATH device=$DEVICE ! queue ! \
gvaclassify model=$CLASS_MODEL_PATH model-proc=$MODEL2_PROC device=$DEVICE ! queue ! \
gvaclassify model=$CLASS_MODEL_PATH1 model-proc=$MODEL3_PROC device=$DEVICE ! queue ! \
gvaclassify model=$CLASS_MODEL_PATH2 model-proc=$MODEL4_PROC device=$DEVICE ! queue ! \
gvafpscounter ! fakesink"


echo ${PIPELINE}
${PIPELINE}