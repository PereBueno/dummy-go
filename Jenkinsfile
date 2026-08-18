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

    parameters {
        string 'DEPLOY_ARTIFACT_ID'
    }

    stages {
        stage('Deploy artifact') {
            steps {
                echo "Deploying artifact ID: ${params.DEPLOY_ARTIFACT_ID}"
            }
        }
        stage('Register deployed artifact') {
            steps {
                registerDeployedArtifactMetadata(
                    artifactId: params.DEPLOY_ARTIFACT_ID,
                    targetEnvironment: "PREPROD"
                )
            }
        }
    }
}