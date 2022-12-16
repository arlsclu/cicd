pipeline {
    agent any
    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub-cred')
        tag  = 'arlsclu/cicd:v0.0.2'
    }
    stages {
        stage ('scm') {
            steps {
                sh 'git pull origin main'
            }
        }
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
                sh "docker build -t ${env.tag} ."
               	sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
                sh "docker push ${env.tag}"
            }
        }
        stage('running') {
            agent {
                docker {
                    image 'arlsclu/cicd:v0.0.2'
                }
            }
            steps {
                echo 'xxx'
            }
        }
    }
    post {
        always {
            sh 'docker logout'
        }
    }
}