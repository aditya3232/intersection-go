node {
    stage('Checkout') {
        checkout scm
    }

    stage('Test') {
        docker.image('golang:1.22-alpine').inside {
            sh '''
                go version
                go test ./... -v
            '''
        }
    }

    stage('Build') {
        docker.image('golang:1.22-alpine').inside {
            sh '''
                go build -o intersection-app
            '''
        }
    }
}
