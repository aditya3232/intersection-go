node {
    properties([
        pipelineTriggers([
            pollSCM('H/2 * * * *')
        ])
    ])

    stage('Checkout') {
        checkout scm
    }

    stage('Test') {
        docker.image('golang:1.25-alpine').inside {
            sh '''
                export GOCACHE=/tmp/go-build
                export GOPATH=/tmp/go

                go version
                go test ./... -v -json > test-report.json
            '''
        }
    }

    stage('Build') {
        docker.image('golang:1.25-alpine').inside {
            sh '''
                export GOCACHE=/tmp/go-build
                export GOPATH=/tmp/go

                go build -o intersection-app
            '''
        }
    }

    stage('Deploy') {
        sh '''
            docker version

            docker build -t intersection-app:latest .

            docker stop intersection-app || true
            docker rm intersection-app || true

            docker run -d \
              --name intersection-app \
              -p 8080:8080 \
              intersection-app:latest
        '''
    }

}
