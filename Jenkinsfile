node {
  stage ("scm") {
    checkout scm
    sh '''
      #!/bin/bash
      ./scripts/jenkins_setup_git.bash
    '''
  }

  wrap([$class: 'AnsiColorBuildWrapper', 'colorMapName': 'XTerm']) {
    stage ("deps") {
      sh '''
        #!/bin/bash
        ./scripts/jenkins_setup_deps.bash
      '''
    }
  }
}
