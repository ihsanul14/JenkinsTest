pipeline {
    agent {
        kubernetes {
        yaml """
        apiVersion: apps/v1
kind: Deployment
metadata:
  name: ms-go-deployment
  labels:
    app: micros-go
spec:
  replicas: 1
  selector:
    matchLabels:
      app: micros-go
  template:
    metadata:
      labels:
        app: micros-go
    spec:
      containers:
        - name: microsrvice-go
          image: ihsanul14/microsergo:tesgo
          ports:
            - containerPort: 8090
        #   env:
        #     - name: MONGO_INITDB_ROOT_USERNAME
        #       valueFrom:
        #         secretKeyRef:
        #           name: mongodb-secret
        #           key: mongo-root-username
        #     - name: MONGO_INITDB_ROOT_PASSWORD
        #       valueFrom:
        #         secretKeyRef:
        #           name: mongodb-secret
        #           key: mongo-root-password
---
apiVersion: v1
kind: Service
metadata:
  name: ms-go-service
spec:
  selector:
    app: micros-go
  type: LoadBalancer
  ports:
    - protocol: TCP
      port: 8090
      targetPort: 8090
      nodePort: 30001

        """
    }
    }
    stages {
        stage('Test') {
            steps {
                echo 'Hello World JEnkins'
                   
            }
        }
    }
}
