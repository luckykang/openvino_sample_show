from __future__ import print_function
import sys
import cv2
import numpy as np

#from openvino.inference_engine import IECore
from openvino.inference_engine import IENetwork, IECore


def main():

    model_xml = '/home/kang/Downloads/public/human-pose-estimation-0001/FP16-INT8/human-pose-estimation-0001.xml'                     #The DL model, IR format 
    model_bin = '/home/kang/Downloads/public/human-pose-estimation-0001/FP16-INT8/human-pose-estimation-0001.bin'

    ie = IECore()                               #Inference-Engine Core object

    net = ie.read_network(model=model_xml, weights=model_bin)    #Read IR  

    input_blob = next(iter(net.inputs))         #first layer of the model
    out_blob = next(iter(net.outputs))          #last layer
    net.batch_size = 1

    #based on last lab,use available device to exercise
    ##-->add your device name here, you can choose taget device to use when you have many device
    exec_net = ie.load_network(network=net, device_name='CPU')


    n, c, h, w = net.inputs[input_blob].shape     #Input dimensions

    image = np.ndarray(shape=(n, c, h, w))

    image = cv2.imread('/home/kang/Desktop/human-20180303-01.jpg')              #read input image
    if image.shape[:-1] != (h, w):                #resize image to match input sizes/shape
        image = cv2.resize(image, (w, h))
        image = image.transpose((2, 0, 1))        # Change data layout from HWC to CHW

    out = exec_net.infer(inputs={input_blob: image})   # Inference 
    out = out[out_blob]

    with open('/home/kang/Desktop/test/第四章/labels.txt', 'r') as f:                 #Read labels file               
        labels_map = [x.split(sep=' ', maxsplit=1)[-1].strip() for x in f]

    for i, probs in enumerate(out):
        probs = np.squeeze(probs)
        top_ind = np.argsort(probs)[-10:][::-1]

        print('\n Class                       Probability')     #print header
        print('---------------------------------------------')

        for id in top_ind:
            det_label = labels_map[id] if labels_map else "{}".format(id)
            print("{:30}{:.7f}".format(det_label, probs[id]))
        print("\n")


        #增加性能计数器（performance counters） 
        performance_counters = exec_net.requests[0].get_perf_counts()

        print('{:<40} {:<15} {:<25} {:<15} {:<10}'.format('name', 'layer_type', 'exet_type', 'status',

      'real_time, us'))

        print("----------------------------------------------------------------------------------------------------------")
        for layer, stats in performance_counters.items():

            print('{:<40} {:<15} {:<25} {:<15} {:<10}'.format(layer, stats['layer_type'], stats['exec_type'],stats['status'], stats['real_time']))



if __name__ == '__main__':
    sys.exit(main() or 0)