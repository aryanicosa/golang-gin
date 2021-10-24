1. init project `go mod init <project-name>`. Then install external library  by using `go get github.com/gin-gonic/gin`

2. create `server.go`

3. create `entity` directory and create `post.go` file for provide Post model

4. create `service` directory and create `post-service.go` file for provide Post service interface
 
5. create `controller` directory and create `post-controller.go` file for provide function controller that handle request to post and get posts

6. create `middleware` directory and create `logger.go & basic-auth.go` files. 
<br>
logger.go provide our own custom log when server received request and save the response into the file `gin.log` and `basic-auth.go` user for creating basic auth security 

7. `go get github.com/tpkeeper/gin-dump` we will user this library to help us debug our web app

8. `go get github.com/go-playground/validator/v10@v10.4.1` adding data binding and validation 

9. Get started to display static html 
- create templates folder then provide `index.html, header.html, footer.html and css file in assets/css folder`
