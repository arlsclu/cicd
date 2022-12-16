pipeline {
    agent any
    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub-cred')
        tag  = 'arlsclu/cicd:v0.01'
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
                sh "docker build -t ${tag} ."
               	sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
                sh "docker push ${tag}"
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