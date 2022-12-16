pipeline {
    agent any
    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub-cred')
    }
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
               	sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
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
    post {
        always {
            sh 'docker logout'
        }
    }
}