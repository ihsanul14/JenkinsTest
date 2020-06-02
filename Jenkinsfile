pipeline {
    agent any
    stages {
        stage('Test') {
            steps {
                echo 'Hello World JEnkins'
                bat 'docker run -p 8090:8090 -d ms-go'      

            }
        }
    }
}
