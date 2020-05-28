pipeline {
    agent { dockerfile true }
    stages {
        stage('Test') {
            steps {
                echo 'Hello World JEnkins'
                sh 'docker run -p 8090:8090 ms-go'
            }
        }
    }
}
