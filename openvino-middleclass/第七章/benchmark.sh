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
dc2-user@10-255-0-77:~/CSDN200/07/exercise-3$ bash face_detection_and_classification.sh road.mp4 
gst-launch-1.0 filesrc location=road.mp4 ! decodebin ! videoconvert ! video/x-raw,format=BGRx ! gvadetect model=/home/dc2-user/CSDN200/07/exercise-3/models/face-detection-adas-0001/FP32/face-detection-adas-0001.xml device=CPU ! queue ! gvaclassify model=/home/dc2-user/CSDN200/07/exercise-3/models/age-gender-recognition-retail-0013/FP32/age-gender-recognition-retail-0013.xml model-proc=./model_proc/age-gender-recognition-retail-0013.json device=CPU ! queue ! gvaclassify model=/home/dc2-user/CSDN200/07/exercise-3/models/emotions-recognition-retail-0003/FP32/emotions-recognition-retail-0003.xml model-proc=./model_proc/emotions-recognition-retail-0003.json device=CPU ! queue ! gvaclassify model=/home/dc2-user/CSDN200/07/exercise-3/models/landmarks-regression-retail-0009/FP32/landmarks-regression-retail-0009.xml model-proc=./model_proc/landmarks-regression-retail-0009.json device=CPU ! queue ! gvafpscounter ! fakesink
Setting pipeline to PAUSED ...
Pipeline is PREROLLING ...
Redistribute latency...
Redistribute latency...
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
            else
                echo $(echo "$paths" | head -n 1)
                exit 0
            fi
        fi
    done

    echo -e "\e[31mModel $model_name file was not found. Please set MODELS_PATH\e[0m" 1>&2
    exit 1
#!/bin/bash
# ==============================================================================
# Copyright (C) 2020 Intel Corporation
#
# SPDX-License-Identifier: MIT
# ==============================================================================

set -e

if [ -z ${1} ]; then
  echo "ERROR set path to video"
  echo "Usage : ./benchmark.sh VIDEO_FILE [DECODE_DEVICE] [INFERENCE_DEVICE] [CHANNELS_COUNT]"
  echo "You can download video with \"cd /path/to/your/video/ && wget https://github.com/intel-iot-devkit/sample-videos/raw/master/head-pose-face-detection-female-and-male.mp4\""
  echo " and run sample ./benchmark.sh /path/to/your/video/head-pose-face-detection-female-and-male.mp4"
  exit
fi

SOURCE_ELEMENT="filesrc location=${1}"
DECODE_DEVICE=${2:-CPU}
INFERENCE_DEVICE=${3:-CPU}
CHANNELS_COUNT=${4:-1}


##please choose your model:
MODEL=face-detection-adas-0001


GET_MODEL_PATH() {
    model_name=$1
    precision="FP32"
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
    echo ./model_proc/$1.json
}

DETECT_MODEL_PATH=$(GET_MODEL_PATH $MODEL )

if [ $DECODE_DEVICE == CPU ]; then
  unset GST_VAAPI_ALL_DRIVERS
  VIDEO_PROCESSING="decodebin ! videoscale ! video/x-raw"
  PRE_PROC=ie
else
  export GST_VAAPI_ALL_DRIVERS=1
  VIDEO_PROCESSING="decodebin ! vaapipostproc ! video/x-raw(memory:VASurface)"
  PRE_PROC=vaapi
fi

PIPELINE=" ${SOURCE_ELEMENT} ! ${VIDEO_PROCESSING} ! \
gvadetect model-instance-id=inf0 model=${DETECT_MODEL_PATH} device=${INFERENCE_DEVICE} pre-process-backend=${PRE_PROC} ! queue ! \
gvafpscounter ! fakesink "

FINAL_PIPELINE_STR=""

for (( CURRENT_CHANNELS_COUNT=0; CURRENT_CHANNELS_COUNT < $CHANNELS_COUNT; ++CURRENT_CHANNELS_COUNT ))
do
  FINAL_PIPELINE_STR+=$PIPELINE
done


echo "gst-launch-1.0 -v ${FINAL_PIPELINE_STR}"
gst-launch-1.0 -v ${FINAL_PIPELINE_STR}