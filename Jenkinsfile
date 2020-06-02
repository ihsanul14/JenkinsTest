pipeline {
	agent any 
    stages {
        stage('Test') {
            steps {
                echo 'Hello World JEnkins'
		
		bat 'kubectl apply -f ms-go.yaml'      
		
            }
        }
    }
}
