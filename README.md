January 2022
My purpose here is to learn basics of Go, containerize, deploy in cloud and be able to program working with 
cloud native services like databases and messaging to be able to create components myself that integrate with Ab Initio for demos.
For example create a cloud service that generate data to feed Ab Initio to demonstrate certain Ab Initio features.

I decided to start with an empty webserver using Golang Tutorials on YouTube

1. Setup environment Go on VSCode
Best I found was this 3 1/2 hour crash course on Go: www.youtube.com/watch?v=yyUHQIec83I
To enable GitHub, watch this: www.youtube.com/watch?v=3Tn58KQvWtU

2. Create program empty-webserver (doesn't do anything)
Comes from www.youtube.com/watch?v=0sRjYzL_oYs

3. Run the program
Create a Terminal window
Type 'go run main.go'
Open Chrome and type localhost:8080

4. Debugging
To debug, we first need to create a configuation. Click on left pane "Run and Debug" icon.
Use menu to create a launch.json in .vscode directory. We also need a go.mod file in there.

5. Run on GCP Compute Engine
Comes from www.youtube.com/watch?v=uL1YERoqGQc
Go to GCP and create a new VM preemtible
Open SSH and install Go from the Go website:
    $sudo apt install wget // apt=Advanced Package Tool, this is to install stuff on Debian Linux
    $wget https://go.dev/dl/go1.17.6.linux-amd64.tar.gz // wget is to get from HTTP -> download Go from Go website
    As per Go website installation instructions (copy/pasted) - MAKE SURE TO ADD SUDO
    $sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz
    $export PATH=$PATH:/usr/local/go/bin
    $go version // Test that installation is good
Copy the source code
    $cd $HOME
    $mkdir -p go/src/github.com/jpchauvet/empty-webserver
    ZIP the file-upload directory on Windows machine
    Use the SSH window wheel menu to Upload the ZIP file
    Then install unzip as follow:
    $sudo apt update && sudo apt install -y unzip
    $unzip empty-webserver.zip -d go/src/github.com/jpchauvet/empty-webserver/
Now run the GO program
    $cd go/src/github.com/jpchauvet/empty-webserver/empty-webserver/
    $go run main.go
    This is NOT going to work because by default, GCP doesn't allow any open port (except 80 and 443)
Establish a new firewall rule to allow 8080
    In GCP console, go to VCP Network/Firewall and Create New Firewall rule
    Set Name, Targets=All instances in the network, Source IPv4 Ranges=0.0.0.0/0 (the whole internet), tcp=8080
    Create. This will create the rule. MAKE SURE YOU DELETE IT AFTERWARDS
Now run the GO program again and test with Chrome connecting to the VM external IP :8080

6. Run on AWS EC2 machine
Go to AWS and create a new free tier EC2 machine
Download the .pem file to ~/Downloads (call it something like aws20220119.pem)
No need to install anything (it seems all installed) but just in case, the equivalent of 'apt' is 'yum'
Install Go like on GCP
Use a Windows cmd and copy the ZIP file as follow:
    >cd c:\Users\JP\Downloads
    >scp -i aws20220119.pem D:\Programming\Go\empty-webserver\empty-webserver.zip ec2-user@ec2-3-138-173-225.us-east-2.compute.amazonaws.com
Then SSH to the machine
    >ssh -i aws20220119.pem ec2-user@ec2-3-138-173-225.us-east-2.compute.amazonaws.com
Establish a new firewall rule to allow 8080
    In Instances, select the EC2 Instance then Security tab, then the Security groups/launch-wizard-N
    In Actions, select Edit Inbound rules
    Add Rule Custom TCP for port 8080 from anywhere. Save Rules
    GO back to Instances, click on EC@ instance and see rule has been added.
Now run the GO program again and test with Chrome connecting to the VM external IP :8080

7. Study how to containerize server with Docker and run on CloudRun
Add a Dockerfile to the project
Go to GCP Console and open the terminal
Check that Go and docker are installed
    $go version
    $docker version
Upload the source code to the GCP terminal
Create a repository called "jp-repo" to upload my container image into GCP Artifact Registry
    $gcloud artifacts repositories create --location us --repository-format docker jp-repo
Then set the name of the container image (in the right repository) we are about to build
    $IMAGE=us-docker.pkg.dev/my-playground-project-332514/jp-repo/empty-webserver
Now build the container in the directory where the source code is
    $docker build . -t $IMAGE -f Dockerfile
Check that the image has been built
    $docker images
Now push it to the Artifact Registry
    $docker push $IMAGE
Now deploy it to CloudRun
    $gcloud run deploy empty-webserver --image $IMAGE --allow-unauthenticated
To shutdown all services running
    $gcloud run services list
    $gcloud run services delete SERVICE



Next: 
- Study Git to understand how it is configured for my project
- Study how to containerize server with Docker on AWS
- Study how to run on serverless CloudRun and Lambda
- Study how server can connect to a cloud database and exchange data
- Study how server can create data that can be used for feeding Ab Initio
