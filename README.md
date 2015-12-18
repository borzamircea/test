## Deploying the image

After cloning the repository using:

`git clone git@github.com:borzamircea/test.git`


do


`cd test`


and finally, deploy the registry on your `localhost:5000` using:


`docker-compose up -d`


## Pushing and pulling images to the registry


#### Pull the ubuntu image, for example, and tag it:



`docker pull ubuntu && docker tag ubuntu localhost:5000/ubuntu`


#### Push and pull it to the above deployed registry:


`docker push localhost:5000/ubuntu` and  `docker pull localhost:5000/ubuntu`


#### Stopping the registry:


`docker stop registry && docker rm -v registry`





