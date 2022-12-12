First clean the docker images 
docker rm -f $(docker ps -aq)
next prune the volume and system for cleaning docker image files from the system
docker volume prune;docker system prune 
start the network using script file (./startfabric.sh)

check docker ps -a

now we need to run test-sdk with "npm start" command

if you don't have node modules you can create with npm i command

now open post man application

we need to enroll admin first
then admin can register the users in org1 and org2 based on our inputs

couch db browser :
http://localhost:5984/_utils/#login
http://localhost:7984/_utils/#login

JSON file:
create:
{
    "userId": "srinivasD@nasa.org",
    "ownerOrgId": "1",
    "data": {
        "modelId": "10",
        "modelName": "Detailed-Report-10",
        "projectId": "PDC-1",
        "modelType": "Detailed",
        "assetType": "MBSEModel"
    }
}
update:
{
  "userId": "srinivasD@nasa.org",
  "ownerOrgId": "1",
  "data" :{
	"modelId":"101",
	"modelName":"Detail-Report-101 updated by srinivas in 2022",
	"projectId":"Project-Devisetti srinivas",
	"modelType":"Detailed",
    "version": "1.1" ,
     "assetType": "MBSEModel"
  }
}
Delete
{
    "userId": "srinivasD@nasa.org",
    "ownerOrgId": "1",
    "data": {
        "modelId": "101",
        "modelType": "Detailed",
         "assetType": "MBSEModel"
    }
}

after updating chain code we need to run below command
deploy chaincode again:

./network.sh deployCC -ccv 2 -ccs 2


