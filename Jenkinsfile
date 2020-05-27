pipeline {
    agent any
    stages {
        stage('Test') {
            steps {
                sh 'go run Tupai.go'
                sh 'go --version'
            }
        }
    }
}