pipeline {
    agent any
    stages {
        stage('Example') {
            steps {
                echo 'hello world'
                sh 'ls /'
                sh 'uname -a'
                sh 'ls -alh /var/jenkins_home/workspace/pipe_0'
            }
        }
    }
}