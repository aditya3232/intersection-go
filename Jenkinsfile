node {
    // Poll SCM tiap 2 menit
    properties([
        pipelineTriggers([
            pollSCM('H/2 * * * *')
        ])
    ])

    stage('Checkout') {
        checkout scm
    }

    stage('Test') {
        echo 'Running unit tests...'
        sh '''
            go version
            go test ./... -v
        '''
    }

    stage('Build') {
        echo 'Building application...'
        sh '''
            go build -o intersection-app
        '''
    }
}
