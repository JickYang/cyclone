/*
Copyright 2016 caicloud authors. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package parser

import (
	"testing"

	"github.com/caicloud/cyclone/worker/ci/yaml"
)

const configStr = `
integration:
  services: # run some micro services container which depend by the integration test
    mongo:
      image: cargo.caicloud.io/caicloud/mongo:3.0.5
      command: mongod --smallfiles
  image: cargo.caicloud.io/caicloud/golang:1.6 # run a golang container to compile executable files and do integration test
  environment:
    - BUILD_DIR=/go/src/github.com/caicloud/ci-demo-go
  commands: 
    - mkdir -p $BUILD_DIR
    - cp ./ -rf $BUILD_DIR
    - cd $BUILD_DIR/code
    - go build -v -o app # compile executable files 
    - $BUILD_DIR/code/app & # run executable files 
    - echo "do some test" # do integration test

pre_build:
  image: cargo.caicloud.io/caicloud/golang:1.6 # run a container to compile publish executable files
  volumes:
    - .:/go/src/github.com/caicloud/ci-demo-go # mount source file to GOPATH
  commands: # compile
    - echo "compile executable files"
    - cd /go/src/github.com/caicloud/ci-demo-go/code
    - go build -v -o app
  outputs: # copy out publish executable files from prebuild container
    - /go/src/github.com/caicloud/ci-demo-go/code/app

build: #build pubilsh image
  dockerfile_name: Dockerfile_publish
`

// TestParser tests parse function.
func TestParser(t *testing.T) {
	config, err := yaml.ParseString(configStr)
	if err != nil {
		t.Error("Expect error to be nil")
	}

	_, err = Load(config)
	if err != nil {
		t.Error("Expect error to be nil")
	}
}
