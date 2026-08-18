pipeline {
    agent {
        kubernetes {
            yaml '''
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: shell
    image: ubuntu
    command:
    - sleep
    args:
    - infinity
    securityContext:
      # ubuntu runs as root by default, it is recommended or even mandatory in some environments (such as pod security admission "restricted") to run as a non-root user.
      runAsUser: 1000
'''
        }
    }

    stages {
        stage('Hello') {
            steps {
                echo 'Hello World'
                sh 'echo "This is my archive" > archive.txt'
                archiveArtifacts artifacts: 'archive.txt', followSymlinks: false
            }
        }
        stage('Register build artifact') {
            steps {
                script {
                    env.ARTIFACT_ID = registerBuildArtifactMetadata(
                        name: "my-artifact",
                        version: "1.0.0",
                        url: "https://sda.general.support.beescloud.com/dse-team-amer/job/ddewhurst/job/unify_dora_mb_test/job/upstream-branch/$BUILD_NUMBER/artifact/archive.txt"
                    )
                }
            }
        }
        stage('Print artifact id') {
            steps {
                echo "Artifact ID: ${env.ARTIFACT_ID}"
            }
        }
    }
}