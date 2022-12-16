pipeline {
    agent any
    stages {
        stage('hello') {
            steps {
                echo 'hello world'
                sh 'pwd'
                sh 'uname -a'
            }
        }
        stage('push image') {
            steps {
                sh 'ls'
                sh 'docker build -t arlsclu/cicd:v0.01 .'
            }
        }
        // stage('testing') {
        //     agent {
        //         docker {
        //             image 'alpine:3.17'
        //         }
        //     }
        //     steps {
        //         /* */
        //     }
        // }
    }
}