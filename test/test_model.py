import os
# import logging
import tensorflow as tf
from tensorflow.contrib.layers import fully_connected
from sklearn.preprocessing import StandardScaler
import numpy as np
from sklearn.metrics import accuracy_score
import matplotlib.pyplot as plt
from tensorflow.contrib.layers import batch_norm





class testNet:
    def __init__(self):
        print("enter")
        self.sess = tf.Session()
        self.sess.as_default()
        with tf.variable_scope('mytest'):
            self.x = tf.placeholder(shape=(2,3),dtype=tf.float32,name="input1")
            self.y = tf.placeholder(shape=(2,3),dtype=tf.float32,name="input2")
            # self.test_var = tf.Variable([1,2,3])
            self.ret1 = tf.add(self.x,self.y,"myadd")
            self.ret2 = tf.matmul(tf.transpose(self.x),self.y,name="mymatmul")
        

    def fit(self):
        init = tf.global_variables_initializer()
        out =self.sess.run([self.ret1,self.ret2],feed_dict={self.x:[[1,1,1],[2,2,2]],self.y:[[3,3,3],[4,4,4]]})
        print("same out0={},out1={}".format(out[0],out[1]))
    def build_save(self):
        builder = tf.saved_model.builder.SavedModelBuilder("golangtestmodel")
        builder.add_meta_graph_and_variables(self.sess, ["mytag"])
        builder.save()
        tf.summary.FileWriter("logs/", self.sess.graph)

    def predict(self,X):
        saver = tf.train.Saver()
        with self.sess.as_default():
            # saver.restore(self.sess,"./mymodel.ckpt")
            proability =self.sess.run(self.logits,feed_dict={self.X:X})
            y_pred =np.argmax(proability,axis=1)
            print("predict test_var={}".format(self.test_var.eval()))
            return y_pred



def buildmode():
    net = testNet()
    net.fit()
    net.build_save()
    return


if __name__ == "__main__":
    print("enter")
    buildmode()