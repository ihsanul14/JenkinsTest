pipeline {
    agent { dockerfile true }
    stages {
        stage('Test') {
            steps {
                echo 'ENV Port = $PORT'
            }
        }
    }
}