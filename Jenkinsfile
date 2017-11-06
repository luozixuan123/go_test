podTemplate(label: 'mypod', containers: [
    containerTemplate(name: 'golang', image: 'golang:1.8.0', ttyEnabled: true, command: 'cat'),
  ]) {

    node('mypod') {

        stage('Get a Golang project') {
            git branch: 'master', credentialsId: 'd7760e1d-9ba9-4026-9725-38ef955da704', url: 'https://github.com/luozixuan123/go_test.git'
            container('golang') {
                stage('Build a Go project') {
                    sh 'go test ./test'
                }
            }
        }

    }
}