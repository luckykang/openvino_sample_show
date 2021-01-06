from __future__ import print_function
import sys
import cv2
import numpy as np

from openvino.inference_engine import IECore

def main():

    model_xml = 'model.xml'                     #The DL model, IR format 
    model_bin = 'model.bin'

    ie = IECore()                               #Inference-Engine Core object

    ##print available devices.
    #--> Your code here
    print('请将可用设备打印于下方：')
    devices = ie.available_devices
    print(devices)   # ['CPU', 'GNA']


if __name__ == '__main__':
    sys.exit(main() or 0)